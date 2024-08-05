<script setup lang="ts">
import { type ChatHistory } from '@/types/ChatHistory'

const props = defineProps<{
  history: ChatHistory;
}>();

function isBotResponse() : boolean {
  return typeof props.history.botResponse !== "undefined" && props.history.botResponse !== "";
}

function chatMessage() : string | undefined {
  return isBotResponse() ? props.history.botResponse : props.history.userMessage;
}
</script>

<template>
  <div :class="['chat', { 'chat-start': isBotResponse(), 'chat-end': !isBotResponse() }]">
    <div class="chat-image avatar">
      <div class="w-10 rounded-full">
        <img
          alt="Tailwind CSS chat bubble component"
          src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp" />
      </div>
    </div>
    <div
      :class="['chat-bubble', { 'chat-bubble-primary': isBotResponse(), 'chat-bubble-secondary': !isBotResponse() }]"
      v-html="chatMessage()"
    >
    </div>
  </div>
</template>

<style scoped>

</style>