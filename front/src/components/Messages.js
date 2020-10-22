import React, { useEffect, useRef, useState } from 'react'
import styled from 'styled-components'
import colors from '../colors'

const ToolTipText = styled.span`
  visibility: hidden;
  background-color: ${colors.layer1Background};
  color: ${colors.layer1Foreground};
  font-weight: 600;
  text-align: center;
  border-radius: 6px;
  border-color: ${colors.layer2Background};
  border-style: solid;
  padding: 5px 0px;
  width: 80%;
  left:0;
  top:20px;
  opacity: 0.8;
  
  /* Position the tooltip */
  position: absolute;
  z-index: 1;
`

const MessageDiv = styled.div`
  // Make it easy to visually differentiate messages
  padding-bottom: 7px;
  position: relative;
  &:hover ${ToolTipText} {
    visibility: visible;
  }
`

const Message = (props) => {
  return (
    <MessageDiv>
      {/* IMPORTANT SPACE               --> */}
      <b>&lt;{props.message.sender}&gt;</b> <span>{props.message.content}</span>
      {props.message.playerID && <ToolTipText>{props.message.playerID}</ToolTipText>}
    </MessageDiv>
  )
}

const MessagesDiv = styled.div`
  flex-basis: 100%;
  padding: 1rem;
  // Scrollbar for message box
  flex: 1 1 auto; 
  overflow-y: auto;
  height: 0px;
  color: ${colors.layer1Foreground};
  // Copied from default React index.css
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
`


const Messages = (props) => {
  /*
   * By default liracer automatically scrolls down when a new message is sent.
   * By scrolling up, liracer enters navigation mode where you can scroll freely without automatic scrolling interfering.
   * Navigation mode is exited by scrolling all the way down.
   */
  const [isInNavigationMode, setIsInNavigationMode] = useState(false)
  const messagesEndRef = useRef(null)

  useEffect(() => {
    !isInNavigationMode && messagesEndRef.current.scrollIntoView({ behavior: "smooth" })
  }, [props.messages])

  const handleScroll = (event) => {
    const  {scrollHeight, scrollTop, offsetHeight} = event.target
    const distanceToBottom =  Math.abs(scrollHeight - scrollTop - offsetHeight)
    /*
     * NB: Testing if `distanceToBottom === 0` is NOT sufficent to verify if the user scrolled all the way down
     * It does not necesarrily evaluate to true on some systems(ubuntu+chromium for example) even though the user scrolled all the way down, such that it is impossible to scroll any longer.
     * See https://github.com/olav35/liracer/pull/49#issuecomment-714450885
     *
     * Instead we test that it is within the range -3..3 which seems to work.
     */
    setIsInNavigationMode(!(distanceToBottom <= 3))
  }

  return (
    <MessagesDiv onScroll={handleScroll}>
      {
        props.messages.map((message, index) => <Message key={index} message={message} />)
      }
      <div ref={messagesEndRef} />
    </MessagesDiv>
  )
}

export default Messages