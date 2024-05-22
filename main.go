package main

import (
	"fmt"

	"github.com/pkoukk/tiktoken-go"
)

func main() {
	text := `
Texto de prueba para tokenizar.
Este es un texto de prueba para tokenizar.
Este es un texto de prueba para tokenizar.
Este es un texto de prueba para tokenizar.
`
	encoding := "cl100k_base"

	// if you don't want download dictionary at runtime, you can use offline loader
	// tiktoken.SetBpeLoader(tiktoken_loader.NewOfflineLoader())
	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return
	}

	// encode
	token := tke.Encode(text, nil, nil)

	//tokens
	// fmt.Println((token))
	// num_tokens
	fmt.Println("total length of tokens. ", len(token))
}
