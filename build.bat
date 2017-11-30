set GOARCH=amd64
set GOOS=linux

go build -o sha256sum-web

docker build -t mailtruck/sha256sum-web:latest .
docker push mailtruck/sha256sum-web:latest