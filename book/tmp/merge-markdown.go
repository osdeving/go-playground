package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func readFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Substitui os links por referências internas
func processMarkdown(mainFile string) (string, error) {
	mainContent, err := readFile(mainFile)
	if err != nil {
		return "", fmt.Errorf("Erro ao ler o arquivo principal: %w", err)
	}

	// Expressão regular para encontrar links do tipo [Título](section-x.x.md)
	//pattern := regexp.MustCompile(`\[(.*?)\]\((section-\d+\.\d+\.md)\)`)
	pattern := regexp.MustCompile(`\[(.*?)\]\(chapters/chapter-\d+/ch\d+-section-\d+\.\d+\.md\)`)
	
	sections := make(map[string]string)
	var orderedKeys []string

	finalContent := pattern.ReplaceAllStringFunc(mainContent, func(match string) string {
		matches := pattern.FindStringSubmatch(match)
		if len(matches) != 3 {
			return match // Caso a regex não capture corretamente, mantém o texto original
		}

		title := matches[1]
		file := matches[2]
		refID := strings.ReplaceAll(strings.ToLower(title), " ", "-") // Criando um ID único

		// Verifica se o conteúdo já foi carregado
		if _, exists := sections[refID]; !exists {
			fileContent, err := readFile(file)
			if err != nil {
				fmt.Printf("⚠️ Erro ao ler %s: %v\n", file, err)
				return fmt.Sprintf("[Erro ao carregar %s](#%s)", file, refID)
			}

			// Adiciona ao mapa e mantém a ordem
			sections[refID] = fmt.Sprintf("\n---\n\n## %s {#%s}\n\n%s\n", title, refID, fileContent)
			orderedKeys = append(orderedKeys, refID)
		}

		// Retorna a referência interna para o conteúdo
		return fmt.Sprintf("[%s](#%s)", title, refID)
	})

	// Adiciona o conteúdo das seções ao final do documento
	var sb strings.Builder
	sb.WriteString(finalContent)
	sb.WriteString("\n\n---\n\n") // Separação do conteúdo original e das seções

	for _, key := range orderedKeys {
		sb.WriteString(sections[key])
	}

	return sb.String(), nil
}

func main() {
	mainFile := "../go-bible.md"
	outputFile := "../go-bible-full.md"

	finalMd, err := processMarkdown(mainFile)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	err = os.WriteFile(outputFile, []byte(finalMd), 0644)
	if err != nil {
		fmt.Println("Erro ao salvar arquivo:", err)
		return
	}

	fmt.Println("✅ Markdown consolidado gerado com sucesso:", outputFile)
}
