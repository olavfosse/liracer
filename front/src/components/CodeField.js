import React from 'react'
import { useState } from 'react'
import Window from './Window'
import styled from 'styled-components'
import colors from '../colors'
import constants from '../utils/constants';

const contentPadding = '1rem'

const Pre = styled.pre`
  padding: ${contentPadding};
  margin: 0;
  color: ${colors.layer1Foreground};
  outline: none;
  box-sizing: border-box;
  height: 100%;
  line-height: calc(100% + 3px); // Make space for the "underline"
`

const mapKeyToChar = (key) => {
  if (key === "Enter") {
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
  const [isFocused, setIsFocused] = useState(false)

  const handleKeyDown = (event) => {
    // For each key, specify a problem that happens if it is not eventPreventDefaulted and one or more environments (browser and os) it occurs in
    const keysToEventPreventDefault = [
      /*
       * problem: Switches focus
       * environment: Chrome/MacOS
       * environment: Safari/MacOS
       */
      'Tab',
      /*
       * problem: Navigates history
       * environment: Firefox/MacOS
       */
      'Backspace'
    ]
    if(keysToEventPreventDefault.includes(event.key)) {
      event.preventDefault()
    }

    const char = mapKeyToChar(event.key)
    if (props.countdownUntilStart !== constants.COUNTDOWN_FINAL_NUMBER) {
      props.setWrongChars(wrongChars => wrongChars + 1)

      return;
    }

    if (char) {
      if (props.wrongChars === 0 && props.snippet.code[props.cursorPosition] === char) {
        props.setCursorPosition(cursorPosition => cursorPosition + 1)
      } else {
        props.setWrongChars(wrongChars => wrongChars + 1)
      }
    } else if(event.key === 'Backspace') {
      if(props.wrongChars > 0) {
        props.setWrongChars(props.wrongChars - 1)
      } else if(props.cursorPosition > 0) {
        props.setCursorPosition(props.cursorPosition - 1)
      }
    }
  }

  if (props.countdownUntilStart !== constants.COUNTDOWN_FINAL_NUMBER)
    return (
      <Window>
        <Countdown onKeyDown={handleKeyDown} tabIndex='0'>
          {props.countdownUntilStart}
        </Countdown>
      </Window>
    )

  return (
    <Window>
      {
        props.snippet && (
          <Pre onKeyDown={handleKeyDown}
               tabIndex='0'
               onBlur={() => setIsFocused(false)}
               onFocus={() => setIsFocused(true)}>
            {
              props.snippet.code.split('').map((char, index) => {
                const isOnPlayerCursor = index === props.cursorPosition
                const isOnOpponentCursor = Object.values(props.opponentCursorPositions).some(position => position === index)
                const isOnLastWrongChar = props.wrongChars > 0 && index === props.cursorPosition + props.wrongChars - 1
                const isOnWrongChar = index >= props.cursorPosition && index < props.cursorPosition + props.wrongChars
                const isOnLastChar = isOnLastWrongChar || (!isOnWrongChar && isOnPlayerCursor)
                let style = {}

                isOnOpponentCursor && (style.background = colors.opponentCursorColor)
                if(props.wrongChars > 0) {
                  if(isOnWrongChar) {
                    if(isFocused) {
                      style.background = colors.wrongCharColor
                    } else {
                      style.borderBottomStyle = 'solid'
                      style.borderColor = colors.wrongCharColor
                    }
                  }
                } else {
                  if(isOnPlayerCursor) {
                    if(isFocused) {
                      style.background = colors.playerCursorColor
                    } else {
                      style.borderBottomStyle = 'solid'
                      style.borderColor = colors.playerCursorColor
                    }
                  }
                }

                // Visualize newlines, by using the ↵ character
                // Only show ↵ when the cursor, or wrongChars markings is on newline
                if(char === "\n" && (isOnLastChar || isOnOpponentCursor)) {
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
