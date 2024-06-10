import { useState, useEffect } from "react";
import PropTypes from "prop-types";
import DOMPurify from "dompurify";
import api from "./Api";

const UpdateSkill = ({ skillId, onSkillUpdated }) => {
  const [skill, setSkill] = useState("");
  const [error, setError] = useState("");

  useEffect(() => {
    // Fetch the current skill description
    const fetchSkill = async () => {
      if (!/^\d+$/.test(skillId)) {
        throw new Error("Invalid skill ID");
      }

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

    if (!/^\d+$/.test(skillId)) {
      throw new Error("Invalid skill ID");
    }

    // Sanitize the input
    const sanitizedSkill = DOMPurify.sanitize(skill);

    try {
      await api.put(`/api/skills/${skillId}`, {
        skill_description: sanitizedSkill,
      });
      onSkillUpdated();
    } catch (error) {
      setError("Failed to update skill.");
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <label htmlFor="skill">Skill:</label>
      <input
        id="skill"
        type="text"
        value={skill}
        onChange={(e) => setSkill(e.target.value)}
      />
      <button type="submit">Update Skill</button>
      {error && <p>{error}</p>}
    </form>
  );
};

UpdateSkill.propTypes = {
  skillId: PropTypes.number.isRequired,
  onSkillUpdated: PropTypes.func.isRequired,
};

export default UpdateSkill;
