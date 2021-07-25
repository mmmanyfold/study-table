FROM golang:alpine as builder

ENV COMMIT ${COMMIT:-undefined}
ENV AIRTABLE_API_KEY ${AIRTABLE_API_KEY:-undefined}
ENV AWS_ACCESS_KEY_ID ${AWS_ACCESS_KEY_ID:-undefined}
ENV AWS_SECRET_ACCESS_KEY ${AWS_SECRET_ACCESS_KEY:-undefined}

WORKDIR /app/study-table-service

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/sts .

FROM alpine:3.14
RUN apk add ca-certificates

COPY --from=builder /app/study-table-service/out/sts /app/sts

EXPOSE 8080

CMD ["/app/sts"]