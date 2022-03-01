<script>
  import { onMount } from "svelte";

  import config from "../config";
  import { Button, Form, Input, Select } from "../UI/UI.svelte";

  let tables = [];
  let JsonInput = "";
  let excludedRows = ["id", "created_at", "updated_at", "deleted","complete"];
  let loaded = false;
  let selectedTable = "";
  // get all tables from api function

  const getTables = async () => {
    const response = await fetch(`${config.apiUrl}/gettables`);
    const data = await response.json();

    data.tables.forEach((table) => {
      tables.push({ value: table, text: table });
    });
    loaded = true;
  };

  // get table schema after select change
  const getTableSchema = async (table) => {
    let columns = {};
    selectedTable = table;
    const response = await fetch(`${config.apiUrl}/gettableschema/${table}`);
    const data = await response.json();
    columns["table"] = table;
    for (let i = 0; i < data.columns.length; i++) {
      let column = data.columns[i];
      if (excludedRows.indexOf(column) === -1) {
        columns[column] = "";
      }
    }

    JsonInput = JSON.stringify(columns, null, 2);
  };

  const sendPost = async () => {
    const response = await fetch(config.apiUrl + "generalinput", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JsonInput,
    });
    JsonInput = "";
    // selectedTable = 0;
    await getTableSchema(selectedTable)
    return await response.json();
  };
  onMount(async () => {
    await getTables();
    console.log(tables);
  });
</script>

<h1>Table input</h1>
<div style="width: 50%; margin-left:25%">
  {#if loaded}
    <Form>
      <div slot="inputs">
        {#if tables.length >= 1}
          <Select
            options={tables}
            withSelectAsFirstOption={true}
            label={"select table"}
            on:change={(e) => {
              selectedTable = e.target.value;
              getTableSchema(e.target.value);
            }}
          />
        {/if}
        <Input
          id="desc"
          label="Description"
          bind:value={JsonInput}
          type="textarea"
        />
      </div>
      <div slot="buttons">
        <Button on:click={sendPost}>submit</Button>
      </div>
    </Form>
  {/if}
</div>
