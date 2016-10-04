# Laboratorio 1

## Problema 1

Usted trabaja en una firma de análisis financiero llamada Economics International Trade (EIT), donde le solicitan
crear un programa que pueda analizar los valores del dólar.
Sin embargo, la empresa no dispone de estos datos, pero uno de tus colegas ha desarrollado una interfaz que realiza
[screen scraping](https://es.wikipedia.org/wiki/Screen_scraping) a la página de Servicio de Impuestos Internos
que obtiene toda esta información en formato CSV.

Con esta información, la empresa necesita resolver las siguientes consultas:

1. Obtener los valores únicos del dólar ordenados de menor a mayor.
1. Buscar el valor de dólar en un fecha específica.
1. Buscar fecha de acuerdo a un valor exacto de dólar o el más cercano.

### CSV

Los valores del CSV vienen ordenados de la siguiente manera:

- Cada línea representa un día del mes; las líneas están ordenadas de manera ascendente.
- Cada valor separado por comas representa un mes, ordenado de manera ascendente.

Por ejemplo, la siguiente tabla:

| Día | Enero | Febrero | Marzo | Abril | Mayo  | Junio | Julio | Agosto | Septiembre | Octubre | Noviembre | Diciembre |
|-----|-------|---------|-------|-------|-------|-------|-------|--------|------------|---------|-----------|-----------|
| 1   | 600.0 | 700.0   | 550.0 | 600.0 | 700.0 | 500.0 |       | 600.0  | 600.0      | 600.0   | 600.0     | 600.0     |
| 2   | 500.0 | 800.0   |       | 600.0 | 700.0 | 500.0 | 600.0 | 600.0  | 600.0      |         | 600.0     | 600.0     |
| 3   | 400.0 | 700.0   | 650.0 | 600.0 | 700.0 | 500.0 | 600.0 |        | 600.0      |         | 600.0     | 600.0     |

Se representa como:

```
600.0,700.0,550.0,600.0,700.0,500.0,,600.0,600.0,600.0,600.0,600.0
500.0,800.0,,600.0,700.0,500.0,600.0,600.0,600.0,,600.0,600.0
400.0,700.0,650.0,600.0,700.0,500.0,600.0,,600.0,,600.0,600.0
```

Para obtener el csv se debe hacer uso del paquete `scraping` de la siguiente manera:

```go
var csv string = scraping.GetCSV()
```

### Formato de entrada

El primer caracter de cada linea indica el número de la pregunta propuesta. La pregunta `2` esta seguida por un string con el
formato `día/mes` y la pregunta `3` es seguida por un decimal que representa el valor del dólar a consultar.

**Ejemplo**

```
1
2 3/4
3 702.32
```

### Formato de salida

Para cada una de las líneas en la entrada debe responder según la pregunta:

- Para la pregunta `1` debe imprimir los números ordenados separados por un espacio.
- Para la pregunta `2` debe el número.
- Para la pregunta `3` debe imprimir una lista de strings con el formato `dia/mes` ordenadas por mes y luego el día, separados
  por un espacio.

## Problema 2

Se le pide construir un programa que verifique si una secuencia
de caracteres contiene todos sus `brackets` equilibrados.

Un `bracket` puede ser cualquiera de los siguientes caracteres:
`(` , `)` , `[` , `]` , `{` o `}`.

Se entiende como bracket equilibrado donde, para cada abertura, existe
su respectivo cierre, por ejemplo:

1. `()`
1. `[]`
1. `{}`
1. `{([()])}`.

Además, debe cumplir las siguientes condiciones:

- Se debe respetar el orden en que se abren y cierran los brackets,
  ya que no es válido cerrar un bracket que no ha sido abierto.
- Un bracket balanceado puede estar contenido en otro bracket balanceado.

Dada `T` secuencias de brackets, determine si cada una de ellas se encuentra
completamente equilibrada. Si la secuencia se encuentra equilibrada, imprima
`YES`, de lo contario imprima `NO`, separando cada caso por un salto de línea.

### Formato de entrada

La primera línea contiene un entero `T`, indicando la cantidad de casos a verificar.
Las siguientes `T` lineas consisten de una secuencia, `s`, de brackets.

### Restricciones

- 1 ≤ `n` ≤ 1000
- 1 ≤ `len(s)` ≤ 1000; donde `len(s)` es largo de la secuencia de brackets.
- Cada caracter dentro de la secuencia será un bracket, (p.e., `(`,`)`, `[`,`]`, `{`,`}`).

### Formato de salida

Para cada secuencia de brackets, imprimir si se ecuentra o no balanceado con un salto de linea.
Si se encuentran balanceados, imprimir `YES`, de lo contrario, imprimir `NO`.

### Ejemplo

#### Entrada

```
({{}}){}[]
({)}[]
{()}[()]]
{}{}[][()]
```

#### Salida

```
YES
NO
NO
YES
```

#### Explicación

1. La secuencia `({{}}){}[]` se encuentra completamente equilibrada, ya que
cumple con ambas condiciones.

2. La secuencia `({)}[]` no se encuentra equilibrada, ya que el primer paréntesis `(`
contiene al bracket `{` no equilibrado.

3. La secuencia `{()}[()]]` no se encuentra completamente equilibrada,
 dado que el último caracter es un bracket `]` que nunca se abrió.

1. La secuencia `{}{}[][()]` se encuentra completamente equilibrada, ya que
cumple con ambas condiciones.
