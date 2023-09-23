FROM golang:1.21-alpine as builder
RUN apk --update add build-base curl

WORKDIR /src/todox
ADD go.mod .
RUN go mod download

ADD . .
RUN go build -o bin/db ./cmd/db
RUN go run ./cmd/build

FROM alpine
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

# Copying binaries
COPY --from=builder /src/todox/bin/app .
COPY --from=builder /src/todox/bin/db .

CMD /bin/db migrate && /bin/app