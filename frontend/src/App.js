import React from "react";
import JobForm from "./components/JobForm";
import JobList from "./components/JobList";

function App() {
  return (
    <div className="app-container">
      <JobForm />
      <JobList />
    </div>
  );
}

export default App;
