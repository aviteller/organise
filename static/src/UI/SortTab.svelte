<script>
  import { createEventDispatcher, onMount } from "svelte";
  const dispatch = createEventDispatcher();
  export let name;
  export let dir = null;
  export let active = false;

  const changeSort = (name, dir) => {
    // console.log("name :>> ", name, dir);
    dispatch("changesort", { name, dir });
  };
</script>

<style>
  .table-sort {
    background-color: blueviolet;
  }

  div {
    text-align: center;
  }

  .sort-active {
    background-color: aquamarine !important;
  }
</style>

<div
  class={`table-sort ${active && 'sort-active'}`}
  on:click={() => changeSort(name, dir === 'asc' ? 'desc' : 'asc')}>
  {name} {dir !== null ? (dir === 'desc' ? 'desc' : 'asc') : ''}
</div>
