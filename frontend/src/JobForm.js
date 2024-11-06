import React, { useState } from "react";

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
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="Job Name"
      />
      <input
        type="number"
        value={duration}
        onChange={(e) => setDuration(e.target.value)}
        placeholder="Duration (seconds)"
      />
      <button type="submit">Submit Job</button>
    </form>
  );
};

export default JobForm;
