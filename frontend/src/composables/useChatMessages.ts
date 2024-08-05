import { ref, onMounted, onBeforeUnmount } from "vue";
import type { ChatHistory } from '@/types/ChatHistory'
import { historyService } from '@/services/historyService';

export function useChatMessages() {
  const histories = ref<ChatHistory[]>([{
    botResponse: "<p>Hola<br>¿Cómo puedo ayudarte?</p>",
  }]);
  const loadingMessage = ref<boolean>(false)

  const handleNewHistory = (history: ChatHistory) => {
    if (typeof history.userMessage !== "undefined" && history.userMessage !== "") {
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