<script setup lang="ts">
import Previewer from '@/components/ui/Previewer.vue'
import { register } from 'swiper/element/bundle';

register();

const props = defineProps<{
  links: string[];
}>();

const spaceBetween = 10
const breakPoints = {
  460: {
    slidesPerView: 1,
    spaceBetween: 10
  },

  920: {
    slidesPerView: 2,
    spaceBetween: 20
  },

  1275: {
    slidesPerView: 3,
    spaceBetween: 40
  },

  1640: {
    slidesPerView: 4,
    spaceBetween: 40
  }
}
</script>

<template>
  <swiper-container
    :slides-per-view="3"
    :space-between="spaceBetween"
    :pagination="{
      hideOnClick: true
    }"
    :breakpoints="breakPoints"
  >
    <swiper-slide
      v-for="(link, index) in props.links"
      :key="index">
      <Previewer :link="link" />
    </swiper-slide>
  </swiper-container>
</template>

<style scoped>

</style>