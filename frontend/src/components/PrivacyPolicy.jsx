import MarkdownRenderer from "./MarkdownRenderer";
import privacyPolicyContent from "../legal/privacy-policy.md";

const PrivacyPolicy = () => {
  return (
    <div className="flex flex-grow justify-center">
      <div className="privacy-policy">
        <h1 className="text-3xl font-bold text-center my-6">Privacy Policy</h1>
        <MarkdownRenderer content={privacyPolicyContent} />
      </div>
    </div>
  );
};

export default PrivacyPolicy;
