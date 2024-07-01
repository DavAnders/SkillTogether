import { useState } from "react";
import InterestItem from "./InterestItem";
import api from "./api";
import PropTypes from "prop-types";

const InterestsList = ({ interests, refreshInterests }) => {
  const [error, setError] = useState(null);

  const handleDelete = async (id) => {
    try {
      setError(null); // Clear any previous errors
      await api.delete(`/api/interests/${id}`);
      refreshInterests();
    } catch (error) {
      console.error("Failed to delete interest:", error);
      setError("Failed to delete interest. Please try again.");
    }
  };

  const handleInterestUpdated = async () => {
    try {
      setError(null); // Clear any previous errors
      await refreshInterests();
    } catch (error) {
      console.error("Failed to refresh interests:", error);
      setError("Failed to update interests list. Please refresh the page.");
    }
  };

  return (
    <div className="interests-list">
      {error && <p className="text-red-500 mb-4">{error}</p>}
      {Array.isArray(interests) && interests.length > 0 ? (
        <ul className="space-y-2">
          {interests.map((interest) => (
            <InterestItem
              key={interest.id}
              interest={interest}
              handleDelete={handleDelete}
              onInterestUpdated={handleInterestUpdated}
            />
          ))}
        </ul>
      ) : (
        <p className="text-gray-500 italic">No interests found.</p>
      )}
    </div>
  );
};

InterestsList.propTypes = {
  interests: PropTypes.array,
  refreshInterests: PropTypes.func.isRequired,
};

export default InterestsList;
