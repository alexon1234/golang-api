FROM golang:alpine AS build

WORKDIR /go/src/github.com/alexon1234/golang-api
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/golang-api -ldflags "$LDFLAGS" cmd/api/main.go

FROM scratch
COPY --from=build /go/bin/golang-api /go/bin/golang-api
EXPOSE 8000
ENTRYPOINT [ "/go/bin/golang-api" ]