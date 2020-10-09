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

const dummyCode = `\
def fibonacci(n)
	n <= 1 ? n : fibonacci(n-1) + fibonacci(n-2)
end

puts fibonacci(gets.to_i)
`

const dummyMessages = [
  {
    sender: 'liracer',
    content: 'Click the JOIN button or type "/join GameID" to join a game.'
  },
  {
    sender: 'liracer',
    content: 'Click the JOIN button or type "/join GameID" to join a game.'
  },
  {
    sender: 'fossegrim',
    content: 'Another sample message'
  }
]

function App() {
  const [code, setCode] = useState(dummyCode)
  const [messages, setMessages] = useState(dummyMessages)
  const [cursorPosition, setCursorPosition] = useState(0)
  const [wrongChars, setWrongChars] = useState(0)

  useEffect(() => {
    let socket
    if(process.env.NODE_ENV !== 'production') {
      socket = io('http://localhost:3101')
    } else {
      socket = io()
    }

    socket.on('connect', () => console.log('hello, world'))
    socket.on('disconnect', () => console.log('goodbye, world'))
  },[])

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
