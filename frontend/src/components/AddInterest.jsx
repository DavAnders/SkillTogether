import { useState } from "react";
import api from "./api";
import PropTypes from "prop-types";

const AddInterest = ({ refreshInterests }) => {
  const [interest, setInterest] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [isDialogOpen, setIsDialogOpen] = useState(false);

  const sanitizeInput = (input) => {
    const sanitizedInput = input.replace(/[<>]/g, "");
    return sanitizedInput;
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    setIsLoading(true);

    try {
      const sanitizedInterest = sanitizeInput(interest);
      const response = await api.post("/api/interests", {
        interest: sanitizedInterest,
      });
      if (response.status === 200) {
        setInterest("");
        refreshInterests();
        setIsDialogOpen(false);
      }
    } catch (error) {
      console.error("Failed to add interest:", error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <>
      <button
        onClick={() => setIsDialogOpen(true)}
        className="bg-violet-400 text-white rounded-full w-8 h-8 flex items-center justify-center shadow-md hover:bg-violet-500 transition-colors"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          className="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth={2}
            d="M12 6v12m6-6H6"
          />
        </svg>
      </button>
      {isDialogOpen && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white p-4 rounded-lg">
            <h2 className="text-lg font-bold mb-4">Add New Interest</h2>
            <form onSubmit={handleSubmit} className="space-y-4">
              <input
                type="text"
                value={interest}
                onChange={(e) => setInterest(e.target.value)}
                placeholder="Enter a new interest"
                className="w-full p-2 border rounded"
              />
              <div className="flex justify-end space-x-2">
                <button
                  type="button"
                  onClick={() => setIsDialogOpen(false)}
                  className="px-4 py-2 bg-gray-200 rounded"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  disabled={isLoading}
                  className="px-4 py-2 bg-violet-400 text-white rounded hover:bg-violet-500"
                >
                  {isLoading ? "Adding..." : "Add Interest"}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </>
  );
};

AddInterest.propTypes = {
  refreshInterests: PropTypes.func.isRequired,
};

export default AddInterest;
