import config from "../compoment/utls/config";
import ajax from "ajax";

//url前缀
const urlPre = config.host + "/articlesApi";

/**
 * 获取对应板块帖子总数
 * @param cid 板块id
 * @param callback 成功回调
 * @param errCallback 失败回调
 */
export function getCount(cid, callback, errCallback) {
    let url = urlPre + "/count?cid=" + cid;
    ajax().get(url).then((response, xhr)=> {
        callback && callback(response);
    }).catch(function (response, xhr) {
        errCallback && errCallback(response);
    });
}

/**
 * 通过cid获取文章列表
 * @param cid 板块id
 * @param minId 取大于该id的数据
 * @param maxId 取小于该id的数据
 * @param callback 成功回调
 * @param errCallback 失败回调
 */
export function getListByCid(cid, minId, maxId, callback, errCallback) {
    let url = urlPre + "/getbycid?cid=" + cid + "&minId=" + minId + "&maxId=" + maxId;
    ajax().get(url).then((rep)=> {
        callback && callback(rep);
    }).catch(rep=> {
        errCallback && errCallback(rep);
    });
}