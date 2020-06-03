FROM golang:1.14-stretch AS build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64 
WORKDIR /build
COPY . .
RUN go get -v ./...
RUN go build -a -installsuffix cgo -o k8s-send-container-logs-to-tg

FROM scratch
COPY --from=build /build/k8s-send-container-logs-to-tg /usr/local/bin/k8s-send-container-logs-to-tg
ENTRYPOINT ["/usr/local/bin/k8s-send-container-logs-to-tg"]
