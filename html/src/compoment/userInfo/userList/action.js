import * as types from "./types";
// import articleService from "../../../../common/js/services/article"

export const setUserCount = ({commit, state}, count) => {
    commit(types.setUserCount, count);
};
