# OpenAI Tokenizer en Go

Este proyecto proporciona una utilidad para tokenizar texto utilizando el tokenizador de OpenAI implementado en Go. La tokenización es un proceso fundamental en el procesamiento de lenguaje natural (NLP) que divide el texto en unidades más pequeñas llamadas tokens, lo que resulta útil para:

- Calcular límites de tokens para APIs de modelos de lenguaje como GPT
- Optimizar el uso de recursos en aplicaciones de NLP
- Preparar texto para análisis y procesamiento avanzado

## 📋 Estructura del Proyecto

El proyecto tiene la siguiente estructura:

```
.
├── go.mod          # Define las dependencias del proyecto
├── go.sum          # Checksums de las dependencias
├── main.go         # Código principal con la lógica de tokenización
└── prueba.go       # Contiene el texto de ejemplo para tokenizar
```

## 🔧 Requisitos

- Go 1.18 o superior (desarrollado con Go 1.22.2)
- Dependencias gestionadas automáticamente con Go Modules

## 📦 Dependencias

El proyecto utiliza las siguientes bibliotecas:

- `github.com/pkoukk/tiktoken-go v0.1.7`: Implementación en Go del tokenizador tiktoken de OpenAI
- `github.com/dlclark/regexp2 v1.10.0`: Biblioteca de expresiones regulares avanzadas
- `github.com/google/uuid v1.3.0`: Generación de identificadores únicos

## 🚀 Instalación

```bash
# Clonar el repositorio
git clone https://github.com/tu-usuario/tokenizerGoOpenAI.git
cd tokenizerGoOpenAI

# Instalar dependencias
go mod download
```

## 💻 Uso

### Ejecución Básica

Para ejecutar el tokenizador con el texto de ejemplo predefinido:

```bash
go run main.go prueba.go
```

El programa mostrará el número total de tokens en el texto proporcionado.

### Personalización del Texto

Puedes modificar la función `textToTokenize()` en el archivo `prueba.go` para cambiar el texto que deseas tokenizar:

```go
func textToTokenize() string {
    // Reemplaza este texto con el que deseas tokenizar
    text := `Tu texto aquí...`
    return text
}
```

### Integración en tu Propio Código

Para usar esta funcionalidad en tu proyecto:

```go
package main

import (
    "fmt"
    "github.com/pkoukk/tiktoken-go"
)

func main() {
    // Selecciona el modelo de codificación (cl100k_base es usado por GPT-4, GPT-3.5-Turbo)
    encoding := "cl100k_base"
    
    // Inicializa el codificador
    tke, err := tiktoken.GetEncoding(encoding)
    if err != nil {
        fmt.Printf("Error al obtener la codificación: %v\n", err)
        return
    }
    
    // Tu texto para tokenizar
    text := "Este es un ejemplo de texto para tokenizar."
    
    // Codifica el texto en tokens
    tokens := tke.Encode(text, nil, nil)
    
    // Muestra el resultado
    fmt.Printf("Número total de tokens: %d\n", len(tokens))
    
    // Para ver los tokens individuales (opcional)
    fmt.Println("Tokens:", tokens)
}
```

## 📊 Modelos de Tokenización Disponibles

La biblioteca `tiktoken-go` soporta varios modelos de tokenización de OpenAI:

- `cl100k_base`: Usado por GPT-4 y GPT-3.5-Turbo
- `p50k_base`: Usado por modelos GPT-3 de texto (Davinci, Curie, etc.)
- `r50k_base`: Usado por modelos anteriores de GPT-3
- Y otros modelos de codificación soportados por la biblioteca

## 🔄 Modo Sin Conexión

Si prefieres no descargar el diccionario en tiempo de ejecución, puedes usar el cargador sin conexión:

```go
import (
    "github.com/pkoukk/tiktoken-go"
    "github.com/pkoukk/tiktoken-go/loader"
)

func main() {
    // Configura el cargador sin conexión
    tiktoken.SetBpeLoader(loader.NewOfflineLoader())
    
    // Continúa con la codificación como se mostró anteriormente
    // ...
}
```

## 🧪 Ejemplos Adicionales

### Decodificación de Tokens a Texto

```go
// Codifica el texto en tokens
text := "Ejemplo de decodificación"
tokens := tke.Encode(text, nil, nil)

// Decodifica los tokens de vuelta a texto
decodedText := tke.Decode(tokens)
fmt.Println("Texto decodificado:", string(decodedText))
```

### Conteo de Tokens para APIs de OpenAI

```go
func estimateTokenCount(text string) int {
    encoding := "cl100k_base"
    tke, _ := tiktoken.GetEncoding(encoding)
    return len(tke.Encode(text, nil, nil))
}

// Ejemplo de uso
prompt := "Explica cómo funciona la tokenización de texto."
tokenCount := estimateTokenCount(prompt)
fmt.Printf("Este prompt utilizará aproximadamente %d tokens\n", tokenCount)
```

## 📝 Nota sobre la Precisión

El conteo de tokens puede variar ligeramente entre diferentes implementaciones de tokenizadores. Esta implementación en Go intenta ser lo más fiel posible al tokenizador oficial de OpenAI.

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Si encuentras algún problema o tienes alguna sugerencia, por favor crea un issue o envía un pull request.

## 📄 Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo LICENSE para más detalles.

---

Creado por [Jorge Salgado Miranda] - [2025]
