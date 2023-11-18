import click
from pydub import AudioSegment


@click.command()
@click.option('--path', '-p', type=str, help='Path of file', required=True)
@click.option('--output', '-o', type=str, help='Path of output file', required=True)
@click.option('--start', '-s', type=int, help='Start of cut file seconds', required=True)
@click.option('--end', '-n', type=int, help='Ended of cut file seconds', required=True)
def cut_audio(path, output, start, end):
    print("Cortando audio...")
    print("Path: ", path)
    print("Output: ", output)
    print("Start: ", start)
    print("End: ", end)
    # Cargar el archivo de audio
    audio = AudioSegment.from_file(path, format="mp3")

    # Calcular los milisegundos correspondientes a los segundos de inicio y finalización
    start_ms = start * 1000
    end_ms = end * 1000

    # Recortar el audio
    audio_recortado = audio[start_ms:end_ms]

    # Guardar el audio recortado en el archivo de salida
    audio_recortado.export(output, format="mp3")


if __name__ == "__main__":
    try:
        cut_audio()
    except Exception as e:
        print(e)
