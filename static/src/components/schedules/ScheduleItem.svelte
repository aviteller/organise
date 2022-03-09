<script>
  import { createEventDispatcher, onMount } from "svelte";

  const dispatch = createEventDispatcher();
  export let item;
  export let catAndFreq = null;
  export let control = false;


  let cat = "";
  let freq = "";
  if (catAndFreq) {
    if (item.schedule_cat_id) {
      cat = catAndFreq.categories.filter(
        (c) => item.schedule_cat_id == c.value
      )[0].text;
    }

    if (item.schedule_frequency_id) {
      freq = catAndFreq.frequencies.filter(
        (f) => item.schedule_frequency_id == f.value
      )[0].text;
    }
  } else {
    cat = item.schedule_cat_id;
    freq = item.schedule_frequency_id;
  }
  let loaded = true;
</script>

{#if loaded}
  <td>{item.desc}</td>
  <td>{item.schedule_time}</td>
  <td>{item.start_date ?? null}</td>
  <td>{item.end_date ?? null}</td>
  <td>{cat}</td>
  <td>{freq}</td>
  {#if control}
  <td
    ><button
      on:click={() => {
        dispatch("deleterow", item.id);
      }}>delete</button
    ></td
  >
  <td
    ><button
      on:click={() => {
        dispatch("editrow", item);
      }}>edit</button
    ></td
  >
  {/if}
{/if}

<style>
  .row {
    display: flex;
    /* flex-direction: row; */
    justify-content: space-between;
  }
</style>
