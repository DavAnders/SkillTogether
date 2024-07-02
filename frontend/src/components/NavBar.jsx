import { Link, useNavigate } from "react-router-dom";
import skillImg from "../assets/ST-dark.png";
import PropTypes from "prop-types";
import { useState } from "react";

const NavBar = ({ user, setUser }) => {
  const navigate = useNavigate();
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  const handleLogout = async (event) => {
    event.preventDefault(); // Had to add because browser wasn't redirecting without it
    event.stopPropagation(); // Prevent the navbar toggle from firing
    try {
      const response = await fetch("http://localhost:8080/api/logout", {
        method: "POST",
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error("Failed to log out");
      }

      setUser(null);
      navigate("/login");
    } catch (error) {
      console.error("Logout error:", error);
      alert("Logout failed. Please try again.");
    }
  };

  return (
    <nav className="bg-gray-800 p-4 text-white sticky top-0 z-10">
      <div className="flex flex-wrap items-center justify-between">
        <div className="flex items-center flex-1 justify-center">
          <div
            className="w-10 h-5 bg-cover bg-center rounded-full"
            style={{ backgroundImage: `url(${skillImg})` }}
          />
          <span className="self-center text-xl font-semibold whitespace-nowrap hidden md:block">
            SkillTogether
          </span>
        </div>

        <button
          onClick={toggleMenu}
          className="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-400 rounded-lg md:hidden hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-600"
        >
          <span className="sr-only">Open main menu</span>
          <svg
            className="w-5 h-5"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 17 14"
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M1 1h15M1 7h15M1 13h15"
            />
          </svg>
        </button>

        <div
          className={`${
            isMenuOpen ? "block" : "hidden"
          } w-full md:block md:w-auto `}
          id="navbar-default"
        >
          <ul className="font-medium flex flex-col p-4 md:p-0 mt-4 border border-gray-700 rounded-lg bg-gray-800 md:flex-row md:space-x-8 md:mt-0 md:border-0 md:bg-transparent">
            <li>
              <Link
                to="/dashboard"
                className="block py-2 px-3 text-white rounded hover:bg-gray-700 md:hover:bg-transparent md:border-0 md:hover:text-blue-500 md:p-0"
                onClick={toggleMenu}
              >
                Dashboard
              </Link>
            </li>
            <li>
              <Link
                to="/about"
                className="block py-2 px-3 text-white rounded hover:bg-gray-700 md:hover:bg-transparent md:border-0 md:hover:text-blue-500 md:p-0"
                onClick={toggleMenu}
              >
                About
              </Link>
            </li>
            <li className="md:hidden">
              {user ? (
                <button
                  onClick={(e) => {
                    toggleMenu();
                    handleLogout(e);
                  }}
                  className="w-full text-left block py-2 px-3 text-white rounded hover:bg-gray-700 md:hover:bg-transparent md:border-0 md:hover:text-blue-500 md:p-0"
                >
                  Logout
                </button>
              ) : (
                <Link
                  to="/login"
                  className="block py-2 px-3 text-white rounded hover:bg-gray-700 md:hover:bg-transparent md:border-0 md:hover:text-blue-500 md:p-0"
                  onClick={toggleMenu}
                >
                  Login
                </Link>
              )}
            </li>
          </ul>
        </div>

        <div className="hidden md:flex items-center gap-3 justify-center flex-1">
          {user ? (
            <>
              <img
                className="rounded-full w-10 h-10"
                src={user?.avatar_url}
                alt={`${user?.username}'s avatar`}
              />
              <button
                onClick={handleLogout}
                className="bg-violet-400 rounded-lg px-4 py-1"
              >
                Logout
              </button>
            </>
          ) : (
            <Link to="/login">
              <button className="bg-violet-400 rounded-lg px-4 py-1">
                Login
              </button>
            </Link>
          )}
        </div>
      </div>
    </nav>
  );
};

NavBar.propTypes = {
  user: PropTypes.object,
  setUser: PropTypes.func,
};

export default NavBar;
