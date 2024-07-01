import fs from "fs";
import path from "path";

export default function markdown() {
  return {
    name: "transform-markdown",
    transform(src, id) {
      if (id.endsWith(".md")) {
        const filePath = path.resolve(id);
        const content = fs.readFileSync(filePath, "utf-8");
        const code = `export default ${JSON.stringify(content)}`;
        return { code };
      }
    },
  };
}
