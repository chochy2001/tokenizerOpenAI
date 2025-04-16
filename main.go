package main

import (
	"fmt"
	"github.com/pkoukk/tiktoken-go"
)

func main() {
	encoding := "cl100k_base"

	// if you don't want download dictionary at runtime, you can use offline loader
	// tiktoken.SetBpeLoader(tiktoken_loader.NewOfflineLoader())
	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return
	}

	// encode
	token := tke.Encode(textToTokenize(), nil, nil)

	fmt.Println("total length of tokens. ", len(token))
}
