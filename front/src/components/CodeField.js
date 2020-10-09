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

const mapKeyToChar = (key) => {
  if(['Shift', 'Meta', 'Alt', 'Control', 'Backspace'].includes(key)){
    return null
  } else if (key === "Enter"){
    return "\n"
  } else if (key === 'Tab') {
    return "\t"
  } else {
    return key
  }
}

const CodeField = (props) => {
  const handleKeyDown = (event) => {
    // Include description of why the key needs to be preventDefaulted.
    const preventDefaultKeys = [
      'Tab' // Iterates through ui elements on Chrome, Firefox
    ]
    preventDefaultKeys.includes(event.key) && event.preventDefault()

    const char = mapKeyToChar(event.key)
    if(char) {
      if(props.code[props.cursorPosition] === char) {
        props.setCursorPosition(props.cursorPosition + 1)
      }
    } else {
      if(event.key === 'Backspace') {
        if(props.wrongChars > 0) {
          props.setWrongChars(props.wrongChars - 1)
        } else if(props.cursorPosition > 0) {
          props.setCursorPosition(props.cursorPosition - 1)
        }
      }
    }
  }

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