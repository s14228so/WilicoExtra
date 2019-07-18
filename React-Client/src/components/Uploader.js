import React from "react";

class Uploader extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    console.log(this.props);
    return (
      <div>
        <h2>Uploaderだよ</h2>
        <form
          action="http://localhost:8080/upload"
          method="post"
          encType="multipart/form-data"
        >
          <input type="file" name="files" multiple />
          <input type="submit" value="投稿" />
        </form>
      </div>
    );
  }
}

export default Uploader;
