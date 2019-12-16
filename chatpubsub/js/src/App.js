import React from 'react'
import styled from 'styled-components'

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
  const [value, setValue] = React.useState('')
  const handleChange = event => setValue(event.target.value)

  const handleSubmit = event => {
    event.preventDefault()
    event.stopPropagation()
    alert('A name was submitted: ' + value)
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
      </SubContainer>
    </Container>
  )
}

export default App
