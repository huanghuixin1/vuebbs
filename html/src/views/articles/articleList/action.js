"use strict";
import * as types from "./types";
import * as typeUtls from "../../../compoment/utls/typeUtls";
import  * as articleService from "../../../service/articlesService";

/**
 * 设置列表数据
 * @param cidOrData 模块id 或是列表数据
 * @param minId 取大于该id的数据
 * @param maxId 取小于该id的数据
 * @param callback 成功回调
 * @param errCallback 失败回调
 */
export const setArticleList = ({commit, state}, {cidOrData, minId, maxId, callback, errCallback}) => {
    //如果是数组
    if (typeUtls.isArray(cidOrData)) {
        commit(types.SetList, cidOrData);
    } else {//否则去服务器
        articleService.getListByCid(cidOrData, minId, maxId, data=> {
            if (data.Status === 0) {
                commit(types.SetList, data.Data);
                callback && callback();
            } else {
                errCallback ? errCallback() : alert("获取数据失败,请刷新页面重试!");
            }
        });
    }
};


//清除列表数据
export const ClearList = ({commit})=>{
    commit(types.ClearList);
} ;

export const IsShowRightDialog = ({commit}, yesorno)=>{
    commit(types.IsShowRightDialog, yesorno);
};