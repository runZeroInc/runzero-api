const yaml = require('js-yaml');
const fs   = require('fs');

// const indexFile = 'runzero-api/index.html';
const indexFile = 'runzero-api/swagger-initializer.js';
const specYAML = 'runzero-api.yml';

try {
  const doc = yaml.load(fs.readFileSync(specYAML, 'utf8'));
  const jsonData = JSON.stringify(doc, null, 2);
  const index = fs.readFileSync(indexFile, 'utf8');
  const newIndex = index.replace(/url:.*swagger.io.*\.json",/m, `  spec: ${jsonData},`);
  if (index === newIndex) {
    console.log("Index replacement failed");
    process.exit(1);
  }
  fs.writeFileSync(indexFile, newIndex, 'utf8');
  console.log(`Updated ${indexFile}`);
} catch (e) {
  console.log(e);
}
