const path = require("path");
const Dotenv = require("dotenv-webpack");

const rules = [
  {
    test: /\.tsx?/,
    exclude: /node_modules/,
    loader: "babel-loader"
  },
  {
    test: /\.css/,
    use: [
      "style-loader",
      {
        loader: "css-loader",
        options: { url: false }
      }
    ]
  }
];
module.exports = {
  target: "node",
  mode: "development",
  entry: ["@babel/polyfill", "./src/index.tsx"],
  output: {
    path: path.join(__dirname, "dist"),
    filename: "bundle.js"
  },
  module: { rules },
  resolve: {
    extensions: [".ts", ".tsx", ".js"]
  },
  devServer: {
    contentBase: "./",
    port: 3000
  },
  plugins: [new Dotenv()]
};
