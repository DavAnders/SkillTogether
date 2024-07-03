import { Link } from "react-router-dom";
import skillImg from "../assets/ST-dark.png";

const Footer = () => {
  return (
    <footer className="bg-gray-800 text-white py-8">
      <div className="container mx-auto px-4">
        <div className="flex flex-col md:flex-row items-center justify-between gap-6">
          <div className="flex items-center justify-center">
            <div
              className="w-10 h-5 bg-cover bg-center rounded-full mr-2"
              style={{ backgroundImage: `url(${skillImg})` }}
            />
            <span className="text-xl font-semibold whitespace-nowrap">
              SkillTogether
            </span>
          </div>

          <ul className="flex flex-wrap justify-center gap-4">
            <li>
              <Link
                to="/privacy"
                className="hover:text-blue-500 transition-colors"
              >
                Privacy Policy
              </Link>
            </li>
            <li>
              <Link
                to="/terms"
                className="hover:text-blue-500 transition-colors"
              >
                Terms of Service
              </Link>
            </li>
          </ul>

          <p className="text-center text-sm">
            &copy; {new Date().getFullYear()} SkillTogether. All rights
            reserved.
          </p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
