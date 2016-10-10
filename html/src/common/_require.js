/**
 * 导入url
 * @param url-是相对于网站根目录的绝对路径
 */
export function requireCss(url) {
    let link = document.createElement("link");
    link.href = url;
    link.rel = "stylesheet";
    document.head.appendChild(link);
}