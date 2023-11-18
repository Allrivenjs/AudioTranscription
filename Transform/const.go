package Transform

const (
	// SpleeterCMD     = "python3.10 -m spleeter separate -p spleeter:2stems -o %s %s"
	// donde: %s: la ruta del archivo de audio
	// -p spleeter:2stems: el modelo de separación de audio
	// -o %s: la ruta de salida
	// %s: la ruta del archivo de audio
	SpleeterCMD = "python3.10 -m spleeter separate -p spleeter:2stems -o %s %s"

	// FfmpegCMD       = "ffmpeg -i %s -acodec pcm_s16le -f s16le -ac 1 -ar 44100 %s"
	// %s: la ruta del archivo de audio
	// -acodec pcm_s16le: el códec de audio
	// -f s16le: el formato de audio
	// -ac 1: el número de canales
	// -ar 44100: la frecuencia de muestreo
	// %s: la ruta del archivo de salida
	FfmpegCMD = "ffmpeg -i %s -acodec pcm_s16le -f s16le -ac 1 -ar 44100 %s"

	// WhispetCMD usa el comando whisper para transcribir el audio
	// %s: la ruta del archivo de audio
	// --model medium: el modelo de transcripción
	// --output_format txt: el formato de salida
	// --language es: el idioma
	// --output_dir %s: la ruta de salida
	WhispetCMD = "whisper %s --model medium --output_format txt --task transcribe --language es --output_dir %s" //whisper on_process_files/audioTest1/vocals.wav --model medium --output_format txt --task transcribe --language es --output_dir on_process_files/audioTest1/

	// getDuration usa el comando ffmpeg para obtener la duración del archivo de audio
	// %s: la ruta del archivo de audio
	// 2>&1: redirige la salida de error a la salida estándar
	// grep Duration: busca la línea que contiene la duración
	// awk '{print $2}': imprime la segunda columna
	// tr -d ,: elimina la coma
	getDuration = "ffmpeg -i %s 2>&1 | grep Duration | awk '{print $2}' | tr -d ,"

	// BaseOutputParts es el formato del nombre de los archivos de audio cortados
	// %s: el prefijo del archivo de salida
	// %d: el número de parte
	BaseOutputParts = "/%s_%d.mp3"

	// CutAudioCMD  usa el comando ffmpeg para cortar el audio donde se le indique
	// %s: la ruta del archivo de audio
	// -ss %s: el punto de corte
	// -t %s: la duración del audio
	// -c copy %s: la ruta del archivo de salida
	CutAudioCMD = "ffmpeg -i %s -ss %s -c copy %s" //"ffmpeg -i path -ss "10 -t 20" -c copy output.mp3"
)
