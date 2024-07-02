import { Outlet, useLocation, useNavigate } from "react-router-dom";
import NavBar from "./components/NavBar";
import { useEffect, useState } from "react";
import api from "./components/Api";
import Footer from "./components/Footer";
import ScrollToTop from "./components/ScrollToTop";

function App() {
  const [user, setUser] = useState(null);

  const navigate = useNavigate();
  const location = useLocation();
  const showNavBar = location.pathname !== "/login";

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await api.get("/api/me");
        setUser(response.data);
      } catch (error) {
        console.error("Failed to fetch user:", error);
        // if (error.response && error.response.status === 401) {
        //   navigate("/login");
        // }
      }
    };

    fetchUser();
  }, [navigate]);

  return (
    <div className="flex flex-col min-h-screen bg-slate-300">
      <ScrollToTop />
      {showNavBar && <NavBar user={user} setUser={setUser} />}
      <main className="flex-grow flex">
        <Outlet context={{ user, setUser }} />
      </main>
      {showNavBar && <Footer />}
    </div>
  );
}

export default App;
