import {SetTitle, SetTotalMembers, SetTotalArticles, SetIsShowDetail} from "./types";

const state = {
    title: "热门推荐",
    totalArticles: 999999, //当前模块的总数量
    totalMembers: 99999, //总会员数量
    isShowHeaderDetail: true // 是否显示头部的详情信息(总会员 总帖子数量)
};

const mutations = {
    //文章列表
    [SetTitle](state, title){
        state.title = title;
    },

    //设置总会员数量
    [SetTotalMembers](state, total){
        state.totalMembers = parseInt(total);
    },

    //设置[当前板块]总帖子数量
    [SetTotalArticles](state, total){
        state.totalArticles = parseInt(total);
    },

    //是否显示头部的详情信息
    [SetIsShowDetail](state, isShow){
        state.isShowHeaderDetail = isShow;
    }
};

export default  {
    state,
    mutations
};