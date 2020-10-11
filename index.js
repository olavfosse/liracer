const express = require('express')
const app = express()
const server = require('http').createServer(app)
const options = {
  serveClient: false
}
const io = require('socket.io')(server, options)

// The frontend assumes that the backend is on same port as backend in production and on port 3101 otherwise
const port = process.env.PORT || 3101
const buildPath = `${__dirname}/front/build`

app.use(express.static(buildPath))
app.get('/*', (_request, response) => {
  response.sendFile(`${buildPath}/index.html`)
})

const snippets = require('./snippets.js')
const randomSnippet = () => snippets[Math.floor(snippets.length * Math.random())]

const games = {}

const createPseudoRandomString = _ => Math.random().toString(36).replace(/[^a-z]+/g, '')
const createGame = _ => ({
  snippet: randomSnippet(),
  roundID: createPseudoRandomString()
})

io.on('connection', socket => {
  let gameID

  const sendPlayerLeftMessage = id => {
    io.to(id).emit('chat message', {
      sender: 'liracer',
      content: 'Player left'
    })
  }

  const sendPlayerJoinedMessage = id => {
    socket.to(id).emit('chat message', {
      sender: 'liracer',
      content: 'Player joined'
    })
  }

  const sendGameCreatedMessage = id => {
    io.to(id).emit('chat message', {
      sender: 'liracer',
      content: 'Game created'
    })
  }

  const sendCurrentSnippetMessage = id => {
    io.to(id).emit('chat message', {
      sender: 'liracer',
      content: `The current snippet is ${games[id].snippet.name}`
    })
  }

  socket.on('disconnecting', () => sendPlayerLeftMessage(gameID))

  socket.on('join game', id => {
    socket.leave(gameID)
    sendPlayerLeftMessage(gameID)

    socket.join(id)
    if(games[id]) {
      sendPlayerJoinedMessage(id)
    } else {
      sendGameCreatedMessage(id)
      games[id] = createGame()
      sendCurrentSnippetMessage(id)
    }

    gameID = id
    socket.emit('game state', games[id])
  })

  // TL;DR: roundID is used to verify that the received 'cursor position update' message refers to the current round
  // roundID is a random hash used to identify a round
  // a new round is started each time a new snippet is used
  // this is used so that messages can be invalidated if they contain an outdated roundID
  // this prevents the following bug:
  //   user1 sends a 'cursor position update' message causing a 'code snippet' message
  //   user2, which has not yet received the 'code snippet' message, sends a 'cursor position update' message.
  //   since user2 has not yet received the 'cursor position update' message the position had not been reset to 0.
  //   therefore the position sent is representative of how much of new 'code snippet' user2 has actually typed.
  //   worst case scenario the inaccurate position user2 sent matches the length of the new code snippet causing user2 to instantly win the round.
  socket.on('cursor position update', ({ position, roundID }) => {
    const game = games[gameID]
    if(game.roundID !== roundID) { // See the wall of text above :^)
      return
    }

    if(position === game.snippet.code.length) {
      game.snippet = randomSnippet()
      game.roundID = createPseudoRandomString()
      sendCurrentSnippetMessage(gameID)
      io.to(gameID).emit('game state', game)
    }
  })
})

console.log(`listening on ${port}`)
server.listen(port)