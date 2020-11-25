FROM golang:latest as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o calcul cmd/main.go


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/calcul .
EXPOSE 8081

CMD ["./calcul"] 