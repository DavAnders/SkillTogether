import skillImg from "../assets/SkillTogether.png";
import AboutContent from "./AboutContent";

const About = () => {
  return (
    <div className="bg-white flex justify-center items-center min-h-screen w-full px-4 py-8">
      <div className="container mx-auto flex flex-col lg:flex-row items-center justify-center gap-8">
        <div className="w-full lg:w-1/2 flex justify-center">
          <img
            src={skillImg}
            alt="SkillTogether"
            className="max-w-full h-auto"
          />
        </div>
        <div className="w-full lg:w-1/2 max-w-2xl">
          <AboutContent />
        </div>
      </div>
    </div>
  );
};

export default About;
