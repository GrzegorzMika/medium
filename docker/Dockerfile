FROM golang:1.22 as build

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o http_server main.go

FROM scratch as production

COPY --from=build /src/http_server .

EXPOSE 8080

CMD [ "/http_server" ]