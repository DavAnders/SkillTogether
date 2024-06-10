import { useState } from "react";
import DOMPurify from "dompurify";
import "../styles/AddSkill.css";
import api from "./Api";

const AddInterest = () => {
  const [interest, setInterest] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();

    // Sanitize the input
    const sanitizedInterest = DOMPurify.sanitize(interest);

    try {
      const response = await api.post("/api/interests", {
        interest: sanitizedInterest,
      });
      if (response.status === 200) {
        setInterest(""); // Reset interest after success
        setError(""); // Reset error
      } else {
        setError("Failed to add interest. Please try again.");
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
    <form className="add-interest-form" onSubmit={handleSubmit}>
      <label htmlFor="interest">Interest:</label>
      <input
        id="interest"
        type="text"
        value={interest}
        onChange={(e) => setInterest(e.target.value)}
      />
      <button type="submit">Add Interest</button>
      {error && <p>{error}</p>}
    </form>
  );
};

export default AddInterest;
