# Reference: https://github.com/GoogleContainerTools/distroless/blob/main/examples/go/Dockerfile

FROM golang:1.21.6-alpine as builder
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11
ENV APP_ENV=production
COPY --from=builder /go/bin/app /
EXPOSE 3000
CMD ["/app"]