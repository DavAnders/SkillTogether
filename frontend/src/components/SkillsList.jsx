import { useState } from "react";
import api from "./api";
import SkillItem from "./SkillItem";
import PropTypes from "prop-types";

const SkillsList = ({ skills, refreshSkills }) => {
  const [error, setError] = useState(null);

  const handleDelete = async (id) => {
    try {
      setError(null); // Clear any previous errors
      await api.delete(`/api/skills/${id}`);
      refreshSkills();
    } catch (error) {
      console.error("Failed to delete skill:", error);
      setError("Failed to delete skill. Please try again.");
    }
  };

  const handleSkillUpdated = async () => {
    try {
      setError(null); // Clear any previous errors
      await refreshSkills();
    } catch (error) {
      console.error("Failed to refresh skills:", error);
      setError("Failed to update skills list. Please refresh the page.");
    }
  };

  return (
    <div className="skills-list">
      {error && <p className="text-red-500 mb-4">{error}</p>}
      {Array.isArray(skills) && skills.length > 0 ? (
        <ul className="space-y-2">
          {skills.map((skill) => (
            <SkillItem
              key={skill.id}
              skill={skill}
              handleDelete={handleDelete}
              onSkillUpdated={handleSkillUpdated}
            />
          ))}
        </ul>
      ) : (
        <p className="text-gray-500 italic">No skills found.</p>
      )}
    </div>
  );
};

SkillsList.propTypes = {
  skills: PropTypes.array,
  refreshSkills: PropTypes.func.isRequired,
};

export default SkillsList;
