<script>
  import Modal from "./Modal.svelte";
  export let goBack;
  export let message;

</script>

<Modal title="An error occured!" on:cancel goback={goBack}>
  <p>{message}</p>
</Modal>