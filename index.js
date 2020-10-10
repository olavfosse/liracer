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

const games = {}

io.on('connection', socket => {
  let gameID

  const sendLeaveMessage = id => {
    io.to(id).emit('chat message', {
      sender: 'liracer',
      content: 'Player left'
    })
  }

  const sendJoinMessage = id => {
    socket.to(id).emit('chat message', {
      sender: 'liracer',
      content: 'Player joined'
    })
  }

  const sendCurrentSnippetMessage = id => {
    io.to(id).emit('chat message', {
      sender: 'liracer',
      content: `The current snippet is ${games[id].snippet.name}`
    })
  }

  socket.on('disconnecting', () => sendLeaveMessage(gameID))

  socket.on('join game', id => {
    socket.leave(gameID)
    sendLeaveMessage(gameID)

    socket.join(id)
    if(games[id]) {
      sendJoinMessage(id)
    } else {
      // Create game
      games[id] = {
        snippet: randomSnippet()
      }
      sendCurrentSnippetMessage(id)
    }

    gameID = id
    socket.emit('code snippet', games[id].snippet.code)
  })

  socket.on('cursor position update', position => {
    const game = games[gameID]
    if(game && position === game.snippet.code.length) {
      game.snippet = randomSnippet()
      sendCurrentSnippetMessage(gameID)
      io.to(gameID).emit('code snippet', games[gameID].snippet.code)
    }
  })
})

console.log(`listening on ${port}`)
server.listen(port)