import { useEffect } from "react";
import useNavigate from "react-router-dom";
import api from "./Api";

const Callback = () => {
  const navigate = useNavigate();

  useEffect(() => {
    const code = new URLSearchParams(window.location.search).get("code");
    if (code) {
      api
        .get(`/auth/discord/callback?code=${code}`, { withCredentials: false })
        .then((response) => {
          const { session_token } = response.data;
          localStorage.setItem("session_token", session_token);
          navigate("/dashboard");
        })
        .catch((error) => {
          // Handle error
          console.error("Error during login", error);
        });
    }
  }, [navigate]);

  return <div>Loading...</div>;
};

export default Callback;
