FROM golang:1.20-alpine as builder
RUN apk --update add build-base curl

RUN wget https://github.com/benbjohnson/litestream/releases/download/v0.3.9/litestream-v0.3.9-linux-amd64-static.tar.gz
RUN tar -xzf litestream-v0.3.9-linux-amd64-static.tar.gz -C /usr/local/bin

WORKDIR /src/todox
ADD go.mod .
RUN go mod download

ADD . .
RUN make build

FROM alpine

RUN apk add --no-cache bash ca-certificates make

WORKDIR /bin/

# Copying binaries
COPY --from=builder /src/todox/bin/app .
COPY --from=builder /src/todox/bin/tools .
COPY --from=builder /usr/local/bin/litestream /usr/local/bin

CMD ./run.sh