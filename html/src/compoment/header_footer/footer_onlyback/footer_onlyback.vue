<!--只有一个返回按钮的页脚-->
<template>
    <footer class="footer-common">
        <a @click="back()" class="footer-back"> &lt; </a>
    </footer>
</template>

<script>
    export default{
        methods: {
            back:function () {
                window.history.go(-1);
            }
        }
    }
</script>