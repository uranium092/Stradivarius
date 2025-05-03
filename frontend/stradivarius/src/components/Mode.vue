<script lang="ts" setup>
import { useFetchingStore } from '../stores/fetching';

const fetchingStore = useFetchingStore();
</script>

<template>
  <div class="overflow-hidden rounded-xl border border-gray-100 p-1 w-[80%] md:w-[35%] m-auto">
    <ul class="flex items-center gap-2 text-sm font-medium">
      <li class="flex-1">
        <span
          :class="[
            'flex items-center justify-center gap-2 rounded-lg hover:bg-white px-3 py-2 cursor-pointer',
            fetchingStore.mode === 'all' ? 'relative shadow-lg' : 'text-gray-500 hover:shadow',
          ]"
          @click="fetchingStore.changeMode('all')"
        >
          All Stock
        </span>
      </li>
      <li class="flex-1">
        <span
          :class="[
            'flex items-center justify-center gap-2 rounded-lg hover:bg-white px-3 py-2 cursor-pointer',
            fetchingStore.mode === 'recomm'
              ? 'relative bg-white shadow-lg'
              : 'text-gray-500 hover:shadow',
          ]"
          @click="fetchingStore.changeMode('recomm')"
        >
          Recommendations</span
        >
      </li>
    </ul>
  </div>
</template>
