FROM golang:alpine as builder
ENV CGO_ENABLED=0
ENV GOOS=linux 
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]

# https://www.cloudreach.com/en/resources/blog/cts-build-golang-dockerfiles
# https://stackoverflow.com/a/62123648 --> avoid errors for during `go build` command
