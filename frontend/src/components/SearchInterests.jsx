import { useState } from "react";
import "../styles/SearchSkills.css";
import api from "./Api";

const SearchInterests = () => {
  const [query, setQuery] = useState("");
  const [interests, setInterests] = useState([]);
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
      const response = await api.get(`/api/search/interests`, {
        params: { q: query },
      });
      setInterests(response.data);
      setIsLoading(false);
    } catch (error) {
      setError(
        "Failed to fetch interests: " +
          (error.response?.data?.error || error.message)
      );
      setIsLoading(false);
    }
  };

  const handleContact = (discordId) => {
    window.location.href = `https://discordapp.com/users/${discordId}`;
  };

  return (
    <div className="search-interests-container">
      <div className="search-interests">
        <h2>Search for Interests</h2>
        <div className="search-input">
          <input
            type="text"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            placeholder="Enter interest to search..."
          />
          <button onClick={handleSearch} disabled={isLoading}>
            {isLoading ? "Searching..." : "Search"}
          </button>
        </div>
        {error && <div style={{ color: "red" }}>{error}</div>}
        <div className="search-results-container">
          {interests.length > 0 ? (
            <ul>
              {interests.map((result, index) => (
                <li key={index}>
                  <div>
                    <strong>Interest:</strong> {result.interest.interest}
                  </div>
                  <div>
                    <strong>Posted by:</strong> {result.user.username}
                  </div>
                  <button onClick={() => handleContact(result.user.discord_id)}>
                    Contact
                  </button>
                </li>
              ))}
            </ul>
          ) : (
            !isLoading && (
              <p>No interests found, please try a different search.</p>
            )
          )}
        </div>
      </div>
    </div>
  );
};

export default SearchInterests;
