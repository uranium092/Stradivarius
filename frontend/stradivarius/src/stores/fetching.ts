import { defineStore } from 'pinia';
import { computed, reactive, ref } from 'vue';

// pagination managment
interface Pagination {
  currentPage: number;
  totalPages: number;
  currentSection: number;
}

// HTTP request concats both keys with '$'; columnName$sortType
interface SortFilter {
  columnName: string;
  sortType: string;
}

// Model for each item from HTTP Response, and <tr> of table
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

// HTTP Response
interface JSONResponse {
  dataStock: ItemStock[];
  totalPages: number;
}

// pass params if it is necessary based on event
interface HTTPFetchData {
  path?: string;
  page?: number;
  search?: string;
  sort?: string;
}

const PAGSPERSECTION: number = 5;

export const useFetchingStore = defineStore('fetching', () => {
  //state
  const mode = ref<string>(''); // all | recommendation
  const isLoading = ref<boolean>(false);
  const hasError = ref<boolean>(false);
  const tableData = ref<ItemStock[]>([]);
  const pagination = reactive<Pagination>({
    currentPage: 1,
    totalPages: 0,
    currentSection: 1,
  });
  const searchFilter = ref<string>(''); // nothing to search
  const sortFilter = reactive<SortFilter>({
    // nothing to sort
    columnName: '',
    sortType: '',
  });

  // One section contains <= 5 pages

  //----- GETTERS

  const getItemsRenderPagination = computed<number[]>(() => {
    //limit pages navigate by section
    const pageLimitBySection: number = pagination.currentSection * PAGSPERSECTION;
    const from: number = pageLimitBySection - (PAGSPERSECTION - 1);
    const to: number = Math.min(pagination.totalPages, pageLimitBySection);
    return Array.from({ length: to - from + 1 }, (_, index) => index + from);
  });

  //----- ACTIONS

  // HTTP Request on any event: changeMode, search, sort o pagination
  const fetchDataTable = async ({ path, page, search, sort }: HTTPFetchData): Promise<boolean> => {
    isLoading.value = true;
    try {
      // take the params and states and build HTTP Request Queries
      const pathReq: string = path ? path : mode.value;
      const pageReq: number = page ? page : pagination.currentPage;
      const searchReq: string = search !== undefined ? search : searchFilter.value; // search can come "" (clear filter)
      const sortReq: string =
        sort !== undefined // sort can come "" (clear filter)
          ? sort // keep empty param to delete filter
          : sortFilter.columnName.length > 0 //build sort query based on if this state is full or empty
          ? `${sortFilter.columnName}$${sortFilter.sortType}`
          : '';

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
      return true; // everything good
    } catch (err: any) {
      // something failed
      setError();
      return false;
    } finally {
      isLoading.value = false;
    }
  };

  const setError = (): void => {
    // Wrong HTTP Response
    hasError.value = true;
    setTimeout(() => {
      hasError.value = false;
    }, 2000);
  };

  const changeMode = async (modeType: string): Promise<void> => {
    if (modeType === mode.value) {
      return;
    }
    // path param to change endpoint; and page to begin from start
    const success: boolean = await fetchDataTable({ path: `${modeType}`, page: 1 });
    if (success) {
      // reflect changes if HTTP response is OK
      pagination.currentPage = 1;
      pagination.currentSection = 1;
      mode.value = modeType;
    }
  };

  const changeSectionPagination = (direction: string): void => {
    //limit section navigate
    const newSection: number =
      direction === 'prev' ? pagination.currentSection - 1 : pagination.currentSection + 1;
    // newSection out of range (0, -1, ...)
    if (newSection <= 0) {
      return;
    }
    // newSection out of range (120, 121, ...), with an e.g. 73 totalPages
    if (newSection * PAGSPERSECTION - (PAGSPERSECTION - 1) > pagination.totalPages) {
      return;
    }
    pagination.currentSection = newSection;
  };

  const changePage = async (page: number): Promise<void> => {
    if (page === pagination.currentPage) {
      return;
    }
    const success: boolean = await fetchDataTable({ page }); // only page param is necessary
    if (success) {
      // reflect changes if HTTP response is OK
      pagination.currentPage = page;
    }
  };

  const execSearch = async (value: string): Promise<void> => {
    if (value === searchFilter.value) {
      return;
    }
    // search param to apply filter; and page to begin from start
    const success: boolean = await fetchDataTable({ search: value, page: 1 });
    if (success) {
      // reflect changes if HTTP response is OK
      pagination.currentPage = 1;
      pagination.currentSection = 1;
      searchFilter.value = value;
    }
  };

  const execSort = async (columnName: string): Promise<void> => {
    let newSortType: string = 'ASC';
    if (columnName === sortFilter.columnName) {
      // filter apply on same column, so reverse sortType
      newSortType = sortFilter.sortType === 'ASC' ? 'DESC' : 'ASC';
    }
    if (columnName === '') {
      // this can come "" (clear filter)
      newSortType = ''; // clear sortType too
    }
    // pass correct sort param; "columnName$sortType" or "" (clear filter)
    const success: boolean = await fetchDataTable({
      sort: columnName.length > 0 ? `${columnName}$${newSortType}` : '',
    });
    if (success) {
      // reflect changes if HTTP response is OK
      sortFilter.columnName = columnName;
      sortFilter.sortType = newSortType;
    }
  };

  return {
    mode,
    isLoading,
    hasError,
    tableData,
    pagination,
    searchFilter,
    sortFilter,
    getItemsRenderPagination,
    fetchDataTable,
    changeMode,
    changeSectionPagination,
    changePage,
    execSearch,
    execSort,
  };
});
