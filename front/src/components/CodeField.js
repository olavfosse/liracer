import React from 'react'
import Window from './Window'
import styled from 'styled-components'

const contentPadding = '1rem'

const Wrapper = styled.div`
  padding: ${contentPadding};
  font-family: monospace;
`

const CodeField = () => {
  return (
    <Window>
      <Wrapper>
        Code field
      </Wrapper>
    </Window>
  )
}

export default CodeField