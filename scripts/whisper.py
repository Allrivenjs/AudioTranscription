import os

import click
import whisper


@click.command()
@click.option('--audio-path', '-a', type=str, help='Ruta del archivo de audio', required=True)
@click.option('--output-path', '-o', type=str, help='Ruta del archivo de salida para el texto', required=True)
def audio_to_text(audio_path, output_path):
    model = whisper.load_model("medium")
    try:
        # Cargar el archivo de audio y realizar la transformación a texto.
        result = model.transcribe(audio_path, 'es-ES')

        # Guardar el texto resultante en el archivo de salida, crear archivo. si no existe.
        if not os.path.exists(output_path):
            os.makedirs(output_path)

        with open(output_path, 'w') as text_file:
            text_file.write(result["text"])

        # Imprimir un mensaje de éxito.
        click.echo(f'Transformación de audio a texto completada. Resultado guardado en {output_path}')
    except Exception as e:
        # En caso de error, mostrar un mensaje de error.
        click.echo(f'Error al transformar audio a texto: {str(e)}', err=True)
        raise click.Abort()


if __name__ == '__main__':
    audio_to_text()
