import React from 'react'
import styled from 'styled-components'
import colors from '../colors'

const MessageDiv = styled.div`
  // Make it easy to visually differentiate messages
  padding-bottom: 7px;
`

const Message = ({ sender, content }) => {
  return (
    <MessageDiv>
      {/* IMPORTANT SPACE --> */}
      <b>&lt;{sender}&gt;</b> <span>{content}</span>
    </MessageDiv>
  )
}

const MessagesDiv = styled.div`
  flex-basis: 100%;
  padding: 1rem;
  color: ${colors.layer1Foreground};
  // Copied from default React index.css
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
`

const Messages = () => {
  return (
    <MessagesDiv>
      <Message sender="liracer" content='Click the JOIN button or type "/join GameID" to join a game.'/>
      <Message sender="liracer" content='Click the JOIN button or type "/join GameID" to join a game.'/>
      <Message sender="fossegrim" content='Another sample message'/>
    </MessagesDiv>
  )
}

export default Messages