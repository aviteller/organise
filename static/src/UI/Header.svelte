<script>
  import Cookies from "../utils/Cookies";
  // import { cookieUser } from "../utils/helpers";
  import { push, pop } from "svelte-spa-router";
  import Button from "./Button.svelte";
  // console.log("cookieUser :>> ", cookieUser);
  let c = new Cookies();

  const logout = async () => {
    await c.eraseCookie("user-jwt");
    push("/#/login");
    window.location.reload();
  };
</script>

<style>
  header {
    /* position: fixed; */
    width: 100%;
    top: 0;
    left: 0;
    height: 4rem;
    background-color: white;
    border-bottom: 12px solid #14a4e3;
    display: flex;
    justify-content: space-between;
    align-content: center;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.26);
    margin-bottom: 20px;
    z-index: 999;
  }
</style>

<header>
  {#if cookieUser.name}
    <h1>Hello {cookieUser.name}</h1>
    <Button color="danger" on:click={() => pop()}>Go Back</Button>
    <Button color="success" on:click={logout}>Logout</Button>
    {#if cookieUser.name === 'admin'}
      <Button mode="outline" href="/#/admin/users">Go to users</Button>
      <Button mode="outline" href="/#/admin/companies">Go to companies</Button>
      <Button mode="outline" href="/#/admin/chatrooms">Go to chatrooms</Button>
    {/if}
  {:else}
    <Button mode="outline" href="#/Register">Register</Button>
    <Button mode="outline" href="#/Login">Login</Button>
  {/if}
</header>
