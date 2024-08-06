import showdown from 'showdown';
import { ref, onMounted, onBeforeUnmount } from "vue";
import type { ChatHistory } from '@/types/ChatHistory'
import { historyService } from '@/services/historyService';

export function useChatMessages() {
  const histories = ref<ChatHistory[]>([{
    botResponse: "<p>Hola<br>¿Cómo puedo ayudarte?</p>",
  }]);
  const loadingMessage = ref<boolean>(false)

  const converter = new showdown.Converter();

  const handleNewHistory = (history: ChatHistory) => {
    if (typeof history.userMessage !== "undefined" && history.userMessage !== "") {
      // If user message is in, but bot response is empty then add default response.
      if (typeof history.botResponse !== "undefined" && history.botResponse === "") {
        history.botResponse = "Disculpe, no pude encontrar una respuesta a esa petición."
      }

      history.userMessage = converter.makeHtml(history.userMessage)

      if (typeof history.botResponse !== "undefined" && history.botResponse !== "") {
        history.botResponse = converter.makeHtml(history.botResponse)
      }

      histories.value.push(history);

      loadingMessage.value = false
    }
  };

  onMounted(() => {
    historyService.setMessageCallback(handleNewHistory);
  });

  onBeforeUnmount(() => {
    historyService.close();
  });

  return {
    histories,
    loadingMessage,
    handleNewHistory,
  };
}