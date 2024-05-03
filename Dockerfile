# Usa una imagen de Go como base
FROM golang:latest

# Instala ffmpeg
RUN apt-get update && apt-get install -y ffmpeg

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de la aplicación al contenedor
COPY . .

# Compila la aplicación de Go
RUN go build -tags netgo -ldflags '-s -w' -o app

# Comando por defecto para ejecutar la aplicación
CMD ["./app"]
