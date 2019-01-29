const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const CopyWebpackPlugin = require("copy-webpack-plugin");
const chalk = require("chalk");
const webpack = require("webpack");
const config = require("./src/config");

console.log(chalk.bold(`\n\nBuild enviroment: ${chalk.green(config.env)}\n`));

const pages = [
  new HtmlWebpackPlugin({
    filename: "index.html",
    template: "./src/index.html",
    inject: true
  }),
  new CopyWebpackPlugin([{ from: "src/resources/static", to: "static" }]),
  new webpack.DefinePlugin({
    "process.env.NODE_SURROUNDING": JSON.stringify(process.env.NODE_SURROUNDING)
  })
];

module.exports = {
  mode: config.isDev ? "development" : "production",

  entry: "./src/index.jsx",

  output: {
    filename: "[name].[hash].js",
    path: path.resolve("./dist"),
    publicPath: "/"
  },

  devtool: config.isDev ? "source-map" : false,
  target: "web",

  module: {
    rules: [
      {
        test: /\.(png|jpe?g)$/,
        exclude: /\/fonts\//,
        use: [
          {
            loader: "file-loader",
            options: {
              name: "img/[name].[hash:16].[ext]"
            }
          },
          {
            loader: "sharp-image-loader",
            options: {
              withMetadata: true,
              jpegQuality: 80,
              jpegProgressive: true,
              pngProgressive: true,
              pngCompressionLevel: 6,
              webpQuality: 80,
              webpAlphaQuality: 100,
              tiffQuality: 80
            }
          }
        ]
      },
      {
        test: /\.(svg|gif)$/,
        exclude: /\/fonts\//,
        use: [
          {
            loader: "file-loader",
            options: {
              name: "img/[name].[hash:16].[ext]"
            }
          }
        ]
      },
      {
        test: /(\.eot|\.otf|\.woff|\.woff2|\.ttf|\/fonts\/.*\.svg)$/,
        use: [
          {
            loader: "file-loader",
            options: {
              name: "fonts/[name].[hash:16].[ext]"
            }
          }
        ]
      },
      {
        test: /favicon\.ico$/,
        use: [
          {
            loader: "file-loader?name=[name].[ext]",
            options: {
              name: "[name].[ext]"
            }
          }
        ]
      },
      {
        test: /\.css$/,
        use: [
          {
            loader: "file-loader",
            options: {
              name: "css/[name].[hash:16].css"
            }
          },
          {
            loader: "extract-loader"
          },
          {
            loader: "css-loader",
            options: {
              minimize: config.isDev,
              sourceMap: config.isDev
            }
          }
        ]
      },
      {
        test: /\.scss$/,
        use: [
          {
            loader: "file-loader",
            options: {
              name: "css/[name].[hash:16].css"
            }
          },
          {
            loader: "extract-loader"
          },
          {
            loader: "css-loader",
            options: {
              minimize: !config.isDev,
              sourceMap: config.isDev
            }
          },
          {
            loader: "resolve-url-loader"
          },
          {
            loader: "sass-loader",
            options: {
              sourceMap: true
            }
          }
        ]
      },
      {
        test: /\.jsx?$/,
        exclude: /(webpack-dev-server|node_modules)/,
        use: [
          {
            loader: "babel-loader"
          }
        ]
      },
      {
        test: /\.html$/,
        use: [
          {
            loader: "html-loader",
            options: {
              interpolate: true,
              minimize: config.isDev,
              removeComments: !config.isDev,
              collapseWhitespace: true,
              attrs: ["img:src", "link:href", "script:src", "div:data-bg"]
            }
          }
        ]
      }
    ]
  },

  devServer: {
    host: process.env.HOST || "localhost",
    port: process.env.PORT || 3000,
    publicPath: "/",
    historyApiFallback: true,
    compress: true,
    hot: false,
    https: false,
    noInfo: false
  },

  resolve: {
    extensions: [".mjs", ".web.js", ".js", ".json", ".web.jsx", ".jsx"]
  },

  plugins: [...pages]
};
