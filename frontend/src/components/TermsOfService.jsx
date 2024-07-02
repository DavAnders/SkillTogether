import MarkdownRenderer from "./MarkdownRenderer";
import tosContent from "../legal/terms-of-service.md";

const TermsOfService = () => {
  return (
    <div className="flex flex-grow justify-center">
      <div className="terms-of-service">
        <h1 className="text-3xl font-bold text-center my-6">
          Terms of Service
        </h1>
        <MarkdownRenderer content={tosContent} />
      </div>
    </div>
  );
};

export default TermsOfService;
