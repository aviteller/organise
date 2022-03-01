<script>
  import { onMount } from "svelte";
  import Todo from "../todos/Todo.svelte";
  import { Form, Input, Modal, Button,Select } from "../../UI/UI.svelte";

  export let person_id;

  let todos = [];
  let loaded = false;
  let displayModal = false;
  let newTodo = {
    content: "",
    person_id: person_id,
    priority: "",
  };

  let priorities = [
    {
      value: 1,
      text: "Low",
    },
    {
      value: 2,
      text: "Medium",
    },
    {
      value: 1,
      text: "High",
    },
  ];

  const getTodos = async () => {
    loaded = false;
    const response = await fetch(`/api/todos/${person_id}`);
    const data = await response.json();
    todos = data.todos;
    console.log(todos);
    loaded = true;
    // return data;
  };
  const showModal = () => (displayModal = !displayModal);
  const submitAddTodo = () => {

  }
  onMount(async () => {
    await getTodos();
  });
</script>


<Button on:click={showModal}>Add Todo</Button>

{#if displayModal}
<Modal title={"Add Todo"} on:cancel={showModal}>
  <Form>
    <div slot="inputs">
      <Input
        id="desc"
        label="Content"
        bind:value={newTodo.content}
        on:input={(e) => (newTodo.content = e.target.value)}
      />
  

      <Select
        options={priorities}
        withSelectAsFirstOption={true}
        label={"assign priority"}
        on:change={(e) => (newTodo.priority = e.target.value)}
      />
 
    </div>
    <div slot="buttons">
      <Button on:click={submitAddTodo}>submit</Button>
    </div>
  </Form>
</Modal>
{/if}

{#if loaded && todos.length > 0}
  {#each todos as todo}
    <Todo {todo} />
  {/each}
{:else}
  <p>No todos yet</p>
{/if}
