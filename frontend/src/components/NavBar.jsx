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
    event.preventDefault();
    event.stopPropagation();
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
      <div className="flex items-center justify-between">
        <div className="flex justify-start items-center flex-1">
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
          className="p-2 w-10 h-10 text-gray-400 rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-600 md:hidden"
        >
          <span className="sr-only">Open main menu</span>
          <svg
            className="w-6 h-6"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M4 6h16M4 12h16m-7 6h7"
            />
          </svg>
        </button>

        <div className="flex-1 justify-center hidden md:flex">
          <ul className="flex justify-center space-x-8">
            <li>
              <Link
                to="/dashboard"
                className="py-2 px-3 text-white hover:bg-gray-700 md:hover:bg-transparent md:border-0 md:p-0 md:hover:text-blue-500"
              >
                Dashboard
              </Link>
            </li>
            <li>
              <Link
                to="/about"
                className="py-2 px-3 text-white hover:bg-gray-700 md:hover:bg-transparent md:border-0 md:p-0 md:hover:text-blue-500"
              >
                About
              </Link>
            </li>
          </ul>
        </div>

        <div className="items-center flex-1 justify-end hidden md:flex">
          {user ? (
            <>
              <img
                className="rounded-full w-10 h-10 mr-4"
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

        {/* Mobile menu dropdown */}
        <div
          className={`absolute top-full right-0 w-full md:hidden bg-gray-800 shadow-md ${
            isMenuOpen ? "block" : "hidden"
          }`}
        >
          <ul className="flex flex-col p-4">
            <li>
              <Link
                to="/dashboard"
                className="block py-2 px-3 text-white rounded hover:bg-gray-700"
                onClick={toggleMenu}
              >
                Dashboard
              </Link>
            </li>
            <li>
              <Link
                to="/about"
                className="block py-2 px-3 text-white rounded hover:bg-gray-700"
                onClick={toggleMenu}
              >
                About
              </Link>
            </li>
            {user && (
              <li>
                <button
                  onClick={(e) => {
                    toggleMenu();
                    handleLogout(e);
                  }}
                  className="w-full text-left py-2 px-3 rounded hover:bg-gray-700"
                >
                  Logout
                </button>
              </li>
            )}
          </ul>
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
