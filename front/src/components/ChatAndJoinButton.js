import React from 'react'
import Window from './Window'
import styled from 'styled-components'
import colors from '../colors'

const contentPadding = '1rem'
const inputHeight = '50px'

const Messages = styled.div`
  height: calc(100% - ${inputHeight} - ${contentPadding} - ${contentPadding});
  padding: ${contentPadding};
`

const ChatForm = styled.form`
  box-sizing: border-box;
  width: auto;
  height: 50px;
  background: ${colors.layer2Background};
  padding: calc(1em / 2);
`

const ChatInput = styled.input`
  width: calc(100% - 1em);
  height: 100%;
  width: 100%;
  border: 0;
  margin: 0;
  padding: 0;
  background: ${colors.layer1Background}
`

const ChatAndJoinButton = () => {
  return (
    <Window>
      <Messages>
        <b>liracer</b> dit
        <br/>
        <b>liracer</b> dat
      </Messages>
      <ChatForm>
        <ChatInput/>
      </ChatForm>
    {/* The chat button is going here, hence the name */}
    </Window>
  )
}

export default ChatAndJoinButton