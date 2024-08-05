<script setup lang="ts">
import ChatMessage from '@/components/chat/ChatMessage.vue'
import type { ChatHistory } from '@/types/ChatHistory'

const props = defineProps<{
  histories: Array<ChatHistory>;
  loadingMessage: boolean;
}>();
</script>

<template>
  <div class="grow h-0 p-8 bg-info-content rounded-box relative overflow-y-auto">
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