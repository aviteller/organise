<script>
  import { createEventDispatcher, onMount } from "svelte";

  const dispatch = createEventDispatcher();

  export let headers;
  export let classes = [];
  export let pagination = false;
  export let sortBy = null;


  import SortTab from "./SortTab.svelte";

  const setSort = (sort) => {
    if (sort.charAt(0) === "-") {
      sortBy = {
        name: sort.slice(1),
        dir: "desc",
      };
    } else {
      sortBy = {
        name: sort,
        dir: "asc",
      };
    }
  };

  const changeSort = (e) => {
    let name = e.detail.name;
    let dir = e.detail.dir;
    dispatch("changesort", { name, dir });
  };
  onMount(() => {
    if (sortBy !== null) setSort(sortBy);
  });
</script>

<table class={classes.length > 0 ? classes.join(" ") : ""}>
  <thead>
    <tr>
      {#each headers as header}
        {#if typeof header == "object"}
          {#if header.col && header.sortable}
            {#if sortBy.name === header.name}
              <th colspan={header.col}>
                >
                <SortTab
                  on:changesort={changeSort}
                  name={header.name}
                  dir={sortBy.dir === "desc" ? "desc" : "asc"}
                  active={true}
                />
              </th>
            {:else}
              <th colspan={header.col}>
                <button
                  on:click|preventDefault={() =>
                    changeSort(header.name, "desc")}
                >
                  {header.name}
                </button>
              </th>
            {/if}
          {:else if header.sortable}
            {#if sortBy.name === header.name}
              <th>
                <SortTab
                  on:changesort={changeSort}
                  name={header.name}
                  dir={sortBy.dir === "desc" ? "desc" : "asc"}
                  active={true}
                />
              </th>
            {:else}
              <th>
                <SortTab on:changesort={changeSort} name={header.name} />
              </th>
            {/if}
          {:else if header.col}
            <th colspan={header.col}>{header.name}</th>
          {/if}
        {:else}
          <th>{header}</th>
        {/if}
      {/each}
    </tr>
  </thead>

  <slot name="body" />

  {#if pagination}
    <slot name="pagination" />
  {/if}
</table>

<style>
  table {
    border-collapse: collapse;
    width: 100%;
  }
  th {
    background-color: #f2f2f2;
  }

  th {
    text-align: center;
    padding: 8px;
    border-bottom: 1px black solid;
  }

  tr:nth-child(even) {
    background-color: #f2f2f2;
  }

  table :global(tr:hover) {
    background-color: grey !important;
  }

  table :global(tr:nth-child(even)) {
    background-color: #f2f2f2;
  }

  table :global(td) {
    border-bottom: 1px solid #ddd;
  }

  :global(.half) {
    width: 50%!important;
  }

  :global(.full) {
    width: 100%;
  }

  :global(.blue) {
    background-color: blue;
  }
</style>
