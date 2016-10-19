import {SetList,ClearList,IsShowRightDialog} from "./types";

const state = {
    listData: [], //列表数据
    showRightDialog:false //是否显示右边的弹出框
};

const mutations = {
    //文章列表
    [SetList](state, listData){
        //判断当前的id  以确定是放在前面还是后面
        if (state.listData.length > 0 && listData.length > 0 && state.listData[0].Id > listData[0].Id) {
            state.listData = state.listData.concat(listData);
        } else {
            state.listData = listData.concat(state.listData);
        }
    },
    //清除list数据
    [ClearList](state){
        state.listData = [];
    },

    //是否显示右边的弹出框
    [IsShowRightDialog](state, yesorno){
        state.showRightDialog = yesorno;
    }
};

export default  {
    state,
    mutations
};
