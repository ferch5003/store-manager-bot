<script setup lang="ts">
  import LinkPrevue from 'link-prevue'

  const _defaultImageURL = "https://http2.mlstatic.com/static/org-img/homesnw/mercado-libre.png?v=2"

  const props = defineProps<{
    link: string;
  }>();

  function openPageLink(url: string) {
    window.open(url, '_blank')
  }

  function getImageURL(image: string) : string {
    if (image !== null && typeof image !== "undefined" && image !== '') {
      return image
    }

    return _defaultImageURL
  }
</script>

<template>
  <LinkPrevue
    :url="props.link">
    <template v-slot="props">
      <div
        @click="openPageLink(props.url)"
        class="card bg-primary w-60 h-60 md:w-80 md:h-80 lg:w-96 lg:h-96 hover:cursor-pointer">
        <figure class="h-1/2">
          <img
            class="text-black w-full h-full object-cover"
            :src="getImageURL(props.img)"
            :alt="props.title" />
        </figure>
        <div class="card-body h-1/2">
          <h2 class="card-title text-lg text-black">{{props.title}}</h2>
          <p class="truncate text-sm text-black">{{props.description}}</p>
        </div>
      </div>
    </template>
  </LinkPrevue>
</template>

<style scoped>

</style>