# AudioTranscription
Sistema de Procesamiento y Transcripción de Audio con Spleeter y Whisper

[//]: # (https://wavesurfer.xyz/)


### Use
##### list
```{"ggml-tiny.en", "ggml-tiny", "ggml-base.en", "ggml-base", "ggml-small.en", "ggml-small", "ggml-medium.en", "ggml-medium", "ggml-large-v1", "ggml-large-v2", "ggml-large-v3"}```
```bash
./build/go-model-download -out models ggml-small.en.bin
```

```bash
./build/go-whisper -model ./models/ggml-small.en.bin -out-file testdata/text2.txt testdata/jfk.wav
```
Donde:
- -model representa el modelo a usar, 
- -out-file representa el archivo donde se escribira los datos. ese se creara automaticamente.
- `samples\jfk.wav` es la ruta donde se encuentra el audio a decodificar. 

#### Generate documentation
```bash
 swag init --parseDependency --parseInternal 
 ```


#### Install air -d
https://github.com/cosmtrek/air

```bash 
go install github.com/cosmtrek/air@latest
```

```bash
air
```



### links 

https://github.com/ggerganov/whisper.cpp/tree/master/bindings/go
https://github.com/djthorpe/go-whisper

https://wavesurfer.xyz/