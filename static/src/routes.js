import Home from './routes/Home.svelte'
import Person from './routes/Person.svelte'
import Task from './routes/Task.svelte'
import GeneralInput from "./routes/GeneralInput.svelte";
import GeneralGet from "./routes/GeneralGet.svelte";
import Calender from "./routes/Calender.svelte";
// import Project from './routes/Project.svelte'
// import Task from './routes/Task.svelte'
import NotFound from './routes/NotFound.svelte'

export const routes = {
    // Exact path
    '/': Home,
    "/input": GeneralInput,
    "/get": GeneralGet,
    "/calender": Calender,
    '/people/:id': Person,
    '/tasks/:id': Task,
    // '/tasks/:id': Task,

    // // Using named parameters, with last being optional
    // '/author/:first/:last?': Author,

    // // Wildcard parameter
    // '/book/*': Book,

    // Catch-all
    // This is optional, but if present it must be the last
    '*': NotFound,
}