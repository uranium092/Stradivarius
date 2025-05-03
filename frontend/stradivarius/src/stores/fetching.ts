import { defineStore } from 'pinia';
import { computed, reactive, ref } from 'vue';

interface ItemStock {
  ticker: string;
  targetFrom: string;
  targetTo: string;
  company: string;
  action: string;
  brokerage: string;
  ratingFrom: string;
  ratingTo: string;
}

interface Pagination {
  currentPage: number;
  totalPages: number;
  currentSection: number;
}

const PAGSPERSECTION: number = 5;

export const useFetchingStore = defineStore('fetching', () => {
  //state
  const mode = ref<string>('all');
  const isLoading = ref<boolean>(false);
  const tableData = ref<ItemStock[]>([]);
  const pagination = reactive<Pagination>({
    currentPage: 6,
    totalPages: 7,
    currentSection: 1,
  });

  //getters
  const getItemsRenderPagination = computed<number[]>(() => {
    //limit pages navigate
    const pageLimitBySection: number = pagination.currentSection * PAGSPERSECTION;
    const from: number = pageLimitBySection - (PAGSPERSECTION - 1);
    const to: number = Math.min(pagination.totalPages, pageLimitBySection);
    return Array.from({ length: to - from + 1 }, (_, index) => index + from);
  });

  //actions
  const changeSectionPagination = (direction: string): void => {
    //limit section navigate
    const newSection: number =
      direction === 'prev' ? pagination.currentSection - 1 : pagination.currentSection + 1;
    if (newSection <= 0) {
      return;
    }
    if (newSection * PAGSPERSECTION - (PAGSPERSECTION - 1) > pagination.totalPages) {
      return;
    }
    pagination.currentSection = newSection;
  };

  const changePage = (page: number): void => {
    pagination.currentPage = page;
  };

  const changeMode = (modeType: string): void => {
    mode.value = modeType;
  };

  return {
    mode,
    isLoading,
    tableData,
    pagination,
    changeSectionPagination,
    getItemsRenderPagination,
    changePage,
    changeMode,
  };
});

// query
