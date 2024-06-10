import { useState } from "react";
import DOMPurify from "dompurify";
import "../styles/AddSkill.css";
import api from "./Api";

const AddSkill = () => {
  const [skill, setSkill] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();

    // Sanitize the input
    const sanitizedSkill = DOMPurify.sanitize(skill);

    try {
      const response = await api.post("/api/skills", {
        description: sanitizedSkill,
      });
      if (response.status === 200) {
        setSkill(""); // Reset skill after success
        setError(""); // Reset error
      } else {
        setError("Failed to add skill. Please try again.");
      }
    } catch (error) {
      if (error.response) {
        setError(`Error: ${error.response.data.error}`);
      } else if (error.request) {
        // The request was made but no response was received
        setError("No response from server. Check your connection.");
      } else {
        setError("Error: " + error.message);
      }
    }
  };

  return (
    <form className="add-skill-form" onSubmit={handleSubmit}>
      <label htmlFor="skill">Skill:</label>
      <input
        id="skill"
        type="text"
        value={skill}
        onChange={(e) => setSkill(e.target.value)}
      />
      <button type="submit">Add Skill</button>
      {error && <p>{error}</p>}
    </form>
  );
};

export default AddSkill;
