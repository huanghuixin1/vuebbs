import config from "../compoment/utls/config";
import * as ajaxUtls from "./ajaxUtls";

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
    ajaxUtls.sendGet(url, callback, errCallback);
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
    ajaxUtls.sendGet(url, callback, errCallback);
}

/**
 * 获取帖子详情
 * @param id
 * @param success 成功回调
 * @param err 失败会滴啊
 */
export function getDetail({id, success, err}) {
    let url = urlPre + "/getDetail?id=" + id;
    ajaxUtls.sendGet(url, success, err);
}