<script setup lang="ts">
import Loading from './components/Loading.vue';
import Mode from './components/Mode.vue';
import SearchBar from './components/SearchBar.vue';
import DataTable from './components/DataTable.vue';
import Error from './components/Error.vue';
</script>

<template>
  <div>
    <Loading />
    <Error />
    <Mode />
    <SearchBar />
    <DataTable />
  </div>
</template>
