/* ========= *
 * CONSTANTS *
 * ========= */
const codefield = document.getElementsByClassName("codefield")[0]
let wsProtocol = 'ws:'
if(location.protocol === 'https:') {
    wsProtocol = 'wss:'
}
const socket = new WebSocket(`${wsProtocol}//${document.location.host}/ws`)

/* ========== *
 * ROOM STATE *
 * ========== */
let snippet = undefined
let opponentCorrectChars = {
	// ID: correctChars
}
let correctChars = 0
let incorrectChars = 0

/* =========== *
 * OTHER STATE *
 * =========== */
let roomID = undefined
let roundID = undefined

/* ========= *
 * FUNCTIONS *
 * ========= */
// renderSnippet renders a snippet snip to the codefield, overwriting any
// previous text.
const renderCodefield = () => {
	codefield.textContent = ""
	snippet.split("").forEach((c, i) => {
		const s = document.createElement("span")
		s.textContent = c === "\n" ? s.textContent = "↵\n" : c

		if (i < correctChars) {
			s.style.setProperty('background', '#c5ddc5')
		}

		Object.values(opponentCorrectChars).forEach(correctChars => {
			if (i === correctChars) {
				s.style.setProperty('background', '#baba70')
			}
		})

		if (i >= correctChars && i < correctChars + incorrectChars) {
			s.style.setProperty('background', '#dec5c5')
		}

		if (i === correctChars + incorrectChars) {
			s.style.setProperty('background', '#cebaa6')
		}

		codefield.appendChild(s)
	})
}

// send sends a JSON representation of obj to the server and logs it to the
// console.
const send = obj => {
	const s = JSON.stringify(obj)
	console.log('write: ' + s)
	socket.send(s)
}

// sendCorrectChars sends the number of correctly written characters, that is
// correctChars, to the server.
const sendCorrectChars = () => {
	send({
		'RoundID': roundID,
		'CorrectCharsMsg': {
			'CorrectChars': correctChars,
			'RoundId': roundID,
		}
	})
}

// typeIncorrectChar "types" a incorrect character, that is increments
// incorrectChars and renders codefield.
const typeIncorrectChar = () => {
	incorrectChars++
	renderCodefield()
}

// deleteIncorrectChar "deletes" a incorrect character, that is it decrements
// incorrectChars and renders the codefield.
const deleteIncorrectChar = () => {
	incorrectChars--
	renderCodefield()
}

// typeCorrectChar "types" a correct character, that is increments correctChars,
// sends the updated correctChars to the server and renders the codefield.
const typeCorrectChar = () => {
	correctChars++
	sendCorrectChars()
	renderCodefield()
}

// deleteCorrectChar "deletes" a correct character, that is it decrements
// correctChars, sends the updated correctChars to the server correctChars and
// renders the codefield.
const deleteCorrectChar = () => {
	correctChars--
	sendCorrectChars()
	renderCodefield()
}

// restart restarts the room state.
const restart = () => {
	correctChars = 0
	incorrectChars = 0
	opponentCorrectChars = {}
	sendCorrectChars()
	renderCodefield()
}

// mapKeyToChar maps a key, as in the key field of a KeyboardEvent, to the
// character it represents.
const mapKeyToChar = key => {
	console.log(key)
	if(['Shift', 'Meta', 'Alt', 'AltGraph', 'Control', 'Backspace'].includes(key)){
		return null
	} else if (key === "Enter"){
		return "\n"
	} else if (key === 'Tab') {
		return "\t"
	} else {
		return key
	}
}

/* =========== *
 * ENTRY POINT *
 * =========== */
// renderCodefield(snip, correctChars, incorrectChars)

// TODO: wait until socket connection is opened before registering event
codefield.addEventListener("keydown", e => {
	if (e.key === "Tab") {
		// WHY: On Safari, pressing tab makes the browser focus the search
		//      field. calling e.preventDefault prevents this.
		e.preventDefault()
	}

	if(e.key === "Backspace") {
		if (incorrectChars > 0) {
			deleteIncorrectChar()
		} else if (correctChars > 0) {
			deleteCorrectChar()
		}
		return
	}

	const char = mapKeyToChar(e.key)
	if(char === null) {
		return
	}

	if(incorrectChars > 0) {
		typeIncorrectChar()
		return
	}

	if(char === snippet[correctChars]) {
		typeCorrectChar()
	} else {
		typeIncorrectChar()
	}

	if(correctChars === snippet.length) {
		restart()
	}
})

socket.addEventListener('message', e => {
	console.log("read: " + e.data)
	const m = JSON.parse(e.data)

	let isMessageHandled = false

	if(m['NewRoundMsg'] !== null) {
		isMessageHandled = true
		const payload = m['NewRoundMsg']

		if(roundID === undefined || payload['RoundId'] === roundID) {
			snippet = payload['Snippet']
			roundID = payload['NewRoundId']
			correctChars = 0
			incorrectChars = 0
			opponentCorrectChars = {}
			renderCodefield()
		}
	}
	if(m['OpponentCorrectCharsMsg'] !== null) {
		isMessageHandled = true
		const payload = m['OpponentCorrectCharsMsg']

		if(roundID === undefined || payload['RoundId'] === roundID) {
			opponentCorrectChars[payload['OpponentID']] = payload['CorrectChars']
			renderCodefield()
		}
	}
	if(m['ChatMessageMsg'] !== null) {
		isMessageHandled = true
		const payload = m['ChatMessageMsg']

		const chatMessage = document.createElement('div')
		chatMessage.className = 'chat-message'

		const chatMessageSender = document.createElement('span')
		chatMessageSender.className = 'chat-message-sender'
		chatMessageSender.textContent = '<' +payload.Sender +'>'
		chatMessage.appendChild(chatMessageSender)

		const chatMessageContent = document.createElement('span')
		chatMessageContent.className = 'chat-message-content'
		chatMessageContent.textContent = payload.Content
		chatMessage.appendChild(chatMessageContent)

		// When the user is scrolled all the way down in .chat-messages the chat
		// automatically scrolls down when new messages are sent so that (s)he does
		// not have to manually scroll down to see the latest message. If however
		// the user scrolls up from the bottom, to look at previous messages, the
		// chat does not automatically scroll down. 
		const chatMessages = document.getElementsByClassName('chat-messages')[0]
		const isChatScrolledAllTheWayDown = Math.abs(chatMessages.scrollTop+chatMessages.offsetHeight-chatMessages.scrollHeight) <= 4
		chatMessages.append(chatMessage)
		if(isChatScrolledAllTheWayDown) chatMessages.scrollTop = chatMessages.scrollHeight
	}
	if(!isMessageHandled) {
		alert('unhandled message: ' + e.data)
	}
})


const chatForm = document.getElementsByClassName('chat-form')[0]
const chatFormTextField = document.getElementsByClassName('chat-form-text-field')[0]
chatForm.addEventListener('submit', event => {
	event.preventDefault()
	send({'ChatMessageMsg': {'Content': chatFormTextField.value}})
	chatFormTextField.value = ''
})
