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
  box-sizing: border-box;
  height: 100%;
`

const mapKeyToChar = (key) => {
  if (key === "Enter"){
    return "\n"
  } else if (key === 'Tab') {
    return "\t"
  } else if (key.length === 1) {
    return key
  } else {
    return null
  }
}

const CodeField = (props) => {
  const handleKeyDown = (event) => {
    event.preventDefault()

    const char = mapKeyToChar(event.key)
    if(char) {
      if (props.wrongChars === 0 && props.snippet.code[props.cursorPosition] === char){
        props.setCursorPosition(props.cursorPosition + 1)
      } else {
        props.setWrongChars(props.wrongChars + 1)
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
      {
        props.snippet && (
          <Pre onKeyDown={handleKeyDown} tabIndex='0'>
            {
              props.snippet.code.split('').map((char, index) => {
                const isOnPlayerCursor = index === props.cursorPosition
                const isOnOpponentCursor = Object.values(props.opponentCursorPositions).some(position => position === index)
                const isOnLastWrongChar = props.wrongChars > 0 && index === props.cursorPosition + props.wrongChars - 1
                const isOnWrongChar = index >= props.cursorPosition && index < props.cursorPosition + props.wrongChars

                let style = {}

                isOnOpponentCursor && (style.background = colors.opponentCursorColor)
                isOnPlayerCursor && (style.background = colors.playerCursorColor)

                if(props.wrongChars > 0) {
                  if(isOnWrongChar)
                  style.background = colors.wrongCharColor
                }

                // Visualize newlines, by using the ↵ character
                // Only show ↵ when the cursor, or wrongChars markings is on newline
                if(char === "\n" && (isOnLastWrongChar || isOnOpponentCursor || (isOnPlayerCursor && !isOnWrongChar))) {
                  char = "↵\n"
                }

                return (
                  <span key={index} style={style}>
                    {
                      char
                    }
                  </span>
                )
              })
            }
          </Pre>
        )
      }
    </Window>
  )
}

export default CodeField