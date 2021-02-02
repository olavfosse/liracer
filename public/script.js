/*
 * CONSTANTS
 */
const codefield = document.getElementsByClassName("codefield")[0]
const snip = `
package main

import "fmt"

func main() {
	fmt.Println("hello, world!")
}`.trim()

/*
 * GAME STATE
 */
let correctChars = 5
let incorrectChars = 10

// renderSnippet renders a snippet snip to the codefield, overwriting any previous text.
const renderSnippet = (snip, correctChars, incorrectChars) => {
	codefield.textContent = ""
	snip.split("").forEach((c, i) => {
		s = document.createElement("span")
		if (c === "\n") {
			s.textContent = "↵\n"
			s.classList.add("codefield-character-newline")
		} else {
			s.textContent = c
		}
		s.classList.add("codefield-character")
		if (i < correctChars) {
			s.classList.add("codefield-character-correct")
		} else if (i < correctChars + incorrectChars) {
			s.classList.add("codefield-character-incorrect")
		} else if (i === correctChars + incorrectChars) {
			s.classList.add("codefield-character-player")
		}

		codefield.appendChild(s)
	})
}

renderSnippet(snip, correctChars, incorrectChars)

codefield.addEventListener("keydown", e => {
	if (incorrectChars === 0) {
		correctChars = e.key === "Backspace" ? correctChars - 1 : correctChars + 1
	} else {
		incorrectChars = e.key === "Backspace" ? Math.max(0, incorrectChars - 1) : incorrectChars + 1
	}

	renderSnippet(snip, correctChars, incorrectChars)
})
