const express = require('express')
const app = express()
const server = require('http').createServer(app)
const options = {
  serveClient: false
}
const io = require('socket.io')(server, options)

const port = process.env.PORT || 3101
const buildPath = `${__dirname}/front/build`

app.use(express.static(buildPath))
app.get('/*', (_request, response) => {
  response.sendFile(`${buildPath}/index.html`)
})

const quotes = [
  {
    name: 'fibonacci.rb',
    code:
    "def fibonacci(n)\n" +
	  "\tn <= 1 ? n : fibonacci(n-1) + fibonacci(n-2)\n" +
    "end\n" +
    "\n" +
    "puts fibonacci(gets.to_i)"
  },
  {
    name: 'hello_world.rb',
    code: "puts 'hello, world'"
  }
]
const randomQuote = () => quotes[Math.floor(quotes.length * Math.random())]

io.on('connection', socket => {
  let quote
  const newQuote = () => {
    quote = randomQuote()
    socket.emit('code snippet', quote.code)
    socket.emit('chat message', {
      sender: 'liracer',
      content: `The current quote is ${quote.name}`
    })
  }

  socket.emit('chat message', {
    sender: 'liracer',
    content: 'Welcome to liracer! Click "JOIN GAME" and enter a GameID, or type "/join GameID" to join a game. If a game by the given GameID exists you join that, otherwise a new game is created.'
  })

  newQuote()

  socket.on('cursor position update', position => {
    console.log(`client on position ${position}`)

    if(quote.code.length === position) {
      newQuote()
    }
  })
})

console.log(`listening on ${port}`)
server.listen(port)