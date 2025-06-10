type ChatMessage = {
  id?: string
  username: string
  message: string
  timestamp: string
  reactions?: number
  hasReacted?: boolean
  type?: string
  messageId?: string
}

type MessageHandler = (msg: ChatMessage) => void

export class ChatWebSocket {
  private ws: WebSocket | null = null
  private url: string
  private onMessageHandler: MessageHandler | null = null

  constructor(url: string) {
    this.url = url
  }

  connect(onMessage: MessageHandler) {
    this.ws = new WebSocket(this.url)
    this.onMessageHandler = onMessage

    this.ws.onopen = () => {}

    this.ws.onmessage = (event: MessageEvent) => {
      try {
        const data = JSON.parse(event.data)
        this.onMessageHandler && this.onMessageHandler(data)
      } catch (e) {
        // Ignore malformed messages
      }
    }

    this.ws.onclose = () => {}
    this.ws.onerror = (err) => {}
  }

  sendMessage(message: string) {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify({ message }))
    }
  }

  sendReaction(messageId: string) {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify({ type: "reaction", messageId }))
    }
  }

  close() {
    if (this.ws) {
      this.ws.close()
      this.ws = null
    }
  }
}

export function getChatWebSocketURL(): string {
  const protocol = window.location.protocol === "https:" ? "wss" : "ws"
  return `${protocol}://${window.location.host}/api/v1/protected/ws/chat`
}
