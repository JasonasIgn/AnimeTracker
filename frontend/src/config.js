const merge = require("lodash/merge");

const config = {
  all: {
    env: process.env.NODE_ENV || "development",
    isDev: process.env.NODE_ENV !== "production",
    isBrowser: typeof window !== "undefined",
    apiUrl: "",
    utils: {
      maxSuggestionsInList: 6
    }
  },
  test: {},
  development: {},
  production: {}
};

module.exports = merge(config.all, config[config.all.env]);
