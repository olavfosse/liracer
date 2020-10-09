import React from 'react'
import Window from './Window'
import styled from 'styled-components'
import colors from '../colors'

const contentPadding = '1rem'

const Pre = styled.pre`
  padding: ${contentPadding};
  margin: 0;
  color: ${colors.layer1Foreground}
`

const CodeField = (props) => {
  return (
    <Window>
      <Pre>
        { props.code }
      </Pre>
    </Window>
  )
}

export default CodeField