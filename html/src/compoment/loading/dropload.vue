<template>
    <div class="top-loading-show" ref="loading">
        ↓ 释放加载
        <article class="overlay-loader" style="display: none">
            <div class="loader">
                <div></div>
                <div></div>
                <div></div>
                <div></div>
                <div></div>
                <div></div>
                <div></div>
            </div>
        </article>
    </div>
</template>

<script>

    let toploading;
    export default{
        mounted(){
            toploading = this.$refs.loading;
        },
        initDropLoad
    }

    //下拉刷新初始化
    function initDropLoad(domContainer, callBack) {
        domContainer.addEventListener("touchstart", function (e) {
            //判断上啦加载组件是否在视口位置
            var rect = this.getBoundingClientRect();

            //如果没有在首屏位置 则不做任何处理
            if (rect.top < 0) {
                return;
            }

            //如果没有处理其他滑动事件才继续
            if (!toploading.processing) {
                this.orignY = event.targetTouches[0].clientY;
                //添加滑动事件
                this.addEventListener("touchmove", touchMoveEvent);
                //设置为正在处理中
                toploading.processing = true;
            }

            domContainer.addEventListener("touchend", touchEnd);
        });

        //触摸移除事件
        function touchEnd() {
            this.removeEventListener("touchmove", touchMoveEvent);

            //判断下拉的距离 大于35才执行数据加载
            setLoadingState(this.lastY - this.orignY > 35 * 5 ? 2 : 3);

            //执行完毕之后移除
            this.removeEventListener("touchend", touchEnd);
        }


        //滑动过程中 改变高度
        function touchMoveEvent() {
            let currY = event.targetTouches[0].clientY;
            this.lastY = currY;

            if (currY > this.orignY)
                event.preventDefault();

            let currHeight = (currY - this.orignY) / 5;
            if (currHeight > 40) {
                setLoadingState(1);
            }
            toploading.style.height = currHeight + "px";
        }

        //设置加载中的文字
        function setLoadingState(state) {
            if (state === 0) { //默认情况
                toploading.innerHTML = "↓ 释放加载";
            } else if (state === 1) {// 已经拉倒一定阶段的情况
                toploading.innerHTML = "↑ 释放加载";
            } else if (state === 2) { // 释放正在加载的情况
                toploading.classList.add("top-loading-show-back");
                toploading.innerHTML = "加载中...";

                //回调
                callBack(()=> {
                    //设置处理完毕
                    setLoadingState(3);
                });

            } else if (state === 3) { //正在加载中的情况
                toploading.processing = false;//设置为加载完毕
                toploading.classList.remove("top-loading-show-back");//移除加载中的class
                toploading.style.height = 0;
                setLoadingState(0);
            }
        }
    }
</script>