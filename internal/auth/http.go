package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	usr "github.com/sekalahita/epirus/internal/domain/appuser/usecase"
)

type AuthHandler struct {
	Service usr.UserUseCase
}

const sessionName = "graphqlapiserver"

type userInfo struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Picture       string `json:"picture"`
}

func (h AuthHandler) LoginStartHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	config := GoogleOauthConfig(r)

	authUrl, err := url.Parse(config.Endpoint.AuthURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Set("response_type", "codes")
	params.Set("client_id", config.ClientID)
	params.Set("scope", strings.Join(config.Scopes, " "))
	params.Set("redirect_uri", config.RedirectURL)
	params.Set("response_type", "code")
	params.Set("state", "code")
	authUrl.RawQuery = params.Encode()

	if err = saveRedirect(w, r); err != nil {
		return nil, err
	}

	http.Redirect(w, r, authUrl.String(), http.StatusFound)
	return nil, nil
}

func saveRedirect(w http.ResponseWriter, r *http.Request) error {
	redirect := "/"
	redirectParam := r.URL.Query()["redirect"]
	if len(redirectParam) != 0 && redirectParam[0] != "" {
		redirect = redirectParam[0]
	} else if referer := r.Referer(); referer != "" {
		redirect = referer
	}

	// Parse it here, so we can ignore the errors later
	_, err := url.Parse(redirect)
	if err != nil {
		return err
	}

	return nil
}

type LoginCallbackResponse struct {
	Token string `json:"token"`
}

func (h AuthHandler) LoginCallbackHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	config := GoogleOauthConfig(r)

	code := r.FormValue("code")
	if code == "" {
		response := "code is empty"
		if r.FormValue("error_reason") == "user_denied" {
			response = "user denied permission"
		}

		return nil, errors.New(response)
	}

	token, err := config.Exchange(r.Context(), code)
	if err != nil {
		return nil, err
	}

	userinfoResponse, err := http.Get(
		"https://openidconnect.googleapis.com/v1/userinfo?access_token=" + url.QueryEscape(token.AccessToken),
	)
	if err != nil {
		return nil, err
	}
	defer userinfoResponse.Body.Close()

	body, err := ioutil.ReadAll(userinfoResponse.Body)
	if err != nil {
		return nil, err
	}

	var userInfo userInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, err
	}

	user, err := h.Service.LoginOrSignUp(r.Context(), usr.LoginOrSignUpParam{
		GoogleID: userInfo.Sub,
		Email:    userInfo.Email,
	})
	if err != nil {
		return nil, err
	}

	claims := NewClaims(
		WithUserID(user.ID),
		WithAudience([]string{"epirus-app"}),
		WithExpiresAt(time.Now().Add(1*time.Hour)),
		WithNotBefore(time.Now()),
	)

	jwt, err := Encode(claims)
	if err != nil {
		return nil, err
	}

	json.NewEncoder(w).Encode(LoginCallbackResponse{
		Token: jwt,
	})

	return nil, nil
}
