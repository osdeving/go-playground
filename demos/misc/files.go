//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Para abrir arquivos podemos usar a funçao os.Create
	// CUIDADO: isso vai truncar o arquivo se ele já existir, ou seja, zerar ele
	// e então deixar disponível para escrita.
	file, _ := os.Create("myfile.txt")
	defer file.Close()

	// Podemos escrever usando bytes com a função myfile.Write
	len1, _ := file.Write([]byte("Escrevendo dados binários no arquivo\n"))
	
	// ou podemos escrever usando strings com a função myfile.WriteString
	len2, _ := file.WriteString("Testando escrita com WriteString\n")

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", len1+len2)

	// ReadFile le todo o arquivo e retorna os bytes
	// Essa função abre o arquivo, le todo conteúdo e fecha o arquivo
	var fileContent []byte
	fileContent, _ = os.ReadFile("myfile.txt")

	fmt.Printf("\nConteudo do arquivo: \n%s\n", fileContent)


	// Lendo um arquivo em pedaços com bufio Reader.
	// Cenário:
	// O buffer interno tem DefaultBufSize que é 4096 (obtido via reader.Size())
	// Escolhemos um buffer nosso para receber o conteúdo
	// E o próprio arquivo tem uma quantidade N de bytes.
	//
	// Exemplo: Arquivo com 12288 bytes (3 * DefaultBufSize que é 4096)
	// O bufio.Reader tem um buffer interno de 4096 bytes (DefaultBufSize)
	// Nosso chunk de leitura é de 10 bytes
	// 
	// Fluxo:
	// 1. Primeira chamada de sistema: carrega 4096 bytes no buffer interno
	// 2. A cada Read(chunk), lê 10 bytes do buffer
	// 3. Após ~410 leituras (4096/10), o buffer interno esvazia
	// 4. Nova chamada de sistema carrega mais 4096 bytes
	// 5. O processo se repete até ler todo arquivo (3 chamadas de sistema no total)
	fmt.Printf("\n\nLendo o arquivo em chunks\n\n")

	file2, _ := os.Open("myfile.txt")
	defer file2.Close()

	reader := bufio.NewReader(file2)
	chunk := make([]byte, 3)
	fileInfo, _ := file2.Stat()

	fmt.Printf("Tamanho do buffer interno: %d\nTamanho do nosso buffer: %d\nTamanho do arquivo: %d\n\n", 
				reader.Size(), len(chunk), fileInfo.Size())
	
	for {
		n, err := reader.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		fmt.Println(string(chunk[:n]))
	}
}
