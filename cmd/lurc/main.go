package main

import (
	"fmt"

	l "github.com/charmbracelet/lipgloss"
)

var styles = struct {
	header l.Style
	prompt l.Style
}{
	header: l.NewStyle().Background(l.Color("56")).Bold(true),
	prompt: l.NewStyle().Foreground(l.Color("248")),
}

func main() {
	fmt.Printf("\n%s\n\n", styles.header.Render(" lurc "))

	url := ""
	fmt.Printf(styles.prompt.Render("URL: "))
	fmt.Scan(&url)

	fmt.Println(url)
}
