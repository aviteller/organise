<script>
  import { link,push } from "svelte-spa-router";
  import { onMount } from "svelte";
  import { Table } from "../UI/UI.svelte";
  import config from "../config";
  let loaded = false;
  let people = [];

  const getPeople = () => {
    fetch(`${config.apiUrl}people`)
      .then((res) => res.json())
      .then((res) => {
        people = res.people;
        loaded = true;
      });
  };

  onMount(() => {

    // this is just for me to use the app alone
    push('/people/1')
    // getPeople();
  });
</script>

{#if loaded}
  {#if people && people.length > 0}
    <Table
      headers={[
        "id",
        "parent_id",
        "person_type",
        "name",
        "dob",
        "created_at",
        "action",
      ]}
    >
      <tbody slot="body">
        {#each people as person}
          <tr>
            <td>{person.id}</td>
            <td>{person.parent_id}</td>
            <td>{person.parent_id != 0 ? "child" : "parent"}</td>
            <td>{person.name}</td>
            <td>{person.dob}</td>
            <td>{person.created_at}</td>
            <td><a href={`#/people/${person.id}`}>GO TO</a></td>
          </tr>
        {/each}
      </tbody>
    </Table>
  {/if}
{/if}
