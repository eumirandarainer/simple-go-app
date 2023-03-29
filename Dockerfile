FROM golang:1.17-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/simple-go-app

FROM scratch
COPY --from=build /bin/simple-go-app /bin/simple-go-app
ENTRYPOINT ["/bin/simple-go-app"]
