const http = require("http");

const basePort = 6000;
const targetBackends = 3;

for (let server = 0; server < targetBackends; server++) {
  const serverPort = server + basePort;

  http
    .createServer((req, res) => {
      res.write(`Server ${server} \r\n\r\n`);
      res.write(JSON.stringify(req.headers));
      res.end();
    })
    .listen(serverPort);
}

console.log(`Port range [${basePort}-${basePort + targetBackends - 1}], started listening`);
