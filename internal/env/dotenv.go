package env

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func InitDotenv() {
	dotenvfiles := []string{}

	appenv := os.Getenv("APP_ENV")
	if appenv != "" {
		dotenvfiles = append(dotenvfiles, ".env."+appenv)
	}

	dir, _ := os.Getwd()

	root := filepath.VolumeName(dir)
	files := []string{}

	for dir != root+"/" && dir != root+"\\" {
		for _, dotenvfile := range dotenvfiles {

			filename := filepath.Join(dir, dotenvfile)
			_, err := os.Stat(filename)
			if os.IsNotExist(err) {
				continue
			}

			files = append(files, filename)
		}

		dir = filepath.Dir(dir)
	}

	godotenv.Load(files...)
}
