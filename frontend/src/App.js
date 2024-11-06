import React from "react";
import JobForm from "./JobForm";
import JobList from "./JobList";

function App() {
  return (
    <div>
      <h1>Job Scheduler</h1>
      <JobForm />
      <JobList />
    </div>
  );
}

export default App;
