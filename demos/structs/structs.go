package main

import (
	"github.com/fatih/color"
)

type Product struct {
	name, category string
	price          float64
}

func main() {
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}


	color.White("\n🛒 Product Description") // Unicode: U+1F6D2
	color.Green("✅ Name: %s Category: %s Price: %.2f\n\n", kayak.name, kayak.category, kayak.price) // Unicode: U+2705

	color.New(color.BlinkRapid, color.Italic, color.FgCyan).Println("🔄 Changing price...\n") // Unicode: U+1F504

	kayak.price = 300

	color.Red("🔴 Changed Price!!!\n") // Unicode: U+1F534

	color.White("\n🛒 Product Description") // Unicode: U+1F6D2
	color.HiBlue("✅ Name: %s Category: %s Price: %.2f\n\n", kayak.name, kayak.category, kayak.price) // Unicode: U+2705

	color.HiWhite("💾💿💡📃📒💰📂📈📌⌛🚨❌❓❔❗❕✅ 0️⃣ 1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣ 8️⃣ 9️⃣ 🔟🔴🟢🔵🟡⚪💬💭")

	color.Red("")
}