import { useEffect, useState } from "react";
import api from "./Api";
import "../styles/SkillsList.css";
import InterestItem from "./InterestItem";

const InterestsList = () => {
  const [interests, setInterests] = useState([]);
  const [error, setError] = useState("");

  useEffect(() => {
    fetchInterests();
  }, []);

  const fetchInterests = async () => {
    try {
      const response = await api.get("/api/interests");
      // Check if response.data is an array, if not, set to empty array
      setInterests(Array.isArray(response.data) ? response.data : []);
      setError("");
    } catch (err) {
      setError("Failed to fetch interests: " + err.message);
      setInterests([]);
    }
  };

  const handleInterestUpdated = () => {
    fetchInterests(); // Refresh the list after update
    setError(""); // Reset error
  };

  const handleDelete = async (id) => {
    try {
      await api.delete(`/api/interests/${id}`);
      fetchInterests(); // Refresh the list after deletion
      setError(""); // Reset error
    } catch (err) {
      setError("Failed to delete the interest: " + err.message);
    }
  };

  return (
    <div className="interests-list">
      <h2>Your Interests</h2>
      {error && <p>{error}</p>}
      <ul>
        {interests.length > 0 ? (
          interests.map((interest) => (
            <InterestItem
              key={interest.id}
              interest={interest}
              handleDelete={handleDelete}
              onInterestUpdated={handleInterestUpdated}
            />
          ))
        ) : (
          <p>No interests found.</p>
        )}
      </ul>
    </div>
  );
};

export default InterestsList;
