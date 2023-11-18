import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns  # Importa seaborn para paletas de colores


def cal_spleeter(df):
    # Convertir los valores de "Cal" de milisegundos a segundos.
    df['Cal'] = df['Cal'] / 1000

    # Agrega el número total de segundos de "Cal" para nombres de archivo duplicados.
    total_cal = df.groupby('Name')['Cal'].sum().reset_index()

    # Agrega el número de elementos promediados al nombre de cada barra.
    total_cal['NameWithCount'] = total_cal.apply(
        lambda row: f'{row["Name"]} ({len(df[df["Name"] == row["Name"]]):d})', axis=1)

    # Crea una paleta de colores para las barras.
    colors = sns.color_palette('hsv', len(total_cal))

    # Crea la gráfica de barras directamente desde el DataFrame.
    ax = total_cal.plot.bar(x='NameWithCount', y='Cal', color=colors, figsize=(10, 6))
    ax.set_xlabel('File Name (Number of Averaged Elements)')
    ax.set_ylabel('Total Seconds')
    ax.set_title('Total Seconds per File')

    # Personaliza el formato de las etiquetas en el eje X.
    plt.xticks(rotation=45, ha='right')

    # Mostrar los valores de total de segundos encima de las barras.
    for p in ax.patches:
        ax.annotate(f'{p.get_height():.2f}', (p.get_x() + p.get_width() / 2., p.get_height() + 100), ha='center')

    # Guarda la gráfica en un archivo de imagen (por ejemplo, en formato PNG).
    plt.savefig('statistics/grafica_resultados.png', dpi=300, bbox_inches='tight')

    plt.tight_layout()
    plt.show()


def cal_whisper(df):
    # Convertir los valores de "Cal" de milisegundos a segundos.
    df['Cal'] = df['Cal'] / 1000

    # Agrega el número total de segundos de "Cal" para nombres de archivo duplicados.
    total_cal = df.groupby('Name')['Cal'].sum().reset_index()

    # Agrega el número de elementos promediados al nombre de cada barra.
    total_cal['NameWithCount'] = total_cal.apply(
        lambda row: f'{row["Name"]} ({len(df[df["Name"] == row["Name"]]):d})', axis=1)

    # Crea una paleta de colores para las barras.
    colors = sns.color_palette('hsv', len(total_cal))

    # Crea la gráfica de barras directamente desde el DataFrame.
    ax = total_cal.plot.bar(x='NameWithCount', y='Cal', color=colors, figsize=(10, 6))
    ax.set_xlabel('File Name (Number of Averaged Elements)')
    ax.set_ylabel('Total Seconds')
    ax.set_title('Total Seconds per File')

    # Personaliza el formato de las etiquetas en el eje X.
    plt.xticks(rotation=45, ha='right')

    # Mostrar los valores de total de segundos encima de las barras.
    for p in ax.patches:
        ax.annotate(f'{p.get_height():.2f}', (p.get_x() + p.get_width() / 2., p.get_height() + 100), ha='center')

    # Guarda la gráfica en un archivo de imagen (por ejemplo, en formato PNG).
    plt.savefig('statistics/grafica_resultados_spleeter.png', dpi=300, bbox_inches='tight')

    plt.tight_layout()
    plt.show()


def cal_table(df):
    # Agrega el número total de segundos de "Cal" para nombres de archivo duplicados.
    total_cal = df.groupby('Name')['Cal'].sum().reset_index()
    # Aplica el formato de estilo al DataFrame.
    styled_df = total_cal.style \
        .format({"Cal": "{:.8f}"}) \
        .format_index(str.upper, axis=1) \
        .set_caption("Promedio de Tiempo de procesado por Archivo")

    # Guarda la tabla generada en un archivo CSV.
    styled_df.to_latex('statistics/tabla_resultados_whisper.tex')
    styled_df.to_excel('statistics/tabla_resultados_whisper.xlsx')
    # Muestra el DataFrame formateado.
    styled_df


def main():
    # Supongamos que tienes un archivo CSV llamado "resultados.csv" con la estructura dada.
    # Carga los datos desde el archivo CSV.
    df = pd.read_csv('statistics/output_spleeter.csv')

    cal_spleeter(df)
    # cal_whisper(df)
    # cal_table(df)


if __name__ == '__main__':
    main()
