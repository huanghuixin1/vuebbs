<template>
    <div>
        <h1>测试</h1>

        <input type="text" v-model="a"/>

        <button @click="get()">获取</button>
    </div>
</template>

<script>
    export default{
        methods:{
            get(){
                alert(this.a);
            }
        }
    }
</script>