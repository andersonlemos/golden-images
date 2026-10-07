const http = require("node:http");

const PORT = process.env.PORT || 3000;

const server = http.createServer((req, res) => {
  res.setHeader("Content-Type", "application/json");

  if (req.url === "/health") {
    res.writeHead(200);
    res.end(
      JSON.stringify({
        status: "ok",
        environment: process.env.NODE_ENV || "undefined",
      })
    );

    return;
  }

  if (req.url === "/") {
    res.writeHead(200);
    res.end(
      JSON.stringify({
        message: "Hello from Node.js Golden Image",
        nodeVersion: process.version,
        environment: process.env.NODE_ENV || "undefined",
      })
    );

    return;
  }

  res.writeHead(404);
  res.end(
    JSON.stringify({
      error: "Not Found",
    })
  );
});

server.listen(PORT, () => {
  console.log(`API running on port ${PORT}`);
});