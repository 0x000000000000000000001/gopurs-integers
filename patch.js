const fs = require('fs');
const code = fs.readFileSync('test/Test/Data/Int.purs', 'utf8');
let i = 0;
const newCode = code.replace(/assert \$/g, () => `log "assert ${++i}" >>= \\_ -> assert $`);
fs.writeFileSync('test/Test/Data/Int.purs', newCode);
