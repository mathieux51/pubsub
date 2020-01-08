import React from 'react'
import { socketAddress } from '../constants'

export const WSContext = React.createContext()
WSContext.displayName = 'WSContext'
const { Provider, Consumer } = WSContext

export { Consumer as WSConsumer }

// Create WebSocket connection.
const socket = new WebSocket(socketAddress)

export function WSProvider(props) {
  const [messages, setMessages] = React.useState([])

  const send = msg => socket.send(msg)

  // Connection opened
  socket.addEventListener('open', function(event) {
    console.log(`WebSocket ${socketAddress} connected`)
  })

  // Listen for messages
  socket.addEventListener('message', function(event) {
    setMessages([...messages, event.data])
  })

  const state = {
    messages,
    send
  }

  return <Provider value={state}>{props.children}</Provider>
}
