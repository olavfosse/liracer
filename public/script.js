/* ========= *
 * CONSTANTS *
 * ========= */
const codefield = document.getElementsByClassName("codefield")[0]
const snip = `
package main

import "fmt"

func main() {
	fmt.Println("hello, world!")
}`.trim()

/* ========== *
 * GAME STATE *
 * ========== */
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

// mapKeyToChar maps a key, as in the key field of a KeyboardEvent, to the character it represents.
const mapKeyToChar = key => {
	if(['Shift', 'Meta', 'Alt', 'Control', 'Backspace'].includes(key)){
		return null
	} else if (key === "Enter"){
		return "\n"
	} else if (key === 'Tab') {
		return "\t"
	} else {
		return key
	}
}

renderSnippet(snip, correctChars, incorrectChars)

codefield.addEventListener("keydown", e => {
	if (e.key === "Tab") {
		/* On Safari, pressing tab makes the browser focus the search field. calling e.preventDefault prevents this. */
		e.preventDefault()
	}

	if(e.key === "Backspace") {
		if (incorrectChars > 0) {
			incorrectChars--
		} else if (correctChars > 0) {
			correctChars--
		}
		renderSnippet(snip, correctChars, incorrectChars)
		return
	}

	char = mapKeyToChar(e.key)
	if(char === null) {
		return
	}

	if(incorrectChars > 0) {
		incorrectChars++
		renderSnippet(snip, correctChars, incorrectChars)
		return
	}

	if(char === snip[correctChars]) {
		correctChars++
	} else {
		incorrectChars++
	}
	renderSnippet(snip, correctChars, incorrectChars)
})
