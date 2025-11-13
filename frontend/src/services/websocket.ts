export class WebSocketService {
  private ws: WebSocket | null = null;
  private reconnectTimeout: number = 1000;

  connect(studentId: string) {
    const wsUrl = `ws://localhost:8080/ws/${studentId}`;
    this.ws = new WebSocket(wsUrl);

    this.ws.onopen = () => {
      console.log('WebSocket connected');
      this.reconnectTimeout = 1000;
    };

    this.ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      // Handle real-time updates
        console.log('WebSocket message received:', data);
        
    };

    this.ws.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    this.ws.onclose = () => {
      // Auto-reconnect with exponential backoff
      setTimeout(() => {
        this.reconnectTimeout = Math.min(this.reconnectTimeout * 2, 30000);
        this.connect(studentId);
      }, this.reconnectTimeout);
    };
  }

  disconnect() {
    this.ws?.close();
  }

  send(message: any) {
    if (this.ws?.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify(message));
    }
  }
}
