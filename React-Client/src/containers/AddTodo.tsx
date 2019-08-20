import { useDispatch } from "react-redux";
import React, { FunctionComponent, useState } from "react";
import { add, addError } from "../actions";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Error from "../types/error";
import newTodo from "../types/newTodo";

interface MessageInputEvent extends React.FormEvent<HTMLInputElement> {
  target: HTMLInputElement;
}

const AddTodo: FunctionComponent = (): JSX.Element => {
  // type FormElem = React.FormEvent<HTMLInputElement>;
  const [newTodo, setTodo] = useState<newTodo>({
    title: "",
    count: 0
  });
  const [error, setError] = useState<Error>({
    message: ""
  });
  const dispatch = useDispatch();

  const handleSubmit = (e: React.MouseEvent<HTMLAnchorElement, MouseEvent>) => {
    e.preventDefault();

    if (newTodo.title === "") {
      setError({ message: "文字を入力してね！" });
      return;
    } else if (newTodo.title.length > 35) {
      setError({ message: "文字数オーバーだよ！" });
      return;
    }
    dispatch(add(newTodo.title));

    setError({ message: "" });
    dispatch(addError(""));

    setTodo({
      title: "",
      count: 0
    });
  };
  const handleChange = (e: MessageInputEvent): void => {
    const length: number = e.target.value.length;
    if (length > 35) {
      setError({ message: "文字数オーバーだよ！" });
    } else {
      setError({ message: "" });
    }
    setTodo({
      title: e.target.value,
      count: length
    });
  };
  // this.props.dispatch(add("aaaaa"));
  return (
    <div
      style={{
        width: "50%",
        textAlign: "center",
        margin: "0 auto"
      }}
    >
      <form>
        <input
          style={{
            padding: "10px 15px",
            width: "80%",
            marginTop: "20px",
            borderRadius: "10px"
          }}
          type="text"
          value={newTodo.title}
          onChange={handleChange}
        />{" "}
        <Button variant="contained" color="primary" onClick={handleSubmit}>
          Add
        </Button>
        <br />
        {newTodo.count}/35
      </form>
      <div className="error" style={{ color: "red" }}>
        {error.message}
      </div>
    </div>
  );
};

export default AddTodo;
