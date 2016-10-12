<!--只有一个返回按钮的页脚-->
<template>
    <footer class="footer-onlyback">
        <a @click="back()"> &lt; </a>
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