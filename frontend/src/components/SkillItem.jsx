import { useState } from "react";
import PropTypes from "prop-types";
import UpdateSkill from "./UpdateSkill";

const SkillItem = ({ skill, handleDelete, onSkillUpdated }) => {
  const [isEditing, setIsEditing] = useState(false);

  if (!skill || typeof skill !== "object") {
    console.error("Interest prop is missing or not an object", skill);
    return null; // Render nothing if interest prop is invalid
  }

  return (
    <li className="bg-white shadow-sm rounded-lg p-3 mb-2 transition-all duration-300 hover:shadow-md">
      <div className="flex items-center justify-between">
        <span className="text-gray-800">{skill.skill_description}</span>
        {!isEditing && (
          <div className="flex space-x-2">
            <button
              onClick={() => handleDelete(skill.id)}
              className="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded text-sm transition-colors duration-300"
            >
              Delete
            </button>
            <button
              onClick={() => setIsEditing(true)}
              className="bg-violet-400 hover:bg-violet-500 text-white px-3 py-1 rounded text-sm transition-colors duration-300"
            >
              Edit
            </button>
          </div>
        )}
      </div>
      {isEditing && (
        <div className="mt-2">
          <UpdateSkill
            skillId={skill.id}
            onSkillUpdated={() => {
              setIsEditing(false);
              onSkillUpdated();
            }}
          />
        </div>
      )}
    </li>
  );
};

SkillItem.propTypes = {
  skill: PropTypes.shape({
    id: PropTypes.number,
    skill_description: PropTypes.string,
  }).isRequired,
  handleDelete: PropTypes.func.isRequired,
  onSkillUpdated: PropTypes.func.isRequired,
};

export default SkillItem;
