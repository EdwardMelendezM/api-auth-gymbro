# Etapa de construcción
FROM golang:1.22 AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar go.mod y go.sum y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/myapp

# Etapa final: usar una imagen base ligera de Alpine
FROM alpine:latest

# Instalar certificados raíz para conexiones HTTPS y zonas horarias
RUN apk --no-cache add ca-certificates tzdata

# Establecer la zona horaria
ENV TZ=America/Lima

# Copiar el ejecutable desde la etapa de construcción
COPY --from=builder /app/myapp /app/myapp

EXPOSE 80

# Establecer el comando de entrada del contenedor, ejecucion
CMD ["/app/myapp"]