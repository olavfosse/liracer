const { clear } = require('console')
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
  roundID: createPseudoRandomString(),
  startingTime: new Date().getTime()
})

io.on('connection', socket => {
  let gameID

  const sendAnonLeftMessage = id => {
    io.to(id).emit('liracer message', 'anon left')
  }

  const sendAnonJoinedMessage = id => {
    io.to(id).emit('liracer message', 'anon joined')
  }

  const sendGameCreatedMessage = id => {
    io.to(id).emit('liracer message', 'Game created')
  }

  const sendCurrentSnippetMessage = id => {
    io.to(id).emit('liracer message', `The current snippet is ${games[id].snippet.name}`)
  }

  const sendSnippetCompletedMessage = id => {
    const timeToComplete = (new Date().getTime() - games[id].startingTime) / 1000
    io.to(id).emit('liracer message', `${games[id].snippet.name} completed in ${timeToComplete} seconds`)
  }

  const clearCursor = id => {
    // This is kind of hackish
    // If position is -1 it won't be rendered, and it will be cleared from front end memory on 'game state' message
    socket.to(id).emit('cursor position update', {
      sid: socket.id,
      position: -1
    })
  }

  socket.on('disconnecting', () => {
    clearCursor(gameID)
    sendAnonLeftMessage(gameID)
  })

  socket.on('join game', id => {
    socket.leave(gameID)
    sendAnonLeftMessage(gameID)

    socket.join(id)
    if(games[id]) {
      sendAnonJoinedMessage(id)
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
    if(games[gameID].roundID !== roundID) { // See the wall of text above :^)
      return
    }

    if(position === games[gameID].snippet.code.length) {
      // New game
      sendSnippetCompletedMessage(gameID)
      games[gameID] = createGame()
      sendCurrentSnippetMessage(gameID)
      io.to(gameID).emit('game state', games[gameID])
    } else {
      // Send cursor position update to other players
      socket.to(gameID).emit('cursor position update', {
        sid: socket.id,
        position
      })
    }
  })

  socket.on('message', (content) => {
    io.to(gameID).emit('anon message', {
      content,
      playerID: socket.id
    })
  })
})

console.log(`listening on ${port}`)
server.listen(port)