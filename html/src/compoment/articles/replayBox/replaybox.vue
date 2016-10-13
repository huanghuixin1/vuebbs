<!--回复帖子的框-->
<template>
    <article class="replayBox">
        <textarea placeholder="想说啥子写进去">

        </textarea>

        <div>
            <i>取消</i>
            <a @click="publishReplay()">发布</a>
        </div>
    </article>
</template>
<script>
    export default{
        methods:{
            publishReplay(){
                alert("发布");
            }
        }
    }
</script>
