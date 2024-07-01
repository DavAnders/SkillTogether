import ReactMarkdown from "react-markdown";
import rehypeRaw from "rehype-raw";
import PropTypes from "prop-types";

const MarkdownRenderer = ({ content }) => {
  return (
    <div className="markdown-content max-w-3xl mx-auto p-6">
      <ReactMarkdown
        rehypePlugins={[rehypeRaw]}
        components={{
          h1: (props) => <h1 className="text-2xl font-bold my-4" {...props} />,
          h2: (props) => (
            <h2 className="text-xl font-semibold my-3" {...props} />
          ),
          h3: (props) => <h3 className="text-lg font-medium my-2" {...props} />,
          p: (props) => <p className="my-2" {...props} />,
          ul: (props) => (
            <ul className="list-disc list-inside my-2" {...props} />
          ),
          ol: (props) => (
            <ol className="list-decimal list-inside my-2" {...props} />
          ),
          li: (props) => <li className="ml-4" {...props} />,
          a: (props) => (
            <a className="text-blue-600 hover:underline" {...props} />
          ),
        }}
      >
        {content}
      </ReactMarkdown>
    </div>
  );
};

MarkdownRenderer.propTypes = {
  content: PropTypes.string.isRequired,
};

export default MarkdownRenderer;
