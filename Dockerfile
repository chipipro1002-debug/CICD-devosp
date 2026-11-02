# -------------------------------------------------------------
# STAGE 1: Builder (Môi trường biên dịch đầy đủ)
# -------------------------------------------------------------
FROM golang:1.23-alpine AS builder

WORKDIR /src

# Tận dụng Docker layer cache cho dependencies
COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY . .

# Biên dịch tĩnh hoàn toàn: loại bỏ debug symbols (-w -s) để giảm dung lượng
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-w -s" -o /bin/app .

# -------------------------------------------------------------
# STAGE 2: Distroless Runtime (Zero-Attack Surface)
# -------------------------------------------------------------
# Image này KHÔNG có shell (/bin/sh), KHÔNG có package manager
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Copy duy nhất binary từ Stage builder
COPY --from=builder /bin/app /app/server

# Chạy dưới tài khoản không đặc quyền (nonroot: UID 65532)
USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/app/server"]