<script setup lang="ts">
import ChatMessage from '@/components/chat/ChatMessage.vue'

import { ref } from "vue";
import { type ChatHistory } from '@/types/ChatHistory'

const histories = ref<ChatHistory[]>([{
  botResponse: "<p>Hola<br>¿Cómo puedo ayudarte?</p>",
}]);
</script>

<template>
  <div class="grow p-8 bg-info-content rounded-box relative overflow-y-auto">
    <ChatMessage v-for="history in histories" :key="history.id" :history="history"/>
  </div>
</template>

<style scoped>

</style>