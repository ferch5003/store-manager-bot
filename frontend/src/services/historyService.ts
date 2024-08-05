import { type ChatHistory } from '@/types/ChatHistory'
import { toSnakeCase, toCamelCase } from '@/utils/caseConverter';

export class HistoryService  {
  private webSocket: WebSocket | null = null;
  private readonly url: string;
  private readonly reconnectInterval: number;
  private reconnectAttempts: number = 0;
  private maxReconnectAttempts: number = 10;
  private messageCallback: ((history: ChatHistory) => void) | null = null;


  constructor(url: string, reconnectInterval: number = 5000) {
    this.url = url;
    this.reconnectInterval = reconnectInterval;

    this.connect()
  }

  private connect(): void {
    const endpoint = this.url + "/histories";
    this.webSocket = new WebSocket(endpoint);

    this.webSocket.onopen = this.onOpen.bind(this);
    this.webSocket.onmessage = this.onMessage.bind(this);
    this.webSocket.onerror = this.onError.bind(this);
    this.webSocket.onclose = this.onClose.bind(this);
  }

  private onOpen(event: Event): void {
    console.log("WebSocket connection opened");
    this.reconnectAttempts = 0;
  }

  private onMessage(event: MessageEvent): void {
    try {
      const historyResponse: ChatHistory = JSON.parse(event.data);

      if (this.messageCallback) {
        this.messageCallback(toCamelCase(historyResponse));
      }
    } catch (error) {
      console.error('Error parsing message data:', error);
    }
  }

  private onError(event: Event):void {
    console.log("WebSocket error: ", event)
  }

  private onClose(event: CloseEvent): void {
    console.log("WebSocket connection closed:", event)
    if (this.reconnectAttempts < this.maxReconnectAttempts) {
      setTimeout(() => {
        this.reconnectAttempts++;
        this.connect();
      }, this.reconnectInterval);
    }
  }

  public send(history: ChatHistory):void {
    if (this.webSocket && this.webSocket.readyState === WebSocket.OPEN) {
      this.webSocket.send(JSON.stringify(toSnakeCase(history)))
    } else {
      console.error('WebSocket connection is not open');
    }
  }

  public close(): void {
    if (this.webSocket) {
      this.webSocket.close()
    }
  }

  public setMessageCallback(callback: (history: ChatHistory) => void): void {
    this.messageCallback = callback;
  }
}

const apiWSURL = import.meta.env.VITE_BACKEND_WS_URL as string;

export const historyService: HistoryService = new HistoryService(apiWSURL);