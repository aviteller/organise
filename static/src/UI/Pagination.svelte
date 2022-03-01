<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  const switchPage = (page) => dispatch("switchpage", page);

  export let pagination;
</script>

{#if pagination.totalPages > 1}
  {#if pagination.prevPage}
    <button on:click|preventDefault={() => switchPage(pagination.prevPage)}>
      prev
    </button>
  {/if}
  {#each Array(pagination.totalPages) as _, page}
    {#if pagination.currentPage == page + 1}
      <button style="color:red">{page + 1}</button>
    {:else}
      <button on:click|preventDefault={() => switchPage(page + 1)}>
        {page + 1}
      </button>
    {/if}
  {/each}
  {#if pagination.nextPage}
    <button on:click|preventDefault={() => switchPage(pagination.nextPage)}>
      next
    </button>
  {/if}
  <button on:click|preventDefault={() => switchPage("All")}>ALL</button>
{/if}
<span>
  <!-- TO BE SERIOUSLY LOOKED AT  -->
  results {pagination.currentPage == 1
    ? 1
    : pagination.limit * pagination.currentPage - pagination.limit + 1}
  - {pagination.current * pagination.limit >= pagination.totalRows
    ? pagination.totalRows
    : pagination.currentPage * pagination.limit < pagination.totalRows
    ? pagination.currentPage * pagination.limit
    : pagination.totalRows}
  out of {pagination.totalRows}
</span>
