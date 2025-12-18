# --- ETAPA 1: Builder (Golang) ---
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copiamos todo y compilamos
COPY . .
# CGO_ENABLED=0 es OBLIGATORIO para scratch (quita dependencias de C)
RUN CGO_ENABLED=0 go build -o netspy main.go

# --- ETAPA 2: Runner (Scratch/Vacío) ---
FROM scratch

# 1. Copiamos el binario
COPY --from=builder /app/netspy /

# 2. Copiamos los certificados SSL (¡CRÍTICO PARA HTTPS!)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Ejecutamos
CMD ["/netspy"]
