<script lang="ts" setup>
import { useFetchingStore } from '../stores/fetching';

const fetchingStore = useFetchingStore();
</script>

<template>
  <div
    v-if="fetchingStore.hasError"
    class="font-regular w-auto absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 rounded-lg bg-red-500 p-4 text-base leading-5 text-white opacity-100"
  >
    Your request could not be processed
  </div>
</template>
