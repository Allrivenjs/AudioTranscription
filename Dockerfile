FROM golang:1.21-alpine3.18 as SOURCE

WORKDIR /app

RUN apk update && apk add --no-cache make git gcc g++ && \
  rm -rf /var/cache/apk/*

COPY docker .

RUN make dependency && make build && \
  mv bin/go-whisper /bin/ && \
  rm -rf bin && \
  apk del make git gcc g++

FROM python:3.11.6

LABEL org.opencontainers.image.source=https://github.com/allrivenjs/AudioTranscription
LABEL org.opencontainers.image.description="Speech-to-Text."
LABEL org.opencontainers.image.licenses=MIT

RUN apt-get update && apt-get install -y --no-install-recommends ffmpeg libsndfile1 && \
  rm -rf /var/lib/apt/lists/*

RUN python3 -m pip install --no-cache-dir \
  git+https://github.com/dunossauro/videomaker-helper.git@1fd99ec

WORKDIR /app

COPY . .

ENTRYPOINT ["go run main.go"]
