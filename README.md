# AudioTranscription
Sistema de Procesamiento y Transcripción de Audio con Spleeter y Whisper

[//]: # (https://wavesurfer.xyz/)


### Use
```bash
./build/go-whisper -model ./models/ggml-small.en.bin -out-file ./samples/text2.txt samples/jfk.wav
```
Donde:
- -model representa el modelo a usar, 
- -out-file representa el archivo donde se escribira los datos. ese se creara automaticamente.
- `samples\jfk.wav` es la ruta donde se encuentra el audio a decodificar. 

#### Generate documentation
```bash swag init --parseDependency --parseInternal ```


### links 

https://github.com/ggerganov/whisper.cpp/tree/master/bindings/go
https://github.com/djthorpe/go-whisper

https://wavesurfer.xyz/