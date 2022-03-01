<script>
  export let label;
  export let colour;
  export let data;
</script>

<div class={`box`} style={`background-color:${colour}`}>
  <div class="data">{data}</div>
  <div class="label">{label}</div>
</div>

<style>
  .box {
    width: 300px;
    height: 300px;
    color: #000;
  }

  .data {
    padding-top: 2rem;
    padding-bottom: 2rem;
    color:lightgray;
    font-size: 40px;
  }

  .label {
    padding-top: 1rem;
    padding-bottom: 2rem;
    word-spacing: 100vw; 
    font-size: 30px;
  }
</style>
