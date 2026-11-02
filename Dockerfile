# -------------------------------------------------------------
# STAGE 1: Builder
# -------------------------------------------------------------
FROM golang:1.24-alpine AS builder

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

# Bien dich tinh hoan toan, loai bo debug symbols
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-w -s" -o /bin/app .

# -------------------------------------------------------------
# STAGE 2: Distroless Runtime (Zero-Attack Surface)
# -------------------------------------------------------------
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /bin/app /app/server

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/app/server"]
