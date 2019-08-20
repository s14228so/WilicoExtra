import React from "react";
import { Link } from "react-router-dom";

const Header = () => {
  const style = {
    textDecoration: "none",
    color: "#000",
    lineHeight: "70px",
    fontSize: "12px"
  };

  const links = [
    {
      title: "Home",
      path: "/"
    },
    {
      title: "About",
      path: "/about"
    },
    {
      title: "Todo",
      path: "/todo"
    }
  ];

  const header = {
    width: "100vw",
    background: "lightGray",
    height: "70px"
  };

  const items = links.map((item, i) => {
    return (
      <Link style={style} key={i} to={item.path}>
        {item.title}
      </Link>
    );
  });
  return (
    <header style={header}>
      <ul
        style={{
          listStyle: "none",
          display: "flex",
          width: "25vw",
          justifyContent: "space-around"
        }}
      >
        {items}
      </ul>
    </header>
  );
};

export default Header;
