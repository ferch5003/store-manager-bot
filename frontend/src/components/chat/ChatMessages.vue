<script setup lang="ts">
import { ref, onUpdated } from "vue";
import ChatMessage from '@/components/chat/ChatMessage.vue'
import type { ChatHistory } from '@/types/ChatHistory'

const chatMessagesContainer = ref<HTMLDivElement>()

const props = defineProps<{
  histories: Array<ChatHistory>;
  loadingMessage: boolean;
}>();

onUpdated(() => {
  scrollToLastMessage()
});

const scrollToLastMessage = () => {
  console.log(chatMessagesContainer)

  if (!chatMessagesContainer.value) {
    return
  }

  const lastChildElement = chatMessagesContainer.value.lastElementChild;

  lastChildElement?.scrollIntoView({
        behavior: 'smooth',
  })
}
</script>

<template>
  <div
    class="grow h-0 p-8 bg-info-content rounded-box relative overflow-y-auto"
    ref="chatMessagesContainer">
    <ChatMessage v-for="history in props.histories" :key="history.id" :history="history"/>
    <div
      class="flex flex-col w-9/12 mt-6"
      v-if="props.loadingMessage">
      <div class="skeleton bg-neutral h-32 w-full"></div>
    </div>
  </div>
</template>

<style scoped>

</style>