import { defineStore } from 'pinia';
import { computed, reactive, ref } from 'vue';

interface Pagination {
  currentPage: number;
  totalPages: number;
  currentSection: number;
}

interface Filters {
  search: string;
  sort: string;
}

interface ItemStock {
  Id: string;
  ticker: string;
  target_from: string;
  target_to: string;
  company: string;
  action: string;
  brokerage: string;
  rating_from: string;
  rating_to: string;
  time: string;
}

interface JSONResponse {
  dataStock: ItemStock[];
  totalPages: number;
}

const PAGSPERSECTION: number = 5;

export const useFetchingStore = defineStore('fetching', () => {
  //state
  const mode = ref<string>('');
  const isLoading = ref<boolean>(false);
  const hasError = ref<boolean>(false);
  const tableData = ref<ItemStock[]>([]);
  const pagination = reactive<Pagination>({
    currentPage: 1,
    totalPages: 0,
    currentSection: 1,
  });
  const filters = reactive<Filters>({
    search: '',
    sort: '',
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
  const setError = (): void => {
    hasError.value = true;
    setTimeout(() => {
      hasError.value = false;
    }, 2000);
  };

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

  //independient
  const fetchDataTable = async ({
    path,
    page,
    search,
    sort,
  }: {
    path?: string;
    page?: number;
    search?: string;
    sort?: string;
  }): Promise<number> => {
    isLoading.value = true;
    try {
      const pathReq: string = path ? path : mode.value;
      const pageReq: number = page ? page : pagination.currentPage;
      const searchReq: string = search ? search : filters.search;
      const sortReq: string = sort ? sort : filters.sort;

      let baseUrl: string = `http://localhost:8080/api/stock/${pathReq}?page=${pageReq}`;

      if (searchReq != '') {
        baseUrl += `&search=${searchReq}`;
      }
      if (sortReq != '') {
        baseUrl += `&sort=${sortReq}`;
      }
      const fetched: Response = await fetch(baseUrl);
      const { dataStock, totalPages }: JSONResponse = await fetched.json();
      pagination.totalPages = totalPages;
      tableData.value = dataStock;
      return totalPages;
    } catch (err: any) {
      setError();
      return 0;
    } finally {
      isLoading.value = false;
    }
  };

  const changePage = async (page: number): Promise<void> => {
    if (page === pagination.currentPage) {
      return;
    }
    await fetchDataTable({ page });
    pagination.currentPage = page;
  };

  const changeMode = async (modeType: string, init?: boolean): Promise<void> => {
    const payload: { path: string; page: number } = { path: `${modeType}`, page: 1 };
    await fetchDataTable(payload);
    pagination.currentPage = 1;
    pagination.currentSection = 1;
    mode.value = modeType;
  };

  return {
    mode,
    filters,
    isLoading,
    hasError,
    tableData,
    pagination,
    changeSectionPagination,
    getItemsRenderPagination,
    changePage,
    changeMode,
  };
});
