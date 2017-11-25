set GOARCH=amd64
set GOOS=linux

go build


docker build -t mailtruck/kitchen:latest .
docker push mailtruck/kitchen:latest