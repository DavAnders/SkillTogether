import { Link } from "react-router-dom";
import stDark from "../assets/ST-dark.png";

const Home = () => {
  return (
    <div className="text-black flex flex-col justify-center items-center flex-grow bg-gray-100 py-10 min-h-full">
      <header className="text-center">
        <h1 className="text-4xl font-bold mb-4">Welcome to SkillTogether</h1>
        <p className="text-lg mb-2">
          Your platform to learn and share skills with others.
        </p>
        <p className="text-lg mb-4">
          Find new friends with similar interests and connect on Discord!
        </p>
        <div className="space-x-4">
          <Link
            to="/login"
            className="px-6 py-3 bg-violet-400 text-white font-semibold rounded-lg shadow-md hover:bg-violet-500 focus:outline-none focus:ring-2 focus:ring-violet-300 focus:ring-opacity-75"
          >
            Get Started
          </Link>
          <Link
            to="/about"
            className="px-6 py-3 bg-gray-600 text-white font-semibold rounded-lg shadow-md hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-75"
          >
            Learn More
          </Link>
        </div>
      </header>

      <div className="mt-10 flex-grow flex items-center justify-center">
        <img
          className="w-64 h-64 rounded-full"
          src={stDark}
          alt="SkillTogether"
        />
      </div>
    </div>
  );
};

export default Home;
