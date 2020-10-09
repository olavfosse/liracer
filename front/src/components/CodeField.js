import React from 'react'
import Window from './Window'
import styled from 'styled-components'
import colors from '../colors'

const contentPadding = '1rem'

const Pre = styled.pre`
  padding: ${contentPadding};
  margin: 0;
  color: ${colors.layer1Foreground};
  outline: none;
`

const CodeField = (props) => {
  return (
    <Window>
      <Pre onKeyDown={handleKeyDown} tabIndex='0'>
        {
          props.code.split('').map((char, index) => {
            let style = {}

            if(props.wrongChars > 0) {
              if(index >= props.cursorPosition && index < props.cursorPosition + props.wrongChars)
                style.background = colors.wrongCharColor
            } else if (index === props.cursorPosition) {
              style.background = 'rgb(207, 186, 165)'
            }

            return (
              <span key={index}
                style={style}>{char}
              </span>
            )
          })
        }
      </Pre>
    </Window>
  )
}

export default CodeField