import { useEffect, useState } from "react";
import SkillsList from "./SkillsList";
import AddSkill from "./AddSkill";
import SearchSkills from "./SearchSkills";
import "../styles/Dashboard.css";
import AddInterest from "./AddInterest";
import InterestsList from "./InterestsList";
import SearchInterests from "./SearchInterests";
import { useNavigate } from "react-router-dom";
import { useCallback } from "react";
import api from "./Api";

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

  const handleLogout = () => {
    document.cookie =
      "session_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    navigate("/login");
  };

  if (loading) {
    return <div>Loading...</div>;
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
    <div className="dashboard">
      <div className="sidebar">
        <div className="user-info">
          <div className="user-info-top">
            <h2>Welcome, {user?.username}!</h2>
            <img src={user?.avatar_url} alt={`${user?.username}'s avatar`} />
          </div>
          <div className="user-info-bottom">
            <button className="logout" onClick={handleLogout}>
              Logout
            </button>
          </div>
        </div>
      </div>
      <div className="main-content">
        <div className="column">
          <div className="section" id="add-skill">
            <AddSkill refreshSkills={fetchData} />
          </div>
          <div className="section" id="add-interest">
            <AddInterest refreshInterests={fetchData} />
          </div>
        </div>
        <div className="column">
          <div className="section scrollable-section" id="search-skills">
            <SearchSkills />
          </div>
          <div className="section scrollable-section" id="search-interests">
            <SearchInterests />
          </div>
        </div>
        <div className="column">
          <div className="section scrollable-section">
            <SkillsList skills={skills} refreshSkills={fetchData} />
          </div>
          <div className="section scrollable-section">
            <InterestsList interests={interests} refreshInterests={fetchData} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
