<script>
  // generate the calendar based on month
  import CalenderItem from "./CalenderItem.svelte";
  import { onMount } from "svelte";
  export let params = {};

  // console.log(params);

  // get todays month and how many days

  const monthNames = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];

  let today = new Date();
  let month = today.getMonth();
  let year = today.getFullYear();
  let days = new Date(year, month + 1, 0).getDate();
  let firstDay = new Date(year, month, 1).getDay();
  let lastDay = new Date(year, month, days).getDay();
  let scheduleItems = [];
  let scheduleItemsArray = [];
  let loaded = false;
  let calendarDays = [];

  // get full date of first day of month YYYY-MM-DD
  let firstDayOfRange = new Date(year, month, 1).toISOString().slice(0, 10);
  let lastDayOfRange = new Date(year, month, days).toISOString().slice(0, 10);

  const getScheduleItemsByDateRange = async (firstDay, lastDay) => {
    // loaded = false;

    calendarDays = [];
    scheduleItems = [];

    const response = await fetch(
      `/api/calendar/${params.id}?between=start_date,${firstDay},${lastDay}`
    );
    const data = await response.json();
    if (data.schedule) scheduleItems = data.schedule;
    // console.lo

    scheduleItems.forEach((s, i) => {
      delete s.deleted, delete s.created_at, delete s.updated_at;
      s.day_of_month = new Date(s.start_date).getDate();
    });

    // loop over all days then add schdeule item to calendarDays
    for (let i = 1; i <= days; i++) {
      let day = {
        day_of_month: i,
        schedule_items: [],
      };

      // if(scheduleItems.length > 0) {
      //   scheduleItems.forEach((s) => {
      //     if (s.day_of_month === i) {
      //       day.schedule_items.push(s);
      //     }
      //   });
      // }

      day.schedule_items = scheduleItems.filter((s) => s.day_of_month === i);
      calendarDays = [...calendarDays, day];
      
    }

    // console.log(calendarDays);

    loaded = true;
  };

  const next = () => {
    loaded = false;
    month++;
    if (month > 11) {
      month = 0;
      year++;
    }
    days = new Date(year, month + 1, 0).getDate();
    firstDay = new Date(year, month, 1).getDay();
    firstDayOfRange = new Date(year, month, 1 + 1).toISOString().slice(0, 10);
    lastDayOfRange = new Date(year, month, days + 1).toISOString().slice(0, 10);
    getScheduleItemsByDateRange(firstDayOfRange, lastDayOfRange);
    loaded = true;
  };
  const prev = () => {
    loaded = false;

    month--;
    if (month < 0) {
      month = 11;
      year--;
    }
    days = new Date(year, month + 1, 0).getDate();
    firstDay = new Date(year, month, 1).getDay();
    firstDayOfRange = new Date(year, month, 1).toISOString().slice(0, 10);
    lastDayOfRange = new Date(year, month, days).toISOString().slice(0, 10);
    getScheduleItemsByDateRange(firstDayOfRange, lastDayOfRange);
    loaded = true;
  };

  onMount(() => {
    getScheduleItemsByDateRange(firstDayOfRange, lastDayOfRange);
  });

  // create the calendar
</script>

<button on:click={prev}>prev</button>
<button on:click={next}>next</button>
{#if loaded}
  {monthNames[month]}
  {year}
  <div class="calender">
    <!-- make boxes for day of week starting from sunday -->

    <div class="day">Sun</div>
    <div class="day">Mon</div>
    <div class="day">Tue</div>
    <div class="day">Wed</div>
    <div class="day">Thu</div>
    <div class="day">Fri</div>
    <div class="day">Sat</div>

    <!-- add empty divs for missing days -->
    {#each Array(firstDay) as _, i}
      <div class="calender-item empty" />
    {/each}
    {#if loaded && calendarDays.length > 0}
    {#each calendarDays as day}
      <CalenderItem day={day.day_of_month} scheduleItems={day.schedule_items} />
    {/each}
    {/if}
  </div>
{/if}

<style>
  .calender {
    display: grid;
    grid-template-columns: repeat(7, 1fr);
    grid-template-rows: repeat(6, 1fr);
    grid-gap: 1rem;
    margin: 1rem;
  }
</style>
