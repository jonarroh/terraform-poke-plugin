# Stage 1: Build
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Instalar herramientas básicas
RUN apk add --no-cache git

# Copiar código fuente
COPY . .

# Descargar dependencias
RUN go mod tidy

# Compilar el provider
RUN go build -o terraform-provider-pokemon

# Stage 2: Imagen mínima para correr el provider
FROM alpine:latest

WORKDIR /app

# Copiar binario compilado desde builder
COPY --from=builder /app/terraform-provider-pokemon .

# El provider debe ser ejecutable
RUN chmod +x terraform-provider-pokemon

ENTRYPOINT ["./terraform-provider-pokemon"]
