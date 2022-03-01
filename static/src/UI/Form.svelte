<script>
  import { onMount, createEventDispatcher } from "svelte";
  export let title = null;
  export let className= "null";

  const dispatch = createEventDispatcher()
  
</script>
<div class="form-container-full">
  {#if title}
  <h1 class="form-title">{title}</h1>
  {/if}
  
  <form class={`${className} form-style`} on:submit|preventDefault>
    <slot name="inputs"/>
    <slot name="buttons"/>
  </form>
</div>


<style>

.form-container {
  width: 30%;
  margin-left: 35%;
  text-align: center;
  background-color: antiquewhite;
}
.form-container-full {

  text-align: center;
  background-color: antiquewhite;
}

.form-title {
  padding-top:20px;
  height: 55px;
  color: antiquewhite;
  background-color: #0a2240;
  text-transform: uppercase;
  text-align: center;
}

.form-style {
  /* width: 30%;
  margin-left: 35%;
  text-align: center; */
}
</style>