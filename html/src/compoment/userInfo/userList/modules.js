import {setUserCount} from "./types";

const state = {
    userCount: 1
};

const mutations = {
    //文章列表
    [setUserCount](state, count){
        state.userCount += count;
    }

};

export default  {
    state,
    mutations
};
