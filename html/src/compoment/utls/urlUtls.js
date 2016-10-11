export const queryStrings = (function (name) {
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
}());

/**
 * @param name 获取的key
 * @param hasHashUrl url上是否有hash
 */
export function getQueryString(name) {
    if (name in queryStrings)
        return decodeURI(queryStrings[name]);
    else
        return null;
}