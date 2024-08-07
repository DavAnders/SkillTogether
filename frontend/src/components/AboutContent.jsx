const AboutContent = () => (
  <div className="bg-white p-5 flex flex-col">
    <h2 className="mb-6 font-bold text-3xl sm:text-4xl">
      What is <span className="text-violet-400">SkillTogether</span>?
    </h2>
    <div className="space-y-4">
      <p className="text-gray-700">
        SkillTogether is a way to connect with people that share similar
        interests and a platform for swapping skills.
      </p>
      <p className="text-gray-700">
        For example, if you&quot;re interested in learning{" "}
        <em className="font-bold">Japanese</em>, and someone else is looking to
        learn <em className="font-bold">English</em>, you can click the link on
        their posting to be redirected to their Discord profile to get in
        contact with them.
      </p>
    </div>
  </div>
);

export default AboutContent;
