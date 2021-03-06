const PROXY_CONFIG = [
  {
    context: [
      "/api/"
    ],
    target: "http://0.0.0.0:8080",
    secure: false,
    logLevel: "debug",
    changeOrigin: true,
    "pathRewrite": {
      "^/api": ""
    },
    cookieDomainRewrite: {
      "*": ""
    }
  }
];

module.exports = PROXY_CONFIG;
