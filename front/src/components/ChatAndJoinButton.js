import React from 'react'
import Window from './Window'
import styled from 'styled-components'
import colors from '../colors'

const FlexBox = styled.div`
  display: flex;
  flex-direction: column;
  height: 100%;
`

const Messages = styled.div`
  flex-basis: 100%;
  padding: 1rem;
`

const ChatInputFormAndJoinButtonContainerSharedCSS = `
  flex-basis: auto;
  height: 50px;
  background: ${colors.layer2Background};
  padding: calc(1rem / 2);
  box-sizing: border-box;
`

const ChatInputForm = styled.form`${ChatInputFormAndJoinButtonContainerSharedCSS}`

const JoinButtonContainer = styled.div`${ChatInputFormAndJoinButtonContainerSharedCSS}`

const JoinButtonAndChatInputSharedCSS = `
  border: 0;
  height: 100%;
  width: 100%;
  padding: 0;
  background: ${colors.layer1Background};
  outline: none; // CONSIDER removing this line. The outline makes it clear what box is selected although it does look ugly. If it is enabled again it should also be enabled on CodeField.
`

const ChatInput = styled.input`${JoinButtonAndChatInputSharedCSS}`

const JoinButton = styled.button`
  ${JoinButtonAndChatInputSharedCSS}
  transition-duration: 0.1s;
  font-size: 1em;
  :hover {
    background: ${colors.layer2Background};;
  }
`

const ChatAndJoinButton = () => {
  return (
    <Window>
      <FlexBox>
        <Messages>
          <b>liracer</b> dit
          <br/>
          <b>liracer</b> dat
        </Messages>
        <ChatInputForm>
          <ChatInput/>
        </ChatInputForm>
        <JoinButtonContainer>
          <JoinButton onClick={() => alert('This will prompt a game id and join it. Alternatively it could turnn into a text field when clicked where the game id can be entered.')}>JOIN GAME</JoinButton>
        </JoinButtonContainer>
      </FlexBox>
    </Window>
  )
}

export default ChatAndJoinButton