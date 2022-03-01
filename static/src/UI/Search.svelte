<script>
  import { createEventDispatcher } from "svelte";
  export let searchValue = "";
  export let label = "Search";
  export let placeholder = "Enter Search";
  const dispatch = createEventDispatcher();

  const getAllSearch = () => dispatch("getallsearch", searchValue);
</script>

<div>
  <label for="search">{label}</label>
  <input
    id="search"
    type="text"
    bind:value={searchValue}
    on:keyup={(e) => {
      e.key !== "Enter" || getAllSearch();
    }}
    {placeholder}
  />
  <button on:click={getAllSearch}>GO</button>
</div>
