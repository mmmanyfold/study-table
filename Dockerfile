FROM golang:alpine as builder
ARG COMMIT
ARG AIRTABLE_API_KEY

ENV COMMIT ${COMMIT:-undefined}
ENV AIRTABLE_API_KEY ${AIRTABLE_API_KEY:-undefined}

RUN echo $COMMIT
RUN echo $AIRTABLE_API_KEY

WORKDIR /app 

COPY . .

RUN ls

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" ./src

FROM scratch

WORKDIR /app

# COPY --from=builder /app/ /usr/bin/

# RUN ls /usr/bin

EXPOSE 8080

# ENTRYPOINT ["dev-to"]