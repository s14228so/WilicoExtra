import React, { Component, Fragment } from "react";
import axios from "axios";
import "./App.css";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      users: []
    };
    this.getData = this.getData.bind(this);
  }

  getData() {
    axios.get("http://localhost:8080/plans").then(results => {
      const data = results.data;
      console.log(data);
      this.setState({
        users: [...data]
      });
    });
  }

  render() {
    const users = this.state.users.map(user => {
      return <li key={user.id}> {user.username} </li>;
    });

    return (
      <Fragment>
        <div>
          <button onClick={this.getData}> getData </button> <ul> {users} </ul>{" "}
        </div>{" "}
      </Fragment>
    );
  }
}

export default App;
