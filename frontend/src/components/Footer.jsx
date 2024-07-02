import { Link } from "react-router-dom";
import skillImg from "../assets/ST-dark.png";

const Footer = () => {
  return (
    <footer className="bg-gray-800 text-white py-8">
      <div className="flex flex-wrap items-center justify-between px-4">
        <div className="flex-1 justify-center flex items-center mb-4 md:mb-0">
          <div
            className="w-10 h-5 bg-cover bg-center rounded-full"
            style={{ backgroundImage: `url(${skillImg})` }}
          />
          <span className="self-center text-xl font-semibold whitespace-nowrap">
            SkillTogether
          </span>
        </div>

        <div className="flex-1 w-full justify-center md:w-auto text-center md:text-left">
          <ul className="flex justify-center gap-5 mb-4 md:mb-0">
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
        </div>

        <div className="flex-1 ml-auto flex justify-center md:w-auto text-center md:text-left">
          <p>
            &copy; {new Date().getFullYear()} SkillTogether. All rights
            reserved.
          </p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
