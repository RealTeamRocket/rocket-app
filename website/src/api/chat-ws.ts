type ChatMessage = {
  username: string
  message: string
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
    // Use ws:// for local, wss:// for production
    // Cookies (jwt_token) will be sent automatically
    this.ws = new WebSocket(this.url)
    this.onMessageHandler = onMessage

    this.ws.onopen = () => {
      // Connection established
    }

    this.ws.onmessage = (event: MessageEvent) => {
      try {
        const data = JSON.parse(event.data)
        if (data.username && data.message) {
          this.onMessageHandler && this.onMessageHandler(data)
        }
      } catch (e) {
        // Ignore malformed messages
      }
    }

    this.ws.onclose = () => {
      // Handle close if needed
    }

    this.ws.onerror = (err) => {
      // Handle error if needed
    }
  }

  sendMessage(message: string) {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify({ message }))
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
