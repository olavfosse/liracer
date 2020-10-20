import React from 'react'
import { useState, useEffect } from 'react'
import styled from 'styled-components'
import io from 'socket.io-client'
import isMobile from 'ismobilejs'
import ChatAndJoinButton from './components/ChatAndJoinButton'
import CodeField from './components/CodeField'
import colors from './colors'
const Grid = styled.div`
// Consume the entire viewport
position: fixed;
top: 0;
left: 0;
bottom: 0;
right: 0;
overflow: auto;

// Makes Pre take up the proper amount of space on Safari
// DON'T TOUCH
height: 100%;
box-sizing: border-box;
// DON'T TOUCH

// Use a grid layout
display: grid;
grid-template-columns: minmax(300px, 500px) minmax(750px, auto);
padding: 2rem;
column-gap: 2rem;

// Color
background: ${colors.layer0Background}
`

function App() {
  const [snippet, setSnippet] = useState()
  const [roundID, setRoundID] = useState()
  const [cursorPosition, setCursorPosition] = useState()
  const [opponentCursorPositions, setOpponentCursorPositions] = useState({})
  const [wrongChars, setWrongChars] = useState()
  const [messages, setMessages] = useState([]) // message content, sender name and optionally playerID
  const [socket, setSocket] = useState()

  const handleClickJoinGame = _ => {
    const gameID = prompt('GameID')
    if(gameID !== null && gameID !== '') {
      socket.emit('join game', gameID)
      history.pushState(undefined, undefined, gameID) // eslint-disable-line no-restricted-globals
    }
  }

  const handleSendMessage = event => {
    event.preventDefault()

    socket.emit('message', event.target.input.value)
    event.target.input.value = ''
  }

  const parseGameIDFromLocation = location => decodeURI(location.pathname).slice(1)

  window.onpopstate = event => {
    const gameID = parseGameIDFromLocation(event.target.location)
    socket.emit('join game', gameID)
    history.replaceState(undefined, undefined, gameID) // eslint-disable-line no-restricted-globals
  }

  const createPseudoRandomString = _ => Math.random().toString(36).replace(/[^a-z]+/g, '')
  useEffect(() => {
    if(!socket){
      return
    }

    if(window.location.pathname === '/') {
      const gameID = createPseudoRandomString()
      socket.emit('join game', gameID)
      history.replaceState(undefined, undefined, gameID) // eslint-disable-line no-restricted-globals
    } else {
      const gameID = parseGameIDFromLocation(window.location)
      socket.emit('join game', gameID)
      history.pushState(undefined, undefined, gameID) // eslint-disable-line no-restricted-globals
    }
  }, [socket])

  useEffect(() => {
    if(process.env.NODE_ENV !== 'production') {
      setSocket(io('http://localhost:3101'))
    } else {
      setSocket(io())
    }
  }, [])

  useEffect(() => {
    if(!socket) {
      return
    }

    socket.on('game state', (game) => {
      setSnippet(game.snippet)
      setWrongChars(0)
      setOpponentCursorPositions({})
      setCursorPosition(0)

      // It is crucial that the roundID is updated after the cursor is set to 0
      setRoundID(game.roundID)
    })


    socket.on('liracer message', (content) => {
      setMessages(messages => [...messages, { sender: 'liracer', content }])
    })

    socket.on('anon message', ({playerID, content}) => {
      setMessages(messages => [...messages, { sender: 'anon', content, playerID }])
    })

    socket.on('cursor position update', ({
      sid, // socket id, identifies the player/client
      position
    }) => {
      setOpponentCursorPositions(opponentCursorPositions => ({...opponentCursorPositions, [sid]: position }))
    })
  }, [socket])

  useEffect(() => {
    if(!socket || cursorPosition === undefined) {
      return
    }

    socket.emit('cursor position update', {
      position: cursorPosition,
      roundID
    })
  }, [roundID, socket, cursorPosition])

  return isMobile(window.navigator).any ? (
    <div>
      <h3>This game is not playable on mobile devices</h3>
      <p>To play liracer, open it on a laptop or desktop computer.</p>
    </div>
  ) : (
    <Grid>
      <ChatAndJoinButton messages={ messages }
                         handleClickJoinGame={ handleClickJoinGame }
                         handleSendMessage={handleSendMessage}/>
      <CodeField snippet={ snippet }
                 cursorPosition={cursorPosition}
                 opponentCursorPositions={opponentCursorPositions}
                 setCursorPosition={setCursorPosition}
                 wrongChars={wrongChars}
                 setWrongChars={setWrongChars} />
    </Grid>
  ) 
}

export default App
