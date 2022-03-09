<script>
  import { createEventDispatcher, onMount } from "svelte";

  const dispatch = createEventDispatcher();
  export let todo;

  const priorities = [
    {
      name: "High",
      value: 1,
    },
    {
      name: "Medium",
      value: 2,
    },
    {
      name: "Low",
      value: 3,
    },
  ];

  let loaded = true;
  const updateTodo = () => {
    let dataToSend = {
      id: todo.id,
      content: todo.content.replace(/(<([^>]+)>)/gi, ""),
      priority: todo.priority,
      complete: todo.complete,
    };

    dispatch("updatetodo", dataToSend);
  };

  // console.log(priorities.filter(p => p.value == todo.priority))
</script>

{#if loaded}
  <div
    class={`todo-item ${todo.complete ? "complete" : ""} ${
      priorities.filter((p) => p.value == todo.priority)[0].name
    }`}
  >
    {#if todo.complete}
      <p>{todo.content}</p>
    {:else}
      <p
        contenteditable="true"
        bind:innerHTML={todo.content}
        on:blur={() => updateTodo()}
      ></p>
    {/if}
    <div class="todo-item-controls">
      <button
        class="todo-item-controls-button"
        on:click={() => {
          todo.complete = !todo.complete;
          updateTodo();
        }}
      >
        {todo.complete ? "Undo" : "Complete"}
      </button>
      <button
        class="todo-item-controls-button"
        on:click={() => {
          todo.priority = todo.priority == 3 ? 1 : todo.priority + 1;
          updateTodo();
        }}
      >
        {priorities.filter((p) => p.value == todo.priority)[0].name}
      </button>
      <!-- delete button -->

      <button
        class="todo-item-controls-button"
        on:click={() => {
          dispatch("deletetodo", todo.id);
        }}>Delete</button
      >
    </div>
  </div>
{/if}

<style>
  [contenteditable] {
    padding: 0.5em;
    border: 1px solid #eee;
    border-radius: 4px;
    background-color: #fff;
    width: 500px;
    white-space: initial;

    word-wrap: break-word;
  }

  .todo-item {
    border-width: 1px;
    border-style: solid;
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    margin-bottom: 10px;
    padding-top: 10px;
  }

  .todo-item > p {
    white-space: initial;
    word-wrap: break-word;
    max-width: 500px;
  }

  :global(.High) {
    border-color: #ff0000;
    background-color: rgba(255, 0, 0, 0.1);
  }
  :global(.Medium) {
    border-color: #ffa500;
    background-color: rgba(255, 165, 0, 0.1);
  }

  :global(.Low) {
    border-color: #00ff00;
    background-color: rgba(0, 255, 0, 0.1);
  }

  :global(.complete) {
    text-decoration: line-through;
  }
</style>
