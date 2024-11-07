import React, { useState, useEffect } from "react";
import "./JobList.css"; // Make sure the styles are applied

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
    <div className="job-list-container">
      <h2>Job List</h2>
      <div className="job-cards">
        {jobs?.map((job) => (
          <div className="job-card" key={job.id}>
            <div className="job-card-header">
              <h3>{job.name}</h3>
              <span className={`status ${job.status.toLowerCase()}`}>
                {job.status}
              </span>
            </div>
            <div className="job-card-body">
              <div className="left">
                <p>Duration: {job.duration / 1000000000}s</p>
              </div>
              <div className="right">
                <p>Remaining: {job.remaining_time / 1000000000}s</p>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default JobList;
