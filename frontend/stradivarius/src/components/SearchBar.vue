<script lang="ts" setup>
import { ref } from 'vue';
import { useFetchingStore } from '../stores/fetching';

const fetchingStore = useFetchingStore();

const value = ref<string>('');

const dicColumnNames = {
  ticker: 'Ticker',
  target_from: 'Target From',
  target_to: 'Target To',
  company: 'Company',
  action: 'Action',
  brokerage: 'Brokerage',
  rating_from: 'Rating From',
  rating_to: 'Rating To',
  datereleased: 'Date',
};
</script>

<template>
  <div class="flex justify-center items-center gap-3 mt-4 flex-col md:flex-row">
    <div class="flex overflow-hidden rounded-md bg-gray-200 focus:outline focus:outline-blue-500">
      <input
        type="text"
        placeholder="Search"
        class="w-full rounded-bl-md rounded-tl-md bg-gray-100 px-4 py-2 text-gray-700 focus:outline-blue-500"
        v-model="value"
      />
      <button
        class="bg-blue-500 px-3.5 text-white duration-150 hover:bg-blue-600"
        @click="fetchingStore.execSearch(value)"
      >
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
    <div class="flex flex-col items-center" v-if="fetchingStore.searchFilter">
      <!-- Search Tag -->
      <div
        class="relative inline-flex w-max items-center border font-sans font-medium rounded-md text-sm p-0.5 shadow-sm bg-stone-500 border-stone-500 text-stone-50"
      >
        <span class="font-sans text-current leading-none my-1 mx-2.5"
          >Searching by: '{{ fetchingStore.searchFilter }}'</span
        >
        <button
          class="grid place-items-center shrink-0 rounded-full p-px -translate-x-1 ms-1 w-5 h-5 stroke-2 cursor-pointer"
          @click="fetchingStore.execSearch('')"
        >
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            color="currentColor"
            class="h-full w-full"
          >
            <path
              d="M6.75827 17.2426L12.0009 12M17.2435 6.75736L12.0009 12M12.0009 12L6.75827 6.75736M12.0009 12L17.2435 17.2426"
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
            ></path>
          </svg>
        </button>
      </div>
    </div>

    <div class="flex flex-col items-center" v-if="fetchingStore.sortFilter.columnName">
      <!-- Sort Tag -->
      <div
        class="relative inline-flex w-max items-center border font-sans font-medium rounded-md text-sm p-0.5 shadow-sm bg-stone-500 border-stone-500 text-stone-50"
      >
        <span class="font-sans text-current leading-none my-1 mx-2.5"
          >Sorting by: '{{ dicColumnNames[fetchingStore.sortFilter.columnName] }}'</span
        >
        <button
          class="grid place-items-center shrink-0 rounded-full p-px -translate-x-1 ms-1 w-5 h-5 stroke-2 cursor-pointer"
          @click="fetchingStore.execSort('')"
        >
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            color="currentColor"
            class="h-full w-full"
          >
            <path
              d="M6.75827 17.2426L12.0009 12M17.2435 6.75736L12.0009 12M12.0009 12L6.75827 6.75736M12.0009 12L17.2435 17.2426"
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
            ></path>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>
