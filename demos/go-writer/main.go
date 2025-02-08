package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// TextSection representa uma seção individual do texto
type TextSection struct {
	ID       string
	Title    string
	Content  string
	Order    int
	FilePath string
}

// Document representa o documento completo
type Document struct {
	Title      string
	Sections   []TextSection
	OutputPath string
}

func main() {
	doc := &Document{
		Title:      "Novo Documento",
		OutputPath: "output.md",
	}

	if err := initializeWorkspace(); err != nil {
		fmt.Printf("Erro ao inicializar workspace: %v\n", err)
		os.Exit(1)
	}

	// Configurar rotas para a interface web
	http.HandleFunc("/", handleIndex(doc))
	http.HandleFunc("/api/sections", handleSections(doc))

	// Servir arquivos estáticos
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Servidor rodando em http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v\n", err)
		os.Exit(1)
	}
}

func handleIndex(doc *Document) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, doc)
	}
}

func handleSections(doc *Document) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implementar API para gerenciar seções
	}
}

// initializeWorkspace cria a estrutura de diretórios necessária
func initializeWorkspace() error {
	dirs := []string{
		"sections", // para armazenar seções individuais
		"output",   // para o arquivo final combinado
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("erro ao criar diretório %s: %v", dir, err)
		}
	}
	return nil
}

// SaveDocument salva o documento completo
func (d *Document) SaveDocument() error {
	// TODO: Implementar a lógica para combinar todas as seções
	// e gerar o arquivo final com TOC
	return nil
}

// AddSection adiciona uma nova seção ao documento
func (d *Document) AddSection(title, content string) error {
	section := TextSection{
		ID:      generateID(),
		Title:   title,
		Content: content,
		Order:   len(d.Sections),
	}

	// Criar arquivo para a seção
	filename := filepath.Join("sections", fmt.Sprintf("%s.md", section.ID))
	section.FilePath = filename

	// TODO: Implementar salvamento do conteúdo no arquivo

	d.Sections = append(d.Sections, section)
	return nil
}

// ReorderSections reordena as seções do documento
func (d *Document) ReorderSections(newOrder []string) error {
	// TODO: Implementar reordenação das seções
	return nil
}

// generateID gera um ID único usando UUID
func generateID() string {
	return uuid.New().String()
}
