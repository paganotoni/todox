FROM golang:1.22-alpine as builder
RUN apk --update add build-base curl

WORKDIR /src/todox
ADD go.mod .
RUN go mod download

ADD . .

# Building TailwindCSS with tailo
RUN go run github.com/paganotoni/tailo/cmd/build@a4899cd

# Installing kit
RUN go install github.com/leapkit/leapkit/kit@v0.0.7

# Building the app
RUN go build -tags osusergo,netgo -buildvcs=false -o bin/app ./cmd/app

FROM alpine
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

# Copying binaries
COPY --from=builder /go/bin/kit .
COPY --from=builder /src/todox/bin/app .
COPY --from=builder /src/todox/internal/migrations migrations

CMD kit db migrate --migrations.folder=migrations && app
