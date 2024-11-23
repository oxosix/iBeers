package art

import (
	"fmt"
	"log"

	"github.com/mbndr/figlet4go"
)

func AsciiArt(s string) (r string) {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
	}
	text := s
	renderedText, err := ascii.RenderOpts(text, options)
	if err != nil {
		log.Fatalf("Erro ao renderizar o texto: %v", err)
	}
	fmt.Print(renderedText)
	return
}
