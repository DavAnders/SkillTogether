import { useState, useEffect } from "react";
import { useParams, useSearchParams, Link } from "react-router-dom";
import api from "./Api";

const SearchResults = () => {
  const [results, setResults] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState("");
  const { type } = useParams();
  const [searchParams] = useSearchParams();
  const query = searchParams.get("q");

  useEffect(() => {
    const fetchResults = async () => {
      setIsLoading(true);
      setError("");
      try {
        const response = await api.get(`/api/search/${type}`, {
          params: { q: query },
        });
        setResults(response.data);
      } catch (error) {
        setError(
          `Failed to fetch ${type}: ${
            error.response?.data?.error || error.message
          }`
        );
      } finally {
        setIsLoading(false);
      }
    };

    if (query) {
      fetchResults();
    }
  }, [type, query]);

  const handleInterested = (discordId) => {
    window.open(`https://discordapp.com/users/${discordId}`, "_blank");
  };

  const getDisplayText = (result) => {
    if (type === "skills") {
      return (
        result.skill?.skill_description || "No skill description available"
      );
    } else if (type === "interests") {
      return result.interest?.interest || "No interest available";
    }
    return "Unknown";
  };

  const formatDate = (dateString) => {
    if (!dateString) return "Unknown date";
    const date = new Date(dateString);
    return date.toLocaleDateString("en-US", {
      year: "numeric",
      month: "long",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit",
    });
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-2xl font-bold mb-4">
        Search Results for &quot;{query}&quot;
      </h1>
      <Link
        to="/dashboard"
        className="text-blue-600 hover:underline mb-4 inline-block"
      >
        &larr; Back to Dashboard
      </Link>
      {isLoading && <p className="text-gray-600">Loading...</p>}
      {error && <p className="text-red-600">{error}</p>}
      {!isLoading && results.length === 0 && (
        <p className="text-gray-600">
          No results found. Please try a different search term.
        </p>
      )}
      <ul className="space-y-4">
        {results.map((result, index) => (
          <li key={index} className="bg-white shadow rounded-lg p-4">
            <div className="font-semibold">
              {type === "skills" ? "Skill:" : "Interest:"}{" "}
              {getDisplayText(result)}
            </div>
            <div className="text-gray-600">
              Posted by: {result.user?.username || "Unknown user"} on{" "}
              {formatDate(result.created_at)}
            </div>
            {result.user?.discord_id && (
              <button
                onClick={() => handleInterested(result.user.discord_id)}
                className="mt-2 bg-violet-400 text-white py-1 px-3 rounded-md hover:bg-violet-500 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              >
                Contact on Discord
              </button>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default SearchResults;
