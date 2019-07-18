import React from "react";
import PlanIndex from "./containers/PlanIndex";
import Uploader from "./containers/Uploader";

class App extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <PlanIndex />
        <Uploader />
      </div>
    );
  }
}

export default App;
