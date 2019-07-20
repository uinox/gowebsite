<template>
    <header>
        <ul>
            <li class="nav1 left" to="/home">Home</li>
            <li v-for="(item, index) in items" :key="index" class="nav left">
                <router-link tag="div" :to="{'path':item.url}">{{item.Title}}</router-link>
            </li>
            <li class="nav2 right">
                <div v-if="exit">
                    <a style="text-decoration:underline;font-size:16px;cursor:default">{{userName}}</a> <a style="font-size:12px;" @click="loginout">Exit</a>
                </div>
                <div v-else>
                    <a @click="login">Login</a> | <a href="/register">Register</a>
                </div>
            </li>
        </ul>
    </header>
</template>
<script type="text/javascript">
    import '@/js/jquery-1.8.3.min.js'
    import '@/js/common.js'
    
    export default {
        inject:['reload'],
        data(){
            return {
                isActive: 'home',
                exit: false,
                items:[
                    {Title:"Home",url:"/home"},
                    {Title:"Category",url:"/category"},
                    {Title:"Article",url:"/article"}
                ],
                addr:null,
                userName:null
            }
        },
        created(){
            
            var name = window.sessionStorage.getItem("token");
            if(name){
                this.userName=name;
                this.exit=true;
            }
        },
        mounted(){
            this.addr=document.querySelector(".isActive").innerHTML.toLocaleLowerCase();
        },
        methods: {
            loginout(){
                window.sessionStorage.removeItem("token");
                // this.common.startHacking(this, 'success', '注销成功');
                if(this.addr=='tech'){
                    this.$router.push({path:'/home'});
                }else{
                    this.userName=null;
                    this.exit=false;
                    this.reload();
                }
            },
            login(){
                if(this.addr){
                    this.$router.push({path:'/',query:{addr:this.addr}});
                }
            }
            
        }
    }

</script>
