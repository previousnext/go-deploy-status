FROM golang:1.8
ADD workspace /go
RUN go get github.com/mitchellh/gox
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/CHANGE_ME_linux_amd64 /usr/local/bin/CHANGE_ME
CMD ["CHANGE_ME"]
