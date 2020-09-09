import React from 'react'
import styled from 'styled-components'
import ChatAndJoinButton from './components/ChatAndJoinButton'
import CodeField from './components/CodeField'
import colors from './colors'
import isMobile from 'ismobilejs'
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
  return isMobile(window.navigator).any ? (
    <div>
      <h3>This game is not playable on mobile devices</h3>
      <p>To play liracer, open it on a laptop or desktop computer.</p>
    </div>
  ) : (
    <Grid>
      <ChatAndJoinButton />
      <CodeField />
    </Grid>
  ) 
}

export default App
