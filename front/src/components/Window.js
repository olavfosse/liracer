import React from 'react'
import styled from 'styled-components'
import colors from '../colors'

const titleBarHeight = '30px';

const Titlebar = styled.div`
  height: ${titleBarHeight};
  background: ${colors.layer2Background};
`
const Content = styled.div`
  height: calc(100% - ${titleBarHeight});
  padding: 1rem;
  background: ${colors.layer1Background};
`

const Window = ({ children }) => {
  return (
    <div>
      <Titlebar/>
      <Content>{children}</Content>
    </div>
  )
}

export default Window