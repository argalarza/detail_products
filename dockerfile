# Usa la imagen oficial de Golang
FROM golang:1.20-alpine

# Establece el directorio de trabajo
WORKDIR /app

# Copia el archivo go.mod y go.sum
COPY go.mod go.sum ./

# Instala las dependencias
RUN go mod tidy

# Copia el código fuente
COPY . .

# Compila la aplicación
RUN go build -o product-service .

# Expone el puerto 8080
EXPOSE 3009

# Comando para ejecutar la aplicación
CMD ["./product-service"]
