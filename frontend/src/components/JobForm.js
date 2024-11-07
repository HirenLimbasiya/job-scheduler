import React, { useState } from "react";
import "./JobForm.css"; // Import the updated CSS

const JobForm = () => {
  const [name, setName] = useState("");
  const [duration, setDuration] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    await fetch("http://localhost:2027/jobs", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name, duration: parseInt(duration) }),
    });

    setName("");
    setDuration("");
  };

  return (
    <div className="job-form-container">
      <div className="job-form-card">
        <h2>Create a New Job</h2>
        <form onSubmit={handleSubmit}>
          <div className="input-group">
            <div>
              <label htmlFor="job-name">Job Name</label>
              <input
                type="text"
                id="job-name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="Enter job name"
              />
            </div>
            <div>
              <label htmlFor="job-duration">Duration (seconds)</label>
              <input
                type="number"
                id="job-duration"
                value={duration}
                onChange={(e) => setDuration(e.target.value)}
                placeholder="Enter duration"
              />
            </div>
          </div>
          <button type="submit">Submit Job</button>
        </form>
      </div>
    </div>
  );
};

export default JobForm;
