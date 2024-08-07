<script setup lang="ts">
import { ref } from 'vue'
import { type ChatHistory } from '@/types/ChatHistory'

const emit = defineEmits(['addNewHistory'])

const fileInput = ref<HTMLInputElement>()

const props = defineProps<{
  loadingMessage: boolean;
}>()

const history = defineModel<ChatHistory>({
  default: {} as ChatHistory,
})

function chooseFile() {
  fileInput.value?.click()
}

function getFile(event: Event) {
  const file = event.target.files[0]
  let reader = new FileReader();
  reader.readAsDataURL(file)
  reader.onload = function () {
    const historyWithImage: ChatHistory = { userMessage: reader.result as string }
    emit("addNewHistory", historyWithImage)
  }
  reader.onerror = function (error) {
    console.log('Error: ', error)
  }
}
</script>

<template>
<form class="flex flex-row gap-x-2" @submit.prevent>
  <div class="flex-none">
    <label>
      <button
        @click="chooseFile()"
        :class="['btn', 'btn-primary', { 'btn-disabled': props.loadingMessage }]">
        <svg class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
          <path fill-rule="evenodd" d="M13 10a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H14a1 1 0 0 1-1-1Z" clip-rule="evenodd"/>
          <path fill-rule="evenodd" d="M2 6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12c0 .556-.227 1.06-.593 1.422A.999.999 0 0 1 20.5 20H4a2.002 2.002 0 0 1-2-2V6Zm6.892 12 3.833-5.356-3.99-4.322a1 1 0 0 0-1.549.097L4 12.879V6h16v9.95l-3.257-3.619a1 1 0 0 0-1.557.088L11.2 18H8.892Z" clip-rule="evenodd"/>
        </svg>
      </button>
      <input @change="getFile" type="file" ref="fileInput" accept="image/gif, image/jpeg, image/png" hidden/>
    </label>
  </div>
  <div class="flex-none">
    <button
      :class="['btn', 'btn-primary', { 'btn-disabled': props.loadingMessage }]">
      <svg class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
        <path fill-rule="evenodd" d="M5 8a1 1 0 0 1 1 1v3a4.006 4.006 0 0 0 4 4h4a4.006 4.006 0 0 0 4-4V9a1 1 0 1 1 2 0v3.001A6.006 6.006 0 0 1 14.001 18H13v2h2a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h2v-2H9.999A6.006 6.006 0 0 1 4 12.001V9a1 1 0 0 1 1-1Z" clip-rule="evenodd"/>
        <path d="M7 6a4 4 0 0 1 4-4h2a4 4 0 0 1 4 4v5a4 4 0 0 1-4 4h-2a4 4 0 0 1-4-4V6Z"/>
      </svg>
    </button>
  </div>
  <div class="grow">
    <input
      type="text"
      v-model="history.userMessage"
      @keyup.enter="$emit('addNewHistory', history); history = {} as ChatHistory;"
      :disabled="props.loadingMessage"
      placeholder="Envia un mensaje a Deal Genius"
      class="input input-bordered w-full" />
  </div>
  <div class="flex-none">
    <button
      :class="['btn', 'btn-primary', { 'btn-disabled': props.loadingMessage }]"
      :disabled="props.loadingMessage"
      @click="$emit('addNewHistory' ,history); history = {} as ChatHistory;">
      <svg class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
        <path d="M5.027 10.9a8.729 8.729 0 0 1 6.422-3.62v-1.2A2.061 2.061 0 0 1 12.61 4.2a1.986 1.986 0 0 1 2.104.23l5.491 4.308a2.11 2.11 0 0 1 .588 2.566 2.109 2.109 0 0 1-.588.734l-5.489 4.308a1.983 1.983 0 0 1-2.104.228 2.065 2.065 0 0 1-1.16-1.876v-.942c-5.33 1.284-6.212 5.251-6.25 5.441a1 1 0 0 1-.923.806h-.06a1.003 1.003 0 0 1-.955-.7A10.221 10.221 0 0 1 5.027 10.9Z"/>
      </svg>
    </button>
  </div>
</form>
</template>

<style scoped>

</style>