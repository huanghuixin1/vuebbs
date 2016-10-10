//Boolean Number String Function Array Date RegExp Object Error
//判断变量类型的工具
export function getType(obj) {
    return Object.prototype.toString.call(obj);
}

export const types = {
    bool: "[object Boolean]",
    num: "[object Number]",
    str: "[object String]",
    fn: "[object Function]",
    arr: "[object Array]",
    date: "[object Date]",
    reg: "[object RegExp]",
    obj: "[object Object]",
    err: "[object Error]",
};

//是否是数组
export function isArray(obj) {
    return getType(obj) === types.arr;
}

//是否是数值
export function isNumber(obj) {
    return getType(obj) === types.num;
}

//是否是日期
export function isDate(obj) {
    return getType(obj) === types.date;
}