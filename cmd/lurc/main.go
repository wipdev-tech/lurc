package main

import (
	"fmt"

	l "github.com/charmbracelet/lipgloss"
)

var styles = struct {
	header l.Style
}{
	header: l.NewStyle().Background(l.Color("56")).Bold(true),
}

func main() {
	fmt.Printf("\n%s\n\n", styles.header.Render(" lurc "))
}
