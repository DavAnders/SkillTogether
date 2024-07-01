import { Link } from "react-router-dom";
import LoginButton from "./LoginButton";

const Login = () => {
  return (
    <div className="bg-slate-50 dark:bg-slate-900">
      <div className="flex h-screen flex-col items-center justify-center">
        <div className="max-h-auto mx-auto max-w-xl">
          <div className="mb-8 space-y-3">
            <p className="text-xl font-semibold dark:text-white">Login</p>
            <p className="text-gray-500 dark:text-white">
              Connect your Discord account to login. <br />
              No need for passwords -- like magic!
            </p>
          </div>
          <form className="w-full">
            <div className="mb-10 space-y-3">
              <LoginButton />
            </div>
          </form>
          <div className="text-center dark:text-white">
            <Link className="text-blue-500 dark:text-blue-200" to="/about">
              About
            </Link>{" "}
            SkillTogether
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
