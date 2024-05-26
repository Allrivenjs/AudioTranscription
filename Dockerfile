# Usa una imagen de Go como base
FROM golang:latest

# Instala ffmpeg
RUN apt-get update && apt-get install -y ffmpeg

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src/app

# Copia los archivos de la aplicación al contenedor
COPY . .
RUN mkdir -p /go/src/app/storage/app/
RUN mkdir -p /go/src/app/storage/app/audio
RUN mkdir -p /go/src/app/storage/app/temp

RUN chmod -R 777 /go/src/app/storage/app/
# Compila la aplicación de Go
RUN go build -tags netgo -ldflags '-s -w' -o app

# Comando por defecto para ejecutar la aplicación
CMD ["./app"]
