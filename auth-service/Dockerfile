FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/myapp .
COPY --from=builder /app/config/model.conf ./config/
COPY --from=builder /app/config/policy.csv ./config/
COPY ./email/verification_email.html /app/email/

COPY .env .
EXPOSE 8080
CMD ["./myapp"]
