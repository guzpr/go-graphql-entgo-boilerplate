set -e
#!/bin/sh

go run -mod=mod entgo.io/ent/cmd/ent new --target internal/ent/schema $@