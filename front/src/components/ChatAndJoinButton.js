import React from 'react'
import Window from './Window'
import Messages from './Messages'
import styled from 'styled-components'
import colors from '../colors'

const FlexBox = styled.div`
  display: flex;
  flex-direction: column;
  height: 100%;
`

const ChatForm = styled.form`
  background: ${colors.layer2Background};
  height: 50px;
  display: flex;
  padding: calc(1rem / 2);
  box-sizing: border-box;
`

const ChatInput = styled.input`
  background: ${colors.layer1Background};
  color: ${colors.layer1Foreground};
  border: 0;
  width: 100%;
  padding: 0 3px 3px 0;
  box-sizing: border-box;
  outline: none;
  border-right: 1em solid ${colors.layer2Background};
`

const buttonCSS = `
  :hover {
    background: ${colors.layer2Background};
    color: ${colors.layer2Foreground};
  }
  background: ${colors.layer1Background};
  color: ${colors.layer1Foreground};
  transition-duration: 0.12s;
  font-weight: bold;
  font-size: 14px;
  padding: 0;
  cursor: pointer;
  border: 0;
  border-radius: 5px;
  outline: none;
`

const ChatSubmit = styled.button`
  ${buttonCSS}
  width: 70px;
  border: solid ${colors.layer2Background} 0;
`

const JoinButtonContainer = styled.div`
  border: calc(1em / 2) solid ${colors.layer2Background};
  border-top: 0;
  height: 40px;
  box-sizing: border-box;
  background: ${colors.layer2Background};
`

const JoinButton = styled.button`
  ${buttonCSS}
  height: 100%;
  width: 100%;
`

const ChatAndJoinButton = (props) => {
  return (
    <Window>
      <FlexBox>
        <Messages messages={props.messages}/>
        <ChatForm onSubmit={props.handleSendMessage}>
          <ChatInput name='input' />
          <ChatSubmit>SEND</ChatSubmit>
        </ChatForm>
        <JoinButtonContainer>
          <JoinButton onClick={props.handleClickJoinGame}>
            JOIN GAME
          </JoinButton>
        </JoinButtonContainer>
      </FlexBox>
    </Window>
  )
}

export default ChatAndJoinButton