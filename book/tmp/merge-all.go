//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run merge-all.go <book-summary.md> <book-full.md")
		os.Exit(1)
	}
	bookSummary := os.Args[1]
	bookFull := os.Args[2]

	// Lê o conteúdo completo do arquivo principal
	data, err := os.ReadFile(bookSummary)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo %s: %v", bookSummary, err)
	}
	summaryContent := string(data)

	// Expressão regular para encontrar links no formato [Título](caminho)
	linkRegex := regexp.MustCompile(`\[(?P<title>[^\]]+)\]\((?P<path>[^)]+)\)`)
	matches := linkRegex.FindAllStringSubmatch(summaryContent, -1)

	var fullContent string

	for _, match := range matches {
		sectionTitle := match[1]
		sectionPath := match[2]

		fullContent += fmt.Sprintf("## %s\n\n", sectionTitle)
		
		sectionData, err := os.ReadFile(sectionPath)
		if err != nil {
			fullContent += "_Esta seção ainda falta ser escrita._\n\n"
		} else {
			fullContent += string(sectionData) + "\n\n"
		}
	}

	appendContent(bookFull, summaryContent, fullContent)
	
	fmt.Println("Arquivo atualizado com as seções extraídas.")
}

func appendContent(bookFull, summaryContent, content string) {
	f, err := os.OpenFile(bookFull, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo %s para append: %v", bookFull, err)
	}
	defer f.Close()

	
	if _, err := f.WriteString(summaryContent); err != nil {
		log.Fatalf("Erro ao escrever no arquivo %s: %v", bookFull, err)
	}

	if _, err := f.WriteString(content); err != nil {
		log.Fatalf("Erro ao escrever no arquivo %s: %v", bookFull, err)
	}
}
