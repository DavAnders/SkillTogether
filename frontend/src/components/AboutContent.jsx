const AboutContent = () => (
  <>
    <div className="bg-white sm:w-1/2 p-5 flex flex-grow flex-col">
      <div>
        <h2 className="my-4 font-bold text-3xl  sm:text-4xl">
          What is <span className="text-violet-400">SkillTogether</span>?
        </h2>
      </div>
      <div>
        <p className="text-gray-700">
          SkillTogether is a way to connect with people that share similar
          interests and a platform for swapping skills.
        </p>
        <p className="text-gray-700">
          For example, if you&apos;re interested in learning{" "}
          <em className="font-bold">Japanese</em>, and someone else is looking
          to learn <em className="font-bold">English</em>, you can click the
          link on their posting to be redirected to their Discord profile to get
          in contact with them.
        </p>
      </div>
    </div>
  </>
);

export default AboutContent;
