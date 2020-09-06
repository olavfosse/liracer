import React from 'react'
import styled from 'styled-components'
import ChatAndButtons from './components/ChatAndButtons'
import CodeField from './components/CodeField'

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
  grid-template-columns: minmax(300px, 500px) min(750px);
`

function App() {
  return (
    <Grid>
      <ChatAndButtons />
      <CodeField />
    </Grid>
  )
}

export default App
