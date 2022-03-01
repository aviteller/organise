<script>
  import { onMount } from "svelte";

  import config from "../config";
  let loaded = false;
  let excludedRows = ["created_at", "updated_at"];
  import { Button, Form, Input, Select, Table } from "../UI/UI.svelte";
  let tables = [];
  let selectColumns = [];
  let selectedTable = "";
  let columnValue = "";
  let selectedColumn = "";
  let tableRows = [];
  let tableColumns = [];
  let tableLoaded = false;

  const getTables = async () => {
    const response = await fetch(`${config.apiUrl}/gettables`);
    const data = await response.json();
    // console.log(data);
    data.tables.forEach((table) => {
      tables.push({ value: table, text: table });
    });
    loaded = true;
  };

  const getTableSchema = async (table) => {
    loaded = false;
    selectColumns = [];
    let columns = {};
    selectedTable = table;
    const response = await fetch(`${config.apiUrl}/gettableschema/${table}`);
    const data = await response.json();
    columns["table"] = table;
    for (let i = 0; i < data.columns.length; i++) {
      let column = data.columns[i];
      if (excludedRows.indexOf(column) === -1) {
        selectColumns.push({ value: column, text: column });
      }
    }
    loaded = true;
  };

  // send post query function
  const sendQuery = async () => {
    tableLoaded = false;
    tableColumns = [];
    tableRows = [];

    let query = `SELECT * FROM ${selectedTable}`;
    console.log(columnValue, selectedColumn);
    if (columnValue !== "") {
      query += ` WHERE ${selectedColumn} = '${columnValue}'`;
    }

    let response = await fetch(`${config.apiUrl}generalquery`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ query: query }),
    });
    const data = await response.json();
    tableRows = data.rows;
    if (!tableRows  && tableRows.length  <0) {
      tableRows = [{}];
    } else {
      tableRows.forEach((row) => {
        tableColumns.push(Object.keys(row));
      });
    }

    selectedTable = "";
    columnValue = "";
    selectedColumn = "";
    tableLoaded = true;
  };

  onMount(async () => {
    await getTables();
  });
</script>

<div style="width: 40%; margin-left:30%">
  <Form>
    <div slot="inputs">
      {#if tables.length >= 1 && loaded}
        <Select
          options={tables}
          withSelectAsFirstOption={true}
          label={"select table"}
          selectedOption={selectedTable}
          on:change={(e) => {
            getTableSchema(e.target.value);
          }}
        />
      {/if}
      {#if selectColumns.length >= 1}
        <Select
          options={selectColumns}
          withSelectAsFirstOption={true}
          label={"select column"}
          on:change={(e) => {
            selectedColumn = e.target.value;
          }}
        />
        <Input
          id="columnValue"
          label="column value"
          value={columnValue}
          on:input={(e) => (columnValue = e.target.value)}
          type="text"
        />
      {/if}
    </div>
    <div slot="buttons">
      {#if selectedTable}
        <Button on:click={sendQuery}>Send Query</Button>
      {/if}
    </div>
  </Form>
</div>
{#if tableLoaded}
  {#if tableRows && tableRows.length > 0}
    <Table headers={tableColumns[0]}>
      <tbody slot="body">
        {#each tableRows as row}
          <tr>
            {#each Object.entries(row) as column}
              <td>{column[1]}</td>
            {/each}
          </tr>
        {/each}
      </tbody>
    </Table>
  {:else}
    <p>No results found</p>
  {/if}
{/if}
