import React, { Component, Fragment } from "react";
import axios from "axios";
import "./App.css";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      plans: []
    };
    this.getPlans = this.getPlans.bind(this);
  }

  componentDidMount() {
    this.getPlans();
  }

  async getPlans() {
    const { data } = await axios.get("http://localhost:8080/plans");
    console.log(data);
    this.setState({
      plans: [...data]
      //展開記法 (スプレッド記法) は、式を複数の引数または複数の要素に展開して、それぞれ関数呼び出しまたは配列リテラルに渡します。
    });
  }

  render() {
    const plans = this.state.plans.map(plan => {
      return <li key={plan.id}> {plan.title} </li>;
    });

    return (
      <Fragment>
        <div>
          <button onClick={this.getPlans}> getData </button> <ul> {plans} </ul>
        </div>
      </Fragment>
    );
  }
}

export default App;
