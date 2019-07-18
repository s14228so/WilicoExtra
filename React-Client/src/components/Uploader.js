import React from "react";

class Uploader extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <h2>Uploaderだよ</h2>
        <form onSubmit={() => this.props.upload()}>
          <input type="file" />
          <input type="submit" value="投稿" />
        </form>
      </div>
    );
  }
}

export default Uploader;
