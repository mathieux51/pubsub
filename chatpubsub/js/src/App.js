import React from 'react'
import styled from 'styled-components'
import { WSContext } from './context/WSContext'
// import { WSProvider } from 'src/context/WSContext'

const Container = styled.div`
  display: flex;
  margin: 1rem;
`
const SubContainer = styled.div`
  flex: 1;
  margin: 1rem;
  padding: 1rem;
  border: 1px solid lightgray;
  border-radius: 8px;
  box-shadow: 2px 2px 3px #0000001f;
`

const Title = styled.h1`
  font-weight: normal;
  font-size: 16px;
`

const Input = styled.input`
  width: 7rem;
  height: 2rem;
`

function App() {
  const { messages, send } = React.useContext(WSContext)
  const [value, setValue] = React.useState('')
  const handleChange = event => setValue(event.target.value)

  const handleSubmit = event => {
    event.preventDefault()
    event.stopPropagation()
    send(value)
  }
  return (
    <Container>
      <SubContainer>
        <Title>Send</Title>
        <form onSubmit={handleSubmit}>
          <label>
            Message:
            <Input type='text' value={value} onChange={handleChange} />
          </label>
          <input type='submit' value='Submit' />
        </form>
      </SubContainer>
      <SubContainer>
        <Title>Receive</Title>
        <ul>
          {messages.map((message, i) => (
            <li key={message + i}>{message}</li>
          ))}
        </ul>
      </SubContainer>
    </Container>
  )
}

export default App
