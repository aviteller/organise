<script>
  export let type = "button";
  export let disabled = false;
  export let href = null;
  export let mode = null;
  export let color = null;
  export let target = null;
</script>

<style>
  button,
  a {
    font-family: brandon-grotesque, sans-serif;
    font-weight: 600;
    border: 1px solid #14a4e3;
    background: #14a4e3;
    padding: 0.5rem 1rem;
    color: white;
    border-radius: 0;
    box-shadow: 1px 1px 3px rgba(0, 0, 0, 0.26);
    cursor: pointer;
    text-decoration: none;
    margin-top: 0.5rem;
  }

  button:focus {
    outline: none;
  }

  button:hover,
  button:active,
  a:hover,
  a:active {
    background: #5bc0de;
    border-color: #5bc0de;
    box-shadow: 1px 1px 8px rgba(77, 51, 51, 0.26);
  }

  button:disabled,
  button:disabled:hover,
  button:disabled:active {
    background: #ccc;
    border-color: #ccc;
    color: #959595;
    box-shadow: none;
    cursor: not-allowed;
  }

  .success {
    background: #01a129;
    border-color: #01a129;
  }

  .success:hover,
  .success:active {
    background: #1ac745;
    border-color: #1ac745;
  }

  .danger {
    background: #d9534f;
    border-color: #d9534f;
  }

  .danger:hover,
  .danger:active {
    background: #cd2a3e ;
    border-color: #cd2a3e ;
  }

  .outline {
    background: transparent;
    color: #14a4e3;
    box-shadow: none;
  }

  .outline:hover,
  .outline:active {
    background: #ffc7de;
    box-shadow: none;
  }

  .outline:disabled,
  .outline:disabled:hover,
  .outline:disabled:active {
    background: transparent;
    color: #ccc;
  }

  .outline.success {
    border-color: #01a129;
    color: #01a129;
  }

  .outline.success:hover,
  .outline.success:active {
    background: #c2ffd1;
  }
</style>

{#if href}
  <a href={href} {target}  class="{mode} {color}">
    <slot />
  </a>
{:else}

  <button class="{mode} {color}" {type} on:click {disabled}>
    <slot />
  </button>
{/if}
