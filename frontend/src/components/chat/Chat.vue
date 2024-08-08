<script setup lang="ts">
import ChatMessages from '@/components/chat/ChatMessages.vue'
import ChatInput from '@/components/chat/ChatInput.vue'

import { useChatMessages } from '@/composables/useChatMessages'
import { type ChatHistory } from '@/types/ChatHistory'
import { historyService } from '@/services/historyService'

const { histories, loadingMessage, handleNewHistory } =  useChatMessages()

const addNewHistory = (newHistory: ChatHistory) => {
  if (newHistory !== null && typeof newHistory.userMessage !== 'undefined' && newHistory.userMessage.trim() !== '') {
    handleNewHistory(Object.create(newHistory))

    historyService.send(newHistory)

    loadingMessage.value = true
  }
}
</script>

<template>
  <div class="flex flex-col space-y-4 h-screen">
    <ChatMessages :histories="histories" :loading-message="loadingMessage" />
    <ChatInput @add-new-history="addNewHistory" :loading-message="loadingMessage" />
  </div>
</template>

<style scoped>

</style>