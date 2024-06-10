import { useEffect, useState } from "react";
import "../styles/SkillsList.css";
import SkillItem from "./SkillItem";
import api from "./Api";

const SkillsList = () => {
  const [skills, setSkills] = useState([]);
  const [error, setError] = useState("");

  useEffect(() => {
    fetchSkills();
  }, []);

  const fetchSkills = async () => {
    try {
      const response = await api.get("/api/skills");
      // Check if response.data is an array, if not, set to empty array
      setSkills(Array.isArray(response.data) ? response.data : []);
      setError("");
    } catch (err) {
      setError("Failed to fetch skills: " + err.message);
      setSkills([]);
    }
  };

  const handleSkillUpdated = () => {
    fetchSkills(); // Refresh the list after update
    setError(""); // Reset error
  };

  const handleDelete = async (id) => {
    try {
      await api.delete(`/api/skills/${id}`);
      fetchSkills(); // Refresh the list after deletion
      setError(""); // Reset error
    } catch (err) {
      setError("Failed to delete the skill: " + err.message);
    }
  };

  return (
    <div className="skills-list">
      <h2>Your Skills</h2>
      {error && <p>{error}</p>}
      <ul>
        {skills.length > 0 ? (
          skills.map((skill) => (
            <SkillItem
              key={skill.id}
              skill={skill}
              handleDelete={handleDelete}
              onSkillUpdated={handleSkillUpdated}
            />
          ))
        ) : (
          <p>No skills found.</p>
        )}
      </ul>
    </div>
  );
};

export default SkillsList;
