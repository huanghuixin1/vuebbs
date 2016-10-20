function initQueryString() {
    let search = window.location.hash.split("?")[1];
    if (!search) {
        return [];
    }
    let searchs = search.split("&");
    let ret = [];
    for (let i = 0; i < searchs.length; i++) {
        let curr = searchs[i].split("=");
        ret[curr[0]] = curr[1];
    }

    return ret;
}

export let queryStrings = initQueryString();

/**
 * @param name 获取的key
 */
export function getQueryString(name, isRefresh = false) {
    if (isRefresh) {
        queryStrings = initQueryString();
    }
    if (name in queryStrings)
        return decodeURI(queryStrings[name]);
    else
        return null;
}

/**
 * 设置url参数, 注: 该方法会重置url
 * @param key
 * @param val
 */
export function setQueryString(key, val) {
    //判断url上是否有该参数
    let orignVal = this.getQueryString(key, true);

    let finalParam = key + "=" + val;

    let url = window.location.href;
    //如果找到了
    if (orignVal) {
        url = url.replace(key + "=" + orignVal, finalParam);
    } else { //如果没找到
        // 这时候判断url是否已经有了参数
        if (url.indexOf("?") > 0) {
            //有 则往后追加
            url += "&" + finalParam;
        } else {
            url += "?" + finalParam;
        }
    }

    //最后跳转
    window.location = url;
}