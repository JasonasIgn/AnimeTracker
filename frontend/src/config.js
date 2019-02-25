const merge = require("lodash/merge");

const config = {
  all: {
    env: process.env.NODE_ENV || "development",
    isDev: process.env.NODE_ENV !== "production",
    isBrowser: typeof window !== "undefined",
    apiUrl: "http://localhost:1337",
    utils: {
      maxSuggestionsInList: 6
    }
  },
  development: {},
  production: {}
};

module.exports = merge(config.all, config[config.all.env]);
