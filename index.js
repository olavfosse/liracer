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

io.on('connection', socket => {
  socket.emit('code snippet', {
    language: 'ruby',
    code:
    "def fibonacci(n)\n" +
	  "\tn <= 1 ? n : fibonacci(n-1) + fibonacci(n-2)\n" +
    "end\n" +
    "\n" +
    "puts fibonacci(gets.to_i)"
  })

  socket.emit('chat message', {
    sender: 'liracer',
    content: 'Welcome to liracer! Click "JOIN GAME" and enter a GameID, or type "/join GameID" to join a game. If a game by the given GameID exists you join that, otherwise a new game is created.'
  })
})

console.log(`listening on ${port}`)
server.listen(port)