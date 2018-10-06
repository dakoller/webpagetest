docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app golang:1.11.1 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o main'

docker build --rm -f "Dockerfile" -t webpagetest:latest .

docker run --rm -it webpagetest:latest
