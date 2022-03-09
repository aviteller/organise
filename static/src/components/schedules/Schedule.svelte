<script>
  import { onMount } from "svelte";
  import ScheduleItem from "./ScheduleItem.svelte";
  import {catfre} from "../../catfre.store"
  import {
    Form,
    Input,
    Modal,
    Button,
    Select,
    Table,
  } from "../../UI/UI.svelte";

  export let person_id;

  let categories = [];
  let frequencies = [];
  let displayModal = false;

  // get todays date
  let today = new Date();
  let dd = String(today.getDate()).padStart(2, "0");
  let mm = String(today.getMonth() + 1).padStart(2, "0"); //January is 0!
  let yyyy = today.getFullYear();
  
  today = yyyy + "-" + mm + "-" + dd;

  let newScheduleItem = {
    person_id: person_id,
    schedule_cat_id: "",
    schedule_frequency_id: "",
    schedule_time: "",
    start_date: today,
    end_date: "",
    desc: "",
  };

  let scheduleItems = [];
  let loaded = false;

  const getScheduleHelpers = () => {
    return fetch(`/api/helpers?type=schedules`)
      .then((res) => res.json())
      .then((res) => {
        res.schedule_categories.forEach((cat) => {
          categories = [...categories, { value: cat.id, text: cat.title }];
        });
        res.schedule_frequencies.forEach((f) => {
          frequencies = [...frequencies, { value: f.id, text: f.title }];
        });

        catfre.set({
          categories: categories,
          frequencies: frequencies,
        });

        // console.log(categories);
        // console.log(frequencies);
      })
      .catch((err) => console.log(err));
  };
  const showModal = () => {
    if (displayModal) {
      displayModal = false;
      newScheduleItem = {
        person_id: person_id,
        schedule_cat_id: "",
        schedule_frequency_id: "",
        schedule_time: "",
        start_date: today,
        end_date: "",
        desc: "",
      };
    } else {
      displayModal = true;
    }
  };
  const editRow = (e) => {
    newScheduleItem = e.detail;
    // console.log("before", newScheduleItem);
    if ("deleted" in newScheduleItem) delete newScheduleItem.deleted;
    if ("created_at" in newScheduleItem) delete newScheduleItem.created_at;
    if ("updated_at" in newScheduleItem) delete newScheduleItem.updated_at;
    // console.log("after", newScheduleItem);
    // console.log(newScheduleItem);
    showModal();
  };

  const deleteRow = (e) => {
    let id = e.detail;

    fetch(`/api/schedule/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        scheduleItems = scheduleItems.filter((t) => t.id !== id);
      });
  };

  const getScheduleItems = async () => {
    loaded = false;
    const response = await fetch(`/api/schedules/${person_id}`);
    const data = await response.json();
    if (data.schedule_items) scheduleItems = data.schedule_items;
    // scheduleItems.forEach((i) => {
    //   delete i.deleted, delete i.created_at, delete i.updated_at;
    // });
    loaded = true;
    // console.log(scheduleItems.length, loaded);
    // return data;
  };

  const submitEditScheduleItem = () => {
    fetch(`/api/schedules/${newScheduleItem.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newScheduleItem),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        scheduleItems = scheduleItems.map((t) => {
          if (t.id === newScheduleItem.id) {
            return newScheduleItem;
          } else {
            return t;
          }
        });
        showModal();
        // sort by completed
      });
  };

  const submitAddScheduleItem = () => {
    // console.log(newTodo);
    fetch("/api/schedules", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newScheduleItem),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data.row);
        scheduleItems = [...scheduleItems, data.row];
        newScheduleItem = {
          person_id: person_id,
          schedule_cat_id: "",
          schedule_frequency_id: "",
          schedule_time: "",
          start_date: today,
          end_date: "",
          desc: "",
        };

        // getTodos();
        showModal();
      });
  };

  onMount(async () => {
    getScheduleHelpers().then(() => getScheduleItems());
    //  getScheduleItems();
  });
</script>

<Button on:click={showModal}>Add Item</Button>

{#if displayModal}
  <Modal
    title={`${newScheduleItem.id ? "edit" : "add"} item`}
    on:cancel={showModal}
  >
    <!--   schedule_cat_id: "",
          schedule_frequency_id: "",
          schedule_time: "",
          start_date: "",
          end_date: "",
          desc: "", -->
    <Form>
      <div slot="inputs">
        <Input
          id="desc"
          label="Content"
          bind:value={newScheduleItem.desc}
          on:input={(e) => (newScheduleItem.desc = e.target.value)}
        />
        <Input
          id="desc"
          label="start time"
          bind:value={newScheduleItem.schedule_time}
          on:input={(e) => (newScheduleItem.schedule_time = e.target.value)}
          type="time"
        />
        <Input
          id="desc"
          label="start date"
          bind:value={newScheduleItem.start_date}
          on:input={(e) => (newScheduleItem.start_date = e.target.value)}
          type="date"
        />
        <Input
          id="desc"
          label="end date"
          bind:value={newScheduleItem.end_date}
          on:input={(e) => (newScheduleItem.end_date = e.target.value)}
          type="date"
        />

        <Select
          options={categories}
          withSelectAsFirstOption={true}
          selectedOption={newScheduleItem.schedule_cat_id}
          label={"assign cat"}
          on:change={(e) => (newScheduleItem.schedule_cat_id = e.target.value)}
        />
        <Select
          options={frequencies}
          withSelectAsFirstOption={true}
          selectedOption={newScheduleItem.schedule_frequency_id}
          label={"assign frequency"}
          on:change={(e) =>
            (newScheduleItem.schedule_frequency_id = e.target.value)}
        />
      </div>
      <div slot="buttons">
        {#if newScheduleItem.id}
          <Button on:click={submitEditScheduleItem}>Save</Button>
        {:else}
          <Button on:click={submitAddScheduleItem}>Add</Button>
        {/if}
      </div>
    </Form>
  </Modal>
{/if}

{#if loaded && scheduleItems.length > 0}
  <Table
    headers={[
      "desc",
      "time",
      "startdate",
      "enddate",
      "category",
      "frequency",
      { name: "actions", col: 2 },
    ]}
  >
    <tbody slot="body">
      {#each scheduleItems as item}
        <tr>
          <ScheduleItem
            {item}
            catAndFreq={{ categories, frequencies }}
            on:deleterow={deleteRow}
            on:editrow={editRow}
            control={true}
          />
        </tr>
      {/each}
    </tbody>
  </Table>
{:else}
  <p>No schedule items yet</p>
{/if}
