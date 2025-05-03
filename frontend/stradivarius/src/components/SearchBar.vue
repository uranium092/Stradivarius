<script lang="ts" setup></script>

<template>
  <div class="flex justify-center items-center gap-3 mt-4">
    <div class="flex overflow-hidden rounded-md bg-gray-200 focus:outline focus:outline-blue-500">
      <input
        type="text"
        placeholder="Search"
        class="w-full rounded-bl-md rounded-tl-md bg-gray-100 px-4 py-2 text-gray-700 focus:outline-blue-500"
      />
      <button class="bg-blue-500 px-3.5 text-white duration-150 hover:bg-blue-600">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="size-6"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"
          />
        </svg>
      </button>
    </div>
  </div>
</template>
