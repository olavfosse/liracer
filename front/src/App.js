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

// Use a grid layout
display: grid;
grid-template-columns: minmax(300px, 500px) minmax(750px, auto);
padding: 2rem;
column-gap: 2rem;

// Color
background: ${colors.layer0Background}
`

function App() {
  const [code, setCode] = useState()
  const [cursorPosition, setCursorPosition] = useState()
  const [wrongChars, setWrongChars] = useState()
  const [messages, setMessages] = useState([])
  const [socket, setSocket] = useState()

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

    socket.on('code snippet', ({ language, code }) => {
      setCode(code)
      setCursorPosition(0)
      setWrongChars(0)
    })

    socket.on('chat message', message => {
      setMessages(messages => [...messages, message])
    })
  }, [socket])

  useEffect(() => {
    if(!socket || cursorPosition === undefined) {
      return
    }

    socket.emit('cursor position update', cursorPosition)
  }, [socket, cursorPosition])

  return isMobile(window.navigator).any ? (
    <div>
      <h3>This game is not playable on mobile devices</h3>
      <p>To play liracer, open it on a laptop or desktop computer.</p>
    </div>
  ) : (
    <Grid>
      <ChatAndJoinButton messages={ messages } />
      <CodeField code={ code }
                 cursorPosition={cursorPosition}
                 setCursorPosition={setCursorPosition}
                 wrongChars={wrongChars}
                 setWrongChars={setWrongChars} />
    </Grid>
  ) 
}

export default App
