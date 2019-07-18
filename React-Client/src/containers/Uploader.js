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
  upload = e => {
    e.preventDefault();
  };

  render() {
    return (
      <div>
        <Upload upload={this.upload} />
      </div>
    );
  }
}

export default Uploader;
