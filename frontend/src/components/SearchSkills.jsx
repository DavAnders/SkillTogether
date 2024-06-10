import { useState } from "react";
import "../styles/SearchSkills.css";
import api from "./Api";

const SearchSkills = () => {
  const [query, setQuery] = useState("");
  const [skills, setSkills] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");

  const handleSearch = async () => {
    if (!query) {
      setError("Please enter a search term.");
      return;
    }
    setIsLoading(true);
    setError("");
    try {
      const response = await api.get("/api/search/skills", {
        params: { q: query },
      });
      setSkills(response.data);
      setIsLoading(false);
    } catch (error) {
      setError(
        "Failed to fetch skills: " +
          (error.response?.data?.error || error.message)
      );
      setIsLoading(false);
    }
  };

  const handleInterested = (discordId) => {
    window.location.href = `https://discordapp.com/users/${discordId}`;
  };

  return (
    <div className="search-skills-container">
      <div className="search-skills">
        <h2>Search for Skills</h2>
        <div className="search-input">
          <input
            type="text"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            placeholder="Enter skill to search..."
          />
          <button onClick={handleSearch} disabled={isLoading}>
            {isLoading ? "Searching..." : "Search"}
          </button>
        </div>
        {error && <div style={{ color: "red" }}>{error}</div>}
        <div className="search-results-container">
          {skills.length > 0 ? (
            <ul>
              {skills.map((result, index) => (
                <li key={index}>
                  <div>
                    <strong>Skill:</strong> {result.skill.skill_description}
                  </div>
                  <div>
                    <strong>Posted by:</strong> {result.user.username}
                  </div>
                  <button
                    onClick={() => handleInterested(result.user.discord_id)}
                  >
                    Interested
                  </button>
                </li>
              ))}
            </ul>
          ) : (
            !isLoading && <p>No skills found, please try a different search.</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default SearchSkills;
