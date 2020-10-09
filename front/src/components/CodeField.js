import React from 'react'
import Window from './Window'
import styled from 'styled-components'
import colors from '../colors'

const contentPadding = '1rem'

const Wrapper = styled.pre`
  padding: ${contentPadding};
  margin: 0;
  color: ${colors.layer1Foreground}
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