import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useCallback } from "react";
import api from "./api";
import LoadingSpinner from "./LoadingSpinner";
import SkillsList from "./SkillsList";
import InterestsList from "./InterestsList";
import Search from "./Search";
import AddSkill from "./AddSkill";
import AddInterest from "./AddInterest";

const Dashboard = () => {
  const [user, setUser] = useState(null);
  const [skills, setSkills] = useState([]);
  const [interests, setInterests] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  const fetchData = useCallback(async () => {
    try {
      const userResponse = await api.get("/api/me");
      setUser(userResponse.data);

      const skillsResponse = await api.get("/api/skills");
      setSkills(skillsResponse.data);

      const interestsResponse = await api.get("/api/interests");
      setInterests(interestsResponse.data);

      setLoading(false);
    } catch (error) {
      console.error("Error fetching data:", error);
      setLoading(false);
      if (error.response && error.response.status === 401) {
        navigate("/login");
      }
    }
  }, [navigate]);

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-screen">
        <div className="my-auto">
          <LoadingSpinner />
        </div>
      </div>
    );
  }

  if (!user) {
    return (
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
        }}
      >
        <p>Not logged in</p>
        <button onClick={() => navigate("/login")}>Go to Login Page</button>
      </div>
    );
  }

  return (
    <div className="bg-gray-100 min-h-80">
      <div className="container mx-auto px-4 py-8">
        <div className="bg-white shadow-lg rounded-lg overflow-hidden">
          {/* User Info Section */}
          <div className="bg-gray-800 text-white p-6">
            <div className="flex items-center">
              <img
                src={user.avatar_url}
                alt={`${user.username}'s avatar`}
                className="w-20 h-20 rounded-full mr-4"
              />
              <div>
                <h1 className="text-2xl font-bold">
                  Welcome, {user.username}!
                </h1>
                <p className="text-blue-200">Dashboard</p>
              </div>
            </div>
          </div>

          {/* Main Content Grid */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 p-6">
            {/* Skills Section */}
            <div className="bg-white shadow rounded-lg p-4 relative">
              <div className="flex justify-between items-center mb-4">
                <h2 className="text-xl font-semibold">Skills</h2>
                <div className="absolute top-4 right-4">
                  <AddSkill refreshSkills={fetchData} />
                </div>
              </div>
              <SkillsList skills={skills} refreshSkills={fetchData} />
            </div>

            {/* Interests Section */}
            <div className="bg-white shadow rounded-lg p-4 relative">
              <div className="flex justify-between items-center mb-4">
                <h2 className="text-xl font-semibold">Interests</h2>
                <div className="absolute top-4 right-4">
                  <AddInterest refreshInterests={fetchData} />
                </div>
              </div>
              <InterestsList
                interests={interests}
                refreshInterests={fetchData}
              />
            </div>

            {/* Search Section */}
            <div className="bg-white shadow rounded-lg p-4">
              <h2 className="text-xl font-semibold mb-4">Search</h2>
              <Search />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
