import { useState, useEffect } from "react";
import PropTypes from "prop-types";
import DOMPurify from "dompurify";
import api from "./Api";

const UpdateInterest = ({ interestId, onInterestUpdated }) => {
  const [interest, setInterest] = useState("");
  const [error, setError] = useState("");

  useEffect(() => {
    // Fetch the current interest
    const fetchInterest = async () => {
      // Only allow numbers for interest ID
      if (!/^\d+$/.test(interestId)) {
        throw new Error("Invalid interest ID");
      }

      try {
        const response = await api.get(`/api/interests/${interestId}`);
        setInterest(response.data.interest);
      } catch (error) {
        setError("Failed to fetch interest details.");
      }
    };

    fetchInterest();
  }, [interestId]);

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (!/^\d+$/.test(interestId)) {
      throw new Error("Invalid interest ID");
    }

    // Sanitize the input
    const sanitizedInterest = DOMPurify.sanitize(interest);

    try {
      await api.put(`/api/interests/${interestId}`, {
        interest: sanitizedInterest,
      });
      onInterestUpdated();
    } catch (error) {
      setError("Failed to update interest.");
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <label htmlFor="interest">Interest:</label>
      <input
        id="interest"
        type="text"
        value={interest}
        onChange={(e) => setInterest(e.target.value)}
      />
      <button type="interest">Update Interest</button>
      {error && <p>{error}</p>}
    </form>
  );
};

UpdateInterest.propTypes = {
  interestId: PropTypes.number.isRequired,
  onInterestUpdated: PropTypes.func.isRequired,
};

export default UpdateInterest;
