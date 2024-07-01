import skillImg from "../assets/SkillTogether.png";
import AboutContent from "./AboutContent";

const About = () => {
  return (
    <div className="flex justify-center items-center h-screen">
      <div className="sm:flex items-center max-w-screen-xl">
        <div className="object-center text-center scale-75">
          <img src={skillImg} alt="SkillTogether" />
        </div>
        <AboutContent />
      </div>
    </div>
  );
};

export default About;
