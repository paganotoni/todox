FROM golang:1.20-alpine as builder
RUN apk --update add build-base curl

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

CMD GO_ENV=production tools migrate && app