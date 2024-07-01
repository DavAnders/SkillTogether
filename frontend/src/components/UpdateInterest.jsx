import { useEffect, useState } from "react";
import api from "./Api";
import PropTypes from "prop-types";

const UpdateInterest = ({ interestId, onInterestUpdated }) => {
  const [interest, setInterest] = useState("");
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchInterest = async () => {
      try {
        const response = await api.get(`/api/interests/${interestId}`);
        setInterest(response.data.description);
      } catch (error) {
        setError("Failed to fetch interest details.");
      }
    };

    fetchInterest();
  }, [interestId]);

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      await api.put(`/api/interests/${interestId}`, {
        interest: interest,
      });
      onInterestUpdated();
    } catch (error) {
      setError("Failed to update interest.");
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mt-4">
      <div className="mb-4">
        <label
          htmlFor="interest"
          className="block text-sm font-medium text-gray-700 mb-2"
        >
          Interest:
        </label>
        <input
          id="interest"
          type="text"
          value={interest}
          onChange={(e) => setInterest(e.target.value)}
          className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        />
      </div>
      <div className="flex justify-end space-x-2">
        <button
          type="submit"
          className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          Update Interest
        </button>
        <button
          type="button"
          onClick={onInterestUpdated}
          className="px-4 py-2 bg-gray-200 text-gray-700 rounded-md hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
        >
          Cancel
        </button>
      </div>
      {error && <p className="mt-2 text-sm text-red-600">{error}</p>}
    </form>
  );
};

UpdateInterest.propTypes = {
  interestId: PropTypes.number.isRequired,
  onInterestUpdated: PropTypes.func.isRequired,
};

export default UpdateInterest;
