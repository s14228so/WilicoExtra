import React from "react";
export default class PlanIndex extends React.Component {
  listrendering() {
    const plans = this.props.plans.map((plan, index) => {
      return (
        <li key={index}>
          {plan.title}, {plan.price}
        </li>
      );
    });

    return <ul>{plans}</ul>;
  }

  render() {
    console.log(this.props);
    return (
      <div>
        <div>{this.listrendering()}</div>
        {/* <button onClick={() => this.props.handleClick}>+</button>
        <button onClick={() => this.props.handleClick}>-</button> */}
      </div>
    );
  }
}
