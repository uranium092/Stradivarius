<script lang="ts" setup>
import { onMounted } from 'vue';
import { useFetchingStore } from '../stores/fetching';

const fetchingStore = useFetchingStore();

onMounted(() => {
  fetchingStore.changeMode('all');
});
</script>

<template>
  <!-- component -->
  <div class="flex flex-col">
    <div class="overflow-x-auto sm:mx-0.5 lg:mx-0.5">
      <div class="inline-block min-w-full h-20 sm:px-6 lg:px-8">
        <div class="overflow-hidden h-[70vh] overflow-y-auto">
          <table class="min-w-full">
            <thead class="bg-white sticky top-0 shadow-md">
              <tr>
                <th scope="col" class="text-sm font-medium text-gray-900 px-4 py-2 text-center">
                  <span class="cursor-pointer" @click="fetchingStore.execSort('ticker')"
                    >Ticker</span
                  >&nbsp;<span
                    class="text-lg"
                    v-if="fetchingStore.sortFilter.columnName === 'ticker'"
                    >{{ fetchingStore.sortFilter.sortType === 'ASC' ? '↑' : '↓' }}</span
                  >
                </th>
                <th scope="col" class="text-sm font-medium text-gray-900 px-4 py-2 text-center">
                  <span class="cursor-pointer" @click="fetchingStore.execSort('target_from')"
                    >Target from</span
                  >&nbsp;<span
                    class="text-lg"
                    v-if="fetchingStore.sortFilter.columnName === 'target_from'"
                    >{{ fetchingStore.sortFilter.sortType === 'ASC' ? '↑' : '↓' }}</span
                  >
                </th>
                <th scope="col" class="text-sm font-medium text-gray-900 px-4 py-2 text-center">
                  <span class="cursor-pointer" @click="fetchingStore.execSort('target_to')"
                    >Target to</span
                  >&nbsp;<span
                    class="text-lg"
                    v-if="fetchingStore.sortFilter.columnName === 'target_to'"
                    >{{ fetchingStore.sortFilter.sortType === 'ASC' ? '↑' : '↓' }}</span
                  >
                </th>
                <th scope="col" class="text-sm font-medium text-gray-900 px-4 py-2 text-center">
                  <span class="cursor-pointer" @click="fetchingStore.execSort('company')"
                    >Company</span
                  >&nbsp;<span
                    class="text-lg"
                    v-if="fetchingStore.sortFilter.columnName === 'company'"
                    x
                    >{{ fetchingStore.sortFilter.sortType === 'ASC' ? '↑' : '↓' }}</span
                  >
                </th>
                <th scope="col" class="text-sm font-medium text-gray-900 px-4 py-2 text-center">
                  <span class="cursor-pointer" @click="fetchingStore.execSort('action')"
                    >Action</span
                  >&nbsp;<span
                    class="text-lg"
                    v-if="fetchingStore.sortFilter.columnName === 'action'"
                    >{{ fetchingStore.sortFilter.sortType === 'ASC' ? '↑' : '↓' }}</span
                  >
                </th>
                <th scope="col" class="text-sm font-medium text-gray-900 px-4 py-2 text-center">
                  <span class="cursor-pointer" @click="fetchingStore.execSort('brokerage')"
                    >Brokerage</span
                  >&nbsp;<span
                    class="text-lg"
                    v-if="fetchingStore.sortFilter.columnName === 'brokerage'"
                    >{{ fetchingStore.sortFilter.sortType === 'ASC' ? '↑' : '↓' }}</span
                  >
                </th>
                <th scope="col" class="text-sm font-medium text-gray-900 px-4 py-2 text-center">
                  <span class="cursor-pointer" @click="fetchingStore.execSort('rating_from')"
                    >Rating from</span
                  >&nbsp;<span
                    class="text-lg"
                    v-if="fetchingStore.sortFilter.columnName === 'rating_from'"
                    >{{ fetchingStore.sortFilter.sortType === 'ASC' ? '↑' : '↓' }}</span
                  >
                </th>
                <th scope="col" class="text-sm font-medium text-gray-900 px-4 py-2 text-center">
                  <span class="cursor-pointer" @click="fetchingStore.execSort('rating_to')"
                    >Rating to</span
                  >&nbsp;<span
                    class="text-lg"
                    v-if="fetchingStore.sortFilter.columnName === 'rating_to'"
                    >{{ fetchingStore.sortFilter.sortType === 'ASC' ? '↑' : '↓' }}</span
                  >
                </th>
              </tr>
            </thead>

            <tbody>
              <tr class="bg-gray-100" v-for="i in fetchingStore.tableData" v-bind:key="i.Id">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {{ i.ticker }}
                </td>
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ i.target_from }}
                </td>
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ i.target_to }}
                </td>
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ i.company }}
                </td>
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ i.action }}
                </td>
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ i.brokerage }}
                </td>
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ i.rating_from }}
                </td>
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ i.rating_to }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
    <nav class="self-end mr-10">
      <ul class="inline-flex items-center -space-x-px">
        <li>
          <span
            class="block py-2 px-3 ml-0 leading-tight text-gray-500 bg-white rounded-l-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700"
            @click="fetchingStore.changeSectionPagination('prev')"
          >
            <svg
              class="w-5 h-5"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
                clip-rule="evenodd"
              ></path>
            </svg>
          </span>
        </li>
        <li v-for="n in fetchingStore.getItemsRenderPagination" :key="n">
          <span
            :class="[
              'py-2 px-3 leading-tight border cursor-pointer',
              n === fetchingStore.pagination.currentPage
                ? 'text-blue-700 bg-blue-200 border border-blue-300 hover:bg-blue-100 hover:text-blue-700'
                : 'text-gray-500 bg-white border-gray-300 hover:bg-gray-100 hover:text-gray-700',
            ]"
            @click="fetchingStore.changePage(n)"
            >{{ n }}</span
          >
        </li>
        <li>
          <span
            class="block py-2 px-3 leading-tight text-gray-500 bg-white rounded-r-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700"
            @click="fetchingStore.changeSectionPagination('next')"
          >
            <svg
              class="w-5 h-5"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                clip-rule="evenodd"
              ></path>
            </svg>
          </span>
        </li>
      </ul>
    </nav>
  </div>
</template>
