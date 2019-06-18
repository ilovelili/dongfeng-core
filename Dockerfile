FROM golang:1.11 as builder

LABEL maintainer="<route666@live.cn>"

# enable go mod
ENV GO111MODULE=on

ENV SRC_DIR=/go/src/dongfeng/dongfeng-core
WORKDIR $SRC_DIR

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . $SRC_DIR

WORKDIR $SRC_DIR/services/server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM scratch
LABEL maintainer="<route666@live.cn>"

WORKDIR /app
ENV BUILDER_DIR=/go/src/dongfeng/dongfeng-core/services/server

COPY --from=builder $BUILDER_DIR/config.*.json $BUILDER_DIR/server /app/

ENTRYPOINT ["/app/server"]