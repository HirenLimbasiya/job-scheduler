import React, { useState, useEffect } from "react";

const JobList = () => {
  const [jobs, setJobs] = useState([]);

  useEffect(() => {
    const fetchJobs = async () => {
      const response = await fetch("http://localhost:2027/jobs");
      const data = await response.json();
      console.log("data", data);
      setJobs(data || []);
    };

    fetchJobs();

    const ws = new WebSocket("ws://localhost:2027/ws");
    ws.onmessage = (event) => {
      setJobs(JSON.parse(event.data));
    };

    return () => ws.close();
  }, []);

  console.log("jobs", jobs);

  return (
    <div>
      <h2>Job List</h2>
      <ul>
        {jobs?.map((job) => (
          <li key={job.id}>
            {job.name} - {job.status} - {job.duration}s - {job.remaining_time}s
          </li>
        ))}
      </ul>
    </div>
  );
};

export default JobList;
