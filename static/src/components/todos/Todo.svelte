<script>
  import { onMount } from "svelte";
  import TodoItem from "./TodoItem.svelte";
  import { Form, Input, Modal, Button, Select } from "../../UI/UI.svelte";

  export let person_id;

  let todos = [];
  let loaded = false;
  let displayModal = false;
  let newTodo = {
    content: "",
    person_id: person_id,
    priority: 3,
  };

  let priorities = [
    {
      value: 3,
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
    if (data.todos) todos = data.todos;

    loaded = true;
    // return data;
  };
  const showModal = () => (displayModal = !displayModal);
  const submitAddTodo = () => {
    console.log(newTodo);
    fetch("/api/todos", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newTodo),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data.row);
        todos = [...todos, data.row];
        newTodo = {
          content: "",
          person_id: person_id,
          priority: 3,
        };

        todos = todos.sort((a, b) => a.priority - b.priority);
        todos = todos.sort((a, b) => a.complete - b.complete);
        // getTodos();
        showModal();
      });
  };

  const updateTodo = (e) => {
    // console.log(e.detail);
    let todo = e.detail;
    fetch(`/api/todos/${todo.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(todo),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        // sort by completed
        todos = todos.sort((a, b) => a.priority - b.priority);
        todos = todos.sort((a, b) => a.complete - b.complete);
      });

  };

  const deleteTodo = (e) => {
    
    let id = e.detail;
    fetch(`/api/todos/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        todos = todos.filter((t) => t.id !== id);
      });
  };

  onMount(async () => {
    await getTodos();
  });
</script>

<Button on:click={showModal}>Add Todo</Button>
<!-- <Button on:click={() => getTodos("complete")}>Add Todo</Button> -->

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
          selectedOption={newTodo.priority}
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
    <TodoItem {todo} on:updatetodo={updateTodo} on:deletetodo={deleteTodo} />
  {/each}
{:else}
  <p>No todos yet</p>
{/if}
