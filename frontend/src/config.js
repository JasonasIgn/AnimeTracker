const merge = require("lodash/merge");

const config = {
  all: {
    env: process.env.NODE_SURROUNDING || "development",
    isProd: process.env.NODE_SURROUNDING == "production",
    isDev: process.env.NODE_SURROUNDING == "development",
    isQa: process.env.NODE_SURROUNDING == "qa",
    isBrowser: typeof window !== "undefined",
    apiUrl: "http://api.uberlotti.local/graphql"
  },
  test: {},
  development: {}
};

module.exports = merge(config.all, config[config.all.env]);
