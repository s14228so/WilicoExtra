// export const plus = num => {
//   return { type: "PLUS", payload: { num: num } };
// };

import axios from "axios";

const startRequest = () => {
  return {
    type: "START_REQUEST"
  };
};

const receiveData = data => {
  return {
    type: "RECEIVE_DATA",
    payload: data
  };
};

const getUrls = word => {
  return dispatch => {
    const url = "http://localhost:8080/plans";
    // リクエスト直前に実行
    dispatch(startRequest());
    axios.get(url).then(res => {
      console.log(res.data);
      const data = res.data;
      const imageUrlList = data.map(item => item.images.downsized.url);
      dispatch(receiveData(imageUrlList));
    });
  };
};

export default getUrls;
