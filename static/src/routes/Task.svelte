<script>
  export let params = {}


  import { onMount } from "svelte";
 
  import config from "../config";
  let loaded = false

  let task = {}
  const getTask = () => {
    fetch(`${config.apiUrl}tasks/${params.id}`)
      .then((res) => res.json())
      .then((res) => {
        task = res.task
        loaded = true
      });
  };
  
  onMount(() => {
    getTask()
  })
</script>

{#if loaded}
  <a href={`/#/people/${task.person_id}`}>go back to parent</a>
  <br>
  {task.desc}
{/if}