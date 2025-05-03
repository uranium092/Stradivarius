<script lang="ts" setup></script>

<template>
  <div
    class="flex items-center justify-center absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2"
  >
    <button
      type="button"
      class="bg-indigo-400 h-max w-max rounded-lg text-white font-bold hover:bg-indigo-300 hover:cursor-not-allowed duration-[500ms,800ms]"
      disabled
    >
      <div class="flex items-center justify-center m-[10px]">
        <div
          class="h-5 w-5 border-t-transparent border-solid animate-spin rounded-full border-white border-4"
        ></div>
        <div class="ml-2">Processing...</div>
      </div>
    </button>
  </div>
</template>
