import React from 'react'
import Window from './Window'
import styled from 'styled-components'

const contentPadding = '1rem'

const Wrapper = styled.div`
  padding: ${contentPadding}
`

const ChatAndButtons = () => {
  return (
    <Window>
      <Wrapper>
        Chat and buttons
      </Wrapper>
    </Window>
  )
}

export default ChatAndButtons