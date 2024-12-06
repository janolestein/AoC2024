const fs = require('node:fs');
fs.readFileSync('rules.txt', 'utf8', (err, rules) => {
  if (err) {
    console.error(err);
    return;
  }
  console.log(rules);
});

fs.readFileSync('updates.txt', 'utf8', (err, updates) => {
  if (err) {
    console.error(err);
    return;
  }
  console.log(updates);
});
