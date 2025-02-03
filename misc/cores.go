package main

import (
    "github.com/fatih/color"
)

func main() {
    // exemplos de como usar fatih color
	color.Red("We have red")
	color.Yellow("We have yellow")
	color.Blue("We have blue")
	color.Green("We have green")
	color.Magenta("We have magenta")
	color.Cyan("We have cyan")
	color.White("We have white")
	color.Black("We have black")
	color.HiRed("We have hi-red")
	color.HiYellow("We have hi-yellow")
	color.HiBlue("We have hi-blue")
	color.HiGreen("We have hi-green")
	color.HiMagenta("We have hi-magenta")
	color.HiCyan("We have hi-cyan")
	color.HiWhite("We have hi-white")
	color.HiBlack("We have hi-black")
	color.HiBlackString("We have hi-black string")
	color.HiRedString("We have hi-red string")
	color.HiYellowString("We have hi-yellow string")
	color.HiBlueString("We have hi-blue string")
	color.HiGreenString("We have hi-green string")
	color.HiMagentaString("We have hi-magenta string")
	color.HiCyanString("We have hi-cyan string")
	color.HiWhiteString("We have hi-white string")
	color.RedString("We have red string")
	color.YellowString("We have yellow string")
	color.BlueString("We have blue string")
	color.GreenString("We have green string")
	color.MagentaString("We have magenta string")
	color.CyanString("We have cyan string")
	color.WhiteString("We have white string")
	color.BlackString("We have black string")
	color.New(color.FgCyan).Println("Prints text in cyan.")
	color.New(color.BgWhite).Println("Prints text with white background.")
	color.New(color.BgHiWhite).Println("Prints text with hi-white background.")
	color.New(color.BgHiWhite, color.FgBlack).Println("Prints text with hi-white background and black foreground.")
	color.New(color.BgHiWhite, color.FgBlack, color.Bold, color.BlinkRapid).Println("Prints text with hi-white background, black foreground, bold and blinking.")
	
}