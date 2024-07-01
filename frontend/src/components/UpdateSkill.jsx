import { useEffect, useState } from "react";
import api from "./api";
import PropTypes from "prop-types";

const UpdateSkill = ({ skillId, onSkillUpdated }) => {
  const [skill, setSkill] = useState("");
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchSkill = async () => {
      try {
        const response = await api.get(`/api/skills/${skillId}`);
        setSkill(response.data.description);
      } catch (error) {
        setError("Failed to fetch skill details.");
      }
    };

    fetchSkill();
  }, [skillId]);

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      await api.put(`/api/skills/${skillId}`, {
        skill_description: skill,
      });
      onSkillUpdated();
    } catch (error) {
      setError("Failed to update skill.");
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mt-4">
      <div className="mb-4">
        <label
          htmlFor="skill"
          className="block text-sm font-medium text-gray-700 mb-2"
        >
          Skill:
        </label>
        <input
          id="skill"
          type="text"
          value={skill}
          onChange={(e) => setSkill(e.target.value)}
          className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        />
      </div>
      <div className="flex justify-end space-x-2">
        <button
          type="submit"
          className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          Update Skill
        </button>
        <button
          type="button"
          onClick={onSkillUpdated}
          className="px-4 py-2 bg-gray-200 text-gray-700 rounded-md hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
        >
          Cancel
        </button>
      </div>
      {error && <p className="mt-2 text-sm text-red-600">{error}</p>}
    </form>
  );
};

UpdateSkill.propTypes = {
  skillId: PropTypes.number.isRequired,
  onSkillUpdated: PropTypes.func.isRequired,
};

export default UpdateSkill;
