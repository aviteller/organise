<script>
  import ScheduleItem from "../schedules/ScheduleItem.svelte";
  import {onMount} from "svelte";
  // import
  import {catfre} from "../../catfre.store"
  import { Table, Modal } from "../../UI/UI.svelte";
  export let day;
  export let scheduleItems;
  let loaded = false;

  let showModal =   false;

  console.log($catfre);

  // make day suffix

  const nth = (n) =>
    `${day}${[, "st", "nd", "rd"][(n / 10) % 10 ^ 1 && n % 10] || "th"}`;

  
  // make day of week
</script>

<div class="box">
  {nth(day)}
  {#if scheduleItems.length > 0}
    {#each scheduleItems as item}
      <div>
        <div>{item.desc} - {item.schedule_time}</div>
      </div>
    {/each}
  {/if}
</div>

<style>
  div.box {
    min-height: 50px;
    width: 80%;
    border: 1px solid black;
    text-align: center;
  }
</style>
