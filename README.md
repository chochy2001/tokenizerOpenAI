Here's the README.md in English and Spanish:

---

# OpenAI Tokenizer in Go

This project provides a simple utility to tokenize text using the OpenAI tokenizer in Go. It helps you determine the number of tokens in a given text, which is useful for adjusting the input to fit within the model's training threshold and text quantity limits.

## Project Structure

The project has the following structure:

```
.
├── go.mod
├── go.sum
└── main.go
```

- `go.mod`: The Go module file that defines the project's dependencies.
- `go.sum`: The checksum file for the project's dependencies.
- `main.go`: The main Go source file containing the tokenizer code.

## Usage

To run the tokenizer, follow these steps:

1. Make sure you have Go installed on your system.
2. Clone this repository and navigate to the project directory.
3. Run the following command to execute the tokenizer:

```
go run main.go
```

The program will output the total number of tokens in the provided text.

## Dependencies

The project relies on the following dependencies:

- `github.com/pkoukk/tiktoken-go`: A Go port of the OpenAI tiktoken library for tokenizing text.

These dependencies are automatically managed by Go modules.

## Customization

You can customize the input text by modifying the `text` variable in the `main.go` file. Simply replace the existing text with your desired input.

Feel free to explore and modify the code to suit your specific requirements.

---

# Tokenizador de OpenAI en Go

Este proyecto proporciona una utilidad simple para tokenizar texto utilizando el tokenizador de OpenAI en Go. Te ayuda a determinar la cantidad de tokens en un texto dado, lo cual es útil para ajustar la entrada y que se ajuste al umbral de entrenamiento y los límites de cantidad de texto del modelo.

## Estructura del Proyecto

El proyecto tiene la siguiente estructura:

```
.
├── go.mod
├── go.sum
└── main.go
```

- `go.mod`: El archivo de módulo de Go que define las dependencias del proyecto.
- `go.sum`: El archivo de suma de comprobación para las dependencias del proyecto.
- `main.go`: El archivo fuente principal de Go que contiene el código del tokenizador.

## Uso

Para ejecutar el tokenizador, sigue estos pasos:

1. Asegúrate de tener Go instalado en tu sistema.
2. Clona este repositorio y navega hasta el directorio del proyecto.
3. Ejecuta el siguiente comando para ejecutar el tokenizador:

```
go run main.go
```

El programa mostrará el número total de tokens en el texto proporcionado.

## Dependencias

El proyecto depende de las siguientes dependencias:

- `github.com/pkoukk/tiktoken-go`: Un puerto en Go de la biblioteca tiktoken de OpenAI para tokenizar texto.

Estas dependencias son gestionadas automáticamente por los módulos de Go.

## Personalización

Puedes personalizar el texto de entrada modificando la variable `text` en el archivo `main.go`. Simplemente reemplaza el texto existente con la entrada deseada.

Siéntete libre de explorar y modificar el código según tus requisitos específicos.
