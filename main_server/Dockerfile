FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

# dependencies
COPY ["go.mod", "go.sum", "/"]
RUN go mod download

# build
COPY . ./
RUN go build -o ./bin/app cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /
COPY configs/config.yaml /config.yaml

# envs
ENV SERVER_CONFIG="config.yaml"

CMD [ "/app" ]