//go:build ignore

package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

const (
	screenWidth  = 100
	screenHeight = 25
)

func main() {
	// Limpa a tela e posiciona o cursor
	clearScreen()
	showMainMenu()
}

func showMainMenu() {
	// Cores principais do sistema (invertidas e ajustadas)
	headerStyle := color.New(color.BgCyan, color.FgBlack)
	menuStyle := color.New(color.BgBlue, color.FgWhite)
	statusStyle := color.New(color.BgCyan, color.FgBlack)

	// Desenha a tela base
	drawScreen()

	// Cabeçalho
	moveCursor(1, 1)
	headerStyle.Printf("%[1]*s", screenWidth, "")
	moveCursor(1, 2)
	centerText(" 🏧 BANCO 24 HORAS - Sistema de Autoatendimento 🏧           ", headerStyle)
	moveCursor(1, 3)
	headerStyle.Printf("%[1]*s", screenWidth, "")

	// Menu principal (ajustado o espaçamento)
	menuItems := []string{
		"1. [💰 Saque]              - Realizar saque em dinheiro",
		"2. [💵 Saldo]              - Consultar saldo atual",
		"3. [📊 Extrato]            - Visualizar extrato bancário",
		"4. [💸 Transferência]      - Transferir valores",
		"5. [⚡ Pix]                - Realizar transferência Pix",
		"6. [📝 Pagamentos]         - Pagar contas e boletos",
		"7. [📈 Investimentos]      - Acessar área de investimentos",
		"8. [⚙️  Configurações]      - Alterar configurações",
		"0. [🚪 Sair]               - Encerrar sessão",
	}

	for i, item := range menuItems {
		moveCursor(5, 6+i)
		menuStyle.Print(item)
	}

	// Área de informações (movida mais para direita)
	infoStyle := color.New(color.BgWhite, color.FgBlack)
	moveCursor(65, 4)
	infoStyle.Print(" ┌──────────── ℹ️  ────────────┐ ")
	moveCursor(65, 5)
	infoStyle.Print(" │    📱 Informações Úteis     │ ")
	moveCursor(65, 6)
	infoStyle.Print(" ├─────────────────────────────┤ ")
	moveCursor(65, 7)
	infoStyle.Print(" │ 🕒 Horário de atendimento:  │ ")
	moveCursor(65, 8)
	infoStyle.Print(" │ 🏧 24 horas - Todos os dias │ ")
	moveCursor(65, 9)
	infoStyle.Print(" └─────────────────────────────┘ ")

	// Barra de status melhorada
	footerStyle := color.New(color.BgCyan, color.FgBlack)
	moveCursor(1, screenHeight-2)
	footerStyle.Printf("%[1]*s", screenWidth, "")
	moveCursor(2, screenHeight-2)
	helpText := "⌨️  Use o teclado numérico | 🔤 [1-9] = Selecionar | ⏎ Enter = Confirmar | ⌫ Esc = Voltar"
	footerStyle.Print(helpText)

	moveCursor(1, screenHeight-1)
	statusStyle.Printf("%[1]*s", screenWidth, "")
	moveCursor(2, screenHeight-1)
	statusText := fmt.Sprintf("🟢 Online | 🕒 %s | 🔔 F1=Ajuda | 🚪 ESC=Sair",
		time.Now().Format("15:04:05"))
	statusStyle.Print(statusText)

	// Prompt com fundo
	promptStyle := color.New(color.BgBlue, color.FgWhite)
	moveCursor(3, screenHeight-4)
	promptStyle.Printf("%[1]*s", screenWidth-6, "")
	moveCursor(5, screenHeight-4)
	promptStyle.Print("💻 Digite sua opção: _")
}

func drawScreen() {
	bgStyle := color.New(color.BgBlue) // Cor de fundo principal alterada
	for i := 1; i <= screenHeight; i++ {
		moveCursor(1, i)
		bgStyle.Printf("%[1]*s", screenWidth, "")
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J") // Limpa a tela
	fmt.Print("\033[?25l")     // Esconde o cursor
}

func moveCursor(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

func centerText(text string, printer *color.Color) {
	padding := (screenWidth - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	printer.Printf("%[1]*s%s%[1]*s", padding, "", text, "")
}

func init() {
	color.NoColor = false
}
