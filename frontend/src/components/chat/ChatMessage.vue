<script setup lang="ts">
import { ref, onMounted } from "vue";
import { type ChatHistory } from '@/types/ChatHistory'
import { extractLinks } from '@/utils/linkMatches'

import botAvatar from "@/assets/bot.png"
import userAvatar from "@/assets/girl.png"

import Carousel from '@/components/ui/Carousel.vue'

const props = defineProps<{
  history: ChatHistory;
}>();

const links = ref<string[]>([])

onMounted(() => {
  if (isBotResponse()) {
    const botResponse = props.history.botResponse as string

    links.value = extractLinks(botResponse)
  }
})

function isBotResponse() : boolean {
  return typeof props.history.botResponse !== "undefined" && props.history.botResponse !== ""
}

function chatMessage() : string | undefined {
  return isBotResponse() ? props.history.botResponse : props.history.userMessage
}

function avatarImage(): string {
  return isBotResponse() ? botAvatar : userAvatar
}
</script>

<template>
  <div>
    <div :class="['chat', { 'chat-start': isBotResponse(), 'chat-end': !isBotResponse() }]">
      <div class="chat-image avatar">
        <div class="w-10 rounded-full">
          <img
            alt="Tailwind CSS chat bubble component"
            :src="avatarImage()" />
<!--          <a href="https://www.flaticon.com/free-icons/bot" title="bot icons">Bot icons created by Freepik - Flaticon</a>-->
<!--          <a href="https://www.flaticon.com/free-icons/girl" title="girl icons">Girl icons created by Freepik - Flaticon</a>-->
        </div>
      </div>
      <div
        :class="['chat-bubble', { 'chat-bubble-primary': isBotResponse(), 'chat-bubble-secondary': !isBotResponse() }]"
        v-html="chatMessage()"
      >
      </div>
    </div>
    <Carousel
      class="my-4"
      v-if="links.length > 0"
      :links="links" />
  </div>
</template>

<style scoped>

</style>