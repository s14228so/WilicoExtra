import React from "react";
import Todo from "./Todo";

const TodoList = ({ todos, remove, filter }) => {
  const items = todos.map(todo => {
    return <Todo key={todo.id} todo={todo} remove={remove} />;
  });
  return <div>{items}</div>;
};

export default TodoList;
