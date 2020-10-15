import React from 'react'
import styled from 'styled-components'
import colors from '../colors'
import './ToolTip.css'

const MessageDiv = styled.div`
  // Make it easy to visually differentiate messages
  padding-bottom: 7px;
`

const Message = (props) => {
  console.log(props.message.playerID)
  return (
    <MessageDiv className="tooltip">
      {/* IMPORTANT SPACE --> */}
      <b>&lt;{props.message.sender}&gt;</b> {props.message.content}
      <span className={props.message.playerID ? 'tooltiptext' : ''}>{props.message.playerID}</span>
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
  return (
    <MessagesDiv>
      {
        props.messages.map((message, index) => <Message key={index} message={ message }/>)
      }
    </MessagesDiv>
  )
}

export default Messages