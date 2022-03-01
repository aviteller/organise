<script>
  export let options;
  export let selectedOption;
  export let label = "Select";
  export let withSelectAsFirstOption = true;
  export let valid = true;
  export let validityMessage = "";
  // options = [
  //   { value: 1, text: "1" },
  //   { value: 2, text: "2" },
  //   { value: 3, text: "3" },
  // ];

  // console.log(options)
  let touched = false;
</script>

<div class="form-control">
  <div id="label">{label}</div>
  <select class:invalid={!valid && touched} id="select" on:change on:blur={() => (touched = true)}>
    {#if withSelectAsFirstOption}
      <option value="0">Please Select</option>
    {/if}
    {#each options as option}
      <option
        value={option.value}
        selected={selectedOption == option.value ? "selected" : ""}
        >{option.text}</option
      >
    {/each}
  </select>
  {#if validityMessage && !valid && touched}
  <p class="error-message">{validityMessage}</p>
{/if}
</div>

<style>
  select{
    display: block;
    width: 100%;
    font: inherit;
    border: none;
    border-bottom: 2px solid #ccc;
    border-radius: 3px 3px 0 0;
    background: white;
    padding: 0.15rem 0.25rem;
    transition: border-color 0.1s ease-out;
  }

  select {
    text-align: center;
  }

  select:focus {
    border-color: #e40763;
    outline: none;
  }

  div#label {
    display: block;
    margin-bottom: 0.5rem;
    width: 100%;
  }

  .form-control {
    padding: 0.5rem 0;
    width: 100%;
    margin: 0.25rem 0;
  }

  .invalid {
    border-color: red;
    background: #fde3e3;
  }

  .error-message {
    color: red;
    margin: 0.25rem 0;
  }
</style>
