import { useState } from "react";
import mattImage from "../assets/matt.jpg";

const Matt = () => {
  const [isDark, setIsDark] = useState(false);

  const imageStyle = {
    width: "10%",
    height: "auto",
    filter: isDark ? "brightness(50%)" : undefined,
    border: isDark ? "3px gold dashed" : undefined,
  };

  const handleClick = () => {
    setIsDark(!isDark);
    console.log(isDark);
  };
  return (
    <div className="h-100 d-flex justify-content-center align-items-center">
      <img
        src={mattImage}
        alt="Matt"
        className="img-fluid rounded"
        style={imageStyle}
        onClick={handleClick}
      />
    </div>
  );
};

export default Matt;
