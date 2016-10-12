import * as types from "./types";
import * as userService from "../../../service/userinfosService";
import * as articleService from "../../../service/articlesService";
import * as categoriesService from "../../../service/categoriesService";

/**
 * 设置标题
 * @param commit
 * @param state
 * @param cid 板块id
 * @param title 标题
 * @constructor
 */
export const SetTitle = ({commit, state}, {cid, title}) => {
    //拿到cid去获取该模块的名称
    if (cid) {
        //去请求数据
        categoriesService.getNameById({
            cid, success: rep=> {
                commit(types.SetTitle, rep.Data);
            }
        });
    } else if (title) {
        commit(types.SetTitle, title);
    }
};

/**
 * 设置会员总数
 * @param total 当<=0或者不传的时候 则去请求获取总数  否则直接设置
 */
export const SetTotalMembers = ({commit, state}, total) => {
    if (!total || total <= 0) {
        userService.getCount(data=> {
            if (data.Status === 0)
                commit(types.SetTotalMembers, data.Data);
        });
    } else {
        commit(types.SetTotalMembers, total);
    }
};

/**
 * 设置总的帖子数量
 * @param commit
 * @param state
 * @param total 总帖子数量
 * @param cid 板块cid(当该参数不为空的时候, 则请求服务器进行获取)
 * @constructor
 */
export const SetTotalArticles = ({commit, state}, options) => {
    let total = options.total, cid = options.cid, success = options.success, errback = options.errback;
    if (cid >= 0) {
        articleService.getCount(cid, data=> {
            if (data.Status === 0) {
                commit(types.SetTotalArticles, data.Data);
                success && success(data);
            } else {
                err && err(data);
            }
        }, err => {
            errback && errback(err);
        });
    } else {
        commit(types.SetTotalArticles, total);
    }
};

//是否显示头部的详情信息
export const SetIsShowHeaderDetail = ({commit, state}, {isShow}) => {
    commit(types.SetIsShowDetail, isShow);
};