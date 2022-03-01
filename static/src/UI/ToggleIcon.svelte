<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  export let toggledMode = false;
  // export let title = null;

  let toggleEditIcon = toggledMode
    ? "gg-toggle-square-off"
    : "gg-toggle-square";

  const toggled = () => {
    dispatch("toggled");
    toggledMode = !toggledMode;
    if (toggledMode) {
      toggleEditIcon = "gg-toggle-square-off";
    } else {
      toggleEditIcon = "gg-toggle-square";
    }
  };
</script>

<style>
  .gg-toggle-square {
    box-sizing: border-box;
    position: relative;
    display: block;
    transform: scale(var(--ggs, 1));
    width: 24px;
    height: 14px;
    border: 2px solid;
    border-radius: 2px;
  }
  .gg-toggle-square::before {
    content: "";
    display: block;
    box-sizing: border-box;
    position: absolute;
    width: 6px;
    height: 6px;
    background: currentColor;
    top: 2px;
    left: 2px;
    border-radius: 1px;
  }

  .gg-toggle-square-off {
    box-sizing: border-box;
    position: relative;
    display: block;
    transform: scale(var(--ggs, 1));
    width: 24px;
    height: 14px;
    border: 2px solid;
    border-radius: 2px;
  }
  .gg-toggle-square-off::before {
    content: "";
    display: block;
    box-sizing: border-box;
    position: absolute;
    width: 6px;
    height: 6px;
    background: currentColor;
    top: 2px;
    left: 12px;
    border-radius: 1px;
  }
</style>

<!-- {title != null && `${title}: `} -->
<i class={toggleEditIcon} on:click={toggled} />
