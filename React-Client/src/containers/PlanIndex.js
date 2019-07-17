import React from "react";
import Plans from "../components/PlanIndex";
// import { increment } from "../actions/app.jsx";

import axios from "axios";

export class PlanIndex extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      plans: []
    };
  }

  componentDidMount() {
    this.getPlans();
  }

  getPlans() {
    const url = "http://localhost:8080/plans";
    // リクエスト直前に実行
    axios.get(url).then(res => {
      console.log(res.data);
      this.setState({
        plans: res.data
      });
    });
  }

  render() {
    return (
      <div>
        <Plans plans={this.state.plans} />
      </div>
    );
  }
}

export default PlanIndex;
