<script>
  export let params = {};

  import { onMount } from "svelte";
  import { Table, Button, Modal, Form, Input, Select } from "../UI/UI.svelte";
  import Todo from "../components/todos/Todo.svelte";
  import Schedule from "../components/schedules/Schedule.svelte";
  import config from "../config";
  import Cookies from "../Cookie";

  let showChildren = false;
  let showTasks = false;
  let showTodos = false;
  let showSchedule = false;

  let tabs = {
    children: false,
    todos: false,
    schedule: false,
  };

  let loaded = false;
  let person = {};
  let children = [];
  let tasks = [];

  let c = new Cookies();

  $: if (params.id) {
    //watch the params.id for changes
    loaded = false;
    getPerson(); //invoke your method to reload data
  }

  let taskTypes = [
    { value: 1, text: "birthdays" },
    { value: 2, text: "homework" },
  ];

  let selectPeople = [];

  const getPerson = () => {
    fetch(`${config.apiUrl}people/${params.id}`)
      .then((res) => res.json())
      .then((res) => {
        person = res.person;
        children = res.children;
        tasks = res.tasks;
        selectPeople = [];
        if (children) {
          selectPeople.push({ value: person.id, text: person.name });
          children.forEach((c) =>
            selectPeople.push({ value: c.id, text: c.name })
          );
        }

        loaded = true;
      });
  };

  var now = new Date();
  now.setMinutes(now.getMinutes() - now.getTimezoneOffset());
  let nowFormatted = now.toISOString().slice(0, 11) + "00:00";
  let newTask = {
    desc: "",
    person_id: 0,
    due_datetime: nowFormatted,
    task_type_id: 0,
  };

  let displayModal = false;
  const showModal = () => (displayModal = !displayModal);
  const submitAddTask = () => {
    //check if newTask.person_id is valid if not assign to params.id
    if (newTask.person_id == 0) {
      newTask.person_id = params.id;
    }

    fetch(`${config.apiUrl}tasks`, {
      method: "POST",
      body: JSON.stringify(newTask),
    })
      .then((res) => res.json())
      .then((res) => {
        showModal();
        console.log(res);
      });
  };

  const checkTabCookies = () => {
    for (let tab in tabs) {
      if (c.getCookie(`tabs-${[tab]}`)) {
        tabs[tab] = true;
      }
    }
  };

  const updateTab = (tab) => {
    
    for (let key in tabs) {
      // console.log(key, tab);
      if (key == tab) {

        if(tabs[key]){
          tabs[key] = false;
          // c.eraseCookie("tabs");
          c.eraseCookie(`tabs-${key}`);
        }else{
          c.setCookie(`tabs-${[key]}`, JSON.stringify({ [key]: true }));
          tabs[key] = true;
        }


        // tabs[key] = !tabs[key];
      }
      // console.log(tabs);
    }
  };

  onMount(() => {
    getPerson();
    checkTabCookies();
  });
</script>

{#if loaded}
  {#if displayModal}
    <Modal title={"Add task"} on:cancel={showModal}>
      <Form>
        <div slot="inputs">
          <Input
            id="desc"
            label="Description"
            bind:value={newTask.desc}
            type="textarea"
          />
          {#if selectPeople.length >= 1}
            <Select
              options={selectPeople}
              withSelectAsFirstOption={true}
              label={"assign task too"}
              on:change={(e) => (newTask.person_id = e.target.value)}
            />
          {/if}
          <Select
            options={taskTypes}
            withSelectAsFirstOption={true}
            label={"assign task type"}
            on:change={(e) => (newTask.task_type_id = e.target.value)}
          />
          <Input
            type="datetime-local"
            label="due datetime"
            value={nowFormatted}
            on:input={(e) => (newTask.due_datetime = e.target.value)}
          />
        </div>
        <div slot="buttons">
          <Button on:click={submitAddTask}>submit</Button>
        </div>
      </Form>
    </Modal>
  {/if}
  <!-- <Button on:click={showModal}>Add Task</Button> -->
  <Button
    on:click={(e) => {
      updateTab("children");
   
    }}>{tabs.children ? "Hide" : "Show"} Children</Button
  >
  <Button
    on:click={(e) => {

      updateTab("schedule");
    }}>{tabs.schedule ? "Hide" : "Show"} Schedule</Button
  >
  <Button
    on:click={(e) => {
      updateTab("todos");
    
    }}>{tabs.todos ? "Hide" : "Show"} Todos</Button
  >
  {#if person.parent_id}
    <a href={`/#/people/${person.parent_id}`}>go back to parent</a>
  {/if}
  <h1>{person.name}s Stuff</h1>
  <div class="grid">
    {#if children && children.length > 0 && tabs.children}
      <div class="grid-member">
        <h2>{person.name}s children</h2>
        <Table headers={["id", "name", "dob", "created_at", "action"]}>
          <tbody slot="body">
            {#each children as child}
              <tr>
                <td>{child.id}</td>
                <td>{child.name}</td>
                <td>{child.dob}</td>
                <td>{child.created_at}</td>
                <td><a href={`/#/people/${child.id}`}>GO TO</a></td>
              </tr>
            {/each}
          </tbody>
        </Table>
      </div>
    {/if}
    {#if tabs.todos}
      <div class="grid-member">
        <h2>{person.name}s Todos</h2>

        <Todo person_id={person.id} />
      </div>
    {/if}
    {#if tabs.schedule}
      <div class="grid-member">
        <h2>{person.name}s schedule</h2>

        <Schedule person_id={person.id} />
      </div>
    {/if}

    {#if tasks && tasks.length > 0}
      <div class="grid-member">
        <h2>{person.name}s tasks</h2>
        <Table
          headers={[
            "id",
            "task_type",
            "due_date",
            "due_time",
            { col: 2, name: "desc" },
            "created_at",
            "action",
          ]}
        >
          <tbody slot="body">
            {#each tasks as task}
              <tr>
                <td>{task.id}</td>

                <td
                  >{task.task_type_id
                    ? taskTypes.find((t) => t.value == task.task_type_id).text
                    : ""}</td
                >
                <td>{task.due_date}</td>
                <td>{task.due_time}</td>
                <td colspan="2">{task.desc}</td>
                <td>{task.created_at}</td>
                <td><a href={`/#/tasks/${task.id}`}>GO TO</a></td>
              </tr>
            {/each}
          </tbody>
        </Table>
      </div>
    {/if}
  </div>
{/if}

<style>
  .grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-gap: 1rem;
  }
</style>
