import * as typeUtls from "./typeUtls";

/**
 * 获取当前时区与utc时间的间隔
 * @returns {number} 返回单位:秒
 */
export function getTimezonOffset() {
    return new Date().getTimezoneOffset() * 60;
}

/**
 * 获取当前的时间戳 单位:秒
 */
export function getUTCTimetamp() {
    return parseInt(Date.now() / 1000) + getTimezonOffset();
}

/**
 * 日期转为字符串
 * @param time 时间戳获取是Date对象
 * @param isUTC 是否是UTC时间
 */
export function time2String(time, isUTC) {
    //如果time是日期类型
    if (typeUtls.isDate(time)) {
        // 转换为时间戳格式 并将单位转为秒
        time = time.getTime() / 1000;
    }

    //如果是utc时间 加上本地时区的差值
    if (isUTC) {
        time += getTimezonOffset();
    }

    //最后再转为日期格式
    let finalDate = new Date(time * 1000);

    //返回字符串
    return finalDate.getFullYear() + "-" + (finalDate.getMonth() + 1) + "-" + finalDate.getDate()
        + " " + finalDate.getHours() + ":" + finalDate.getMinutes();
}

/**
 * 获取时间间隔的描述字符串
 * @param time 参数必须是UTC时间戳
 */
export function getTimespanDesc(time) {
    let timespan = getUTCTimetamp() - time;

    //1分钟内
    if (timespan < 60)
        return parseInt(timespan) + " 秒前";

    timespan = timespan / 60;
    //1小时内
    if (timespan < 60)
        return parseInt(timespan) + " 分钟前";

    timespan = timespan / 60;
    //一天内
    if (timespan < 24) {
        return parseInt(timespan) + " 小时前";
    } else if (parseInt(timespan) < 48) { //两天内
        return "昨天";
    }

    //如果超过了两天 则直接显示创建时间
    return time2String(time, true);
}

