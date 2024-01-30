FROM hub.hamdocker.ir/library/golang:1.18 as builder

WORKDIR /app


COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server


FROM hub.hamdocker.ir/library/alpine

COPY --from=builder /app/server /app/server

CMD ["/app/server"]

