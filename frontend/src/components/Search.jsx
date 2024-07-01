import { useState } from "react";
import { useNavigate } from "react-router-dom";

const Search = () => {
  const [query, setQuery] = useState("");
  const [searchType, setSearchType] = useState("skills");
  const navigate = useNavigate();

  const handleSearch = (e) => {
    e.preventDefault();
    if (query.trim()) {
      navigate(`/search/${searchType}?q=${encodeURIComponent(query)}`);
    }
  };

  return (
    <div className="bg-white shadow rounded-lg p-4">
      <h2 className="text-xl font-semibold mb-4">Search for:</h2>
      <form onSubmit={handleSearch} className="space-y-4">
        <div>
          <input
            type="text"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            placeholder="Enter search term..."
            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div className="flex space-x-4">
          <label className="inline-flex items-center">
            <input
              type="radio"
              value="skills"
              checked={searchType === "skills"}
              onChange={() => setSearchType("skills")}
              className="form-radio h-5 w-5 text-violet-400 border-gray-300 focus:ring-violet-500"
            />
            <span className="ml-2">Skills</span>
          </label>
          <label className="inline-flex items-center">
            <input
              type="radio"
              value="interests"
              checked={searchType === "interests"}
              onChange={() => setSearchType("interests")}
              className="form-radio h-5 w-5 text-violet-400 border-gray-300 focus:ring-violet-500"
            />
            <span className="ml-2">Interests</span>
          </label>
        </div>
        <button
          type="submit"
          className="w-full bg-violet-400 text-white py-2 px-4 rounded-md hover:bg-violet-500 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:ring-offset-2"
        >
          Search
        </button>
      </form>
    </div>
  );
};

export default Search;
