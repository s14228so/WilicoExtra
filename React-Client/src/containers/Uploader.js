import React from "react";
import Upload from "../components/Uploader";
// import { increment } from "../actions/app.jsx";

import axios from "axios";

export class Uploader extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      plans: []
    };
  }
  upload(file) {
    const url = "http://localhost:8080/upload";
    // リクエスト直前に実行
    axios.post(url, {
      params: { file: file }
    });
  }

  render() {
    return (
      <div>
        <Upload getPlans={this.upload()} />
      </div>
    );
  }
}

export default Uploader;
