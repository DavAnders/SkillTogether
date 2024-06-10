import PropTypes from "prop-types";
import { useState } from "react";
import UpdateSkill from "./UpdateSkill";
import "../styles/SkillItem.css";

function SkillItem({ skill, handleDelete, onSkillUpdated }) {
  const [isEditing, setIsEditing] = useState(false);

  return (
    <div className="skill-item">
      <li key={skill.id}>
        {skill.skill_description}
        {!isEditing && (
          <div
            style={{ display: "flex", gap: "10px" }}
            className="skill-buttons"
          >
            <button onClick={() => handleDelete(skill.id)}>Delete</button>
            <button onClick={() => setIsEditing(true)}>Edit</button>
          </div>
        )}
        {isEditing && (
          <UpdateSkill
            skillId={skill.id}
            onSkillUpdated={() => {
              setIsEditing(false);
              onSkillUpdated();
            }}
          />
        )}
      </li>
    </div>
  );
}

SkillItem.propTypes = {
  skill: PropTypes.shape({
    id: PropTypes.number,
    skill_description: PropTypes.string,
  }).isRequired,
  handleDelete: PropTypes.func.isRequired,
  onSkillUpdated: PropTypes.func.isRequired,
};

export default SkillItem;
