import React from "react";
import PlanIndex from "./containers/PlanIndex";

class App extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <PlanIndex />
      </div>
    );
  }
}

export default App;
