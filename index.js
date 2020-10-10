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

const snippets = require('./snippets.js')
const randomSnippet = () => snippets[Math.floor(snippets.length * Math.random())]

io.on('connection', socket => {
  let snippet
  const useNewSnippet = () => {
    snippet = randomSnippet()
    socket.emit('code snippet', snippet.code)
    socket.emit('chat message', {
      sender: 'liracer',
      content: `The current snippet is ${snippet.name}`
    })
  }

  socket.emit('chat message', {
    sender: 'liracer',
    content: 'Welcome to liracer! Click "JOIN GAME" and enter a GameID, or type "/join GameID" to join a game. If a game by the given GameID exists you join that, otherwise a new game is created.'
  })

  useNewSnippet()

  socket.on('cursor position update', position => {
    console.log(`client on position ${position}`)

    if(snippet.code.length === position) {
      useNewSnippet()
    }
  })
})

console.log(`listening on ${port}`)
server.listen(port)