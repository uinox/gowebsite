<template>
    <div class="login">
        <form method="get" action="/login" role="form" ref="form">
            <div class="form-group">
                <label for="name">Account：</label>
                <input type="text" class="form-control" id="name" pattern="[a-zA-Z0-9]{5,12}" placeholder="Enter account" name="name" value="">
            </div>
            <div class="form-group">
                <label for="pwd">Password：</label>
                <input type="password" class="form-control" id="pwd" pattern="\d{5,12}" placeholder="Password" name="password" value="">
            </div>
            <div class="checkbox">
               <p style="line-height:30px;height:30px;"><input type="checkbox" name="autoLogin" value="" style="vertical-align:middle;">&nbsp;Auto Login</p>
            </div>
            
            <button type="submit" class="btn btn-default" @click.prevent="checkInput()">Login</button>
            <button class="btn btn-default" href="/home">Back</button>
            
        </form>
            <p>没有账号，请&nbsp;<a href="/register">注册</a></p>
    </div>
</template>
<script type="text/javascript">
    import headTop from '@/components/header/head'
    import footGuide from '@/components/footer/footGuide'
    import '@/js/jquery-1.8.3.min.js'
    import '@/js/common.js'
    import Qs from 'qs'
    
    export default {

        data(){
            return {
                topics:null
            }
        },
        components:{
            headTop,
            footGuide,
        },
        created(){
            
        },
        methods:{
            checkInput(){
                var self=this;
                var formData=new FormData(self.$refs.form);
                formData.append('name',self.$refs.form.name.value);
                formData.append('password',self.$refs.form.password.value);
                formData.append('autoLogin',self.$refs.form.autoLogin.value);
                var res=null;
                var url;
                url='http://www.tmubei.com:8080/loginat';

                var ses = window.sessionStorage;
    
                $.ajax({
                    type:"post",
                    url:url,
                    async: false,
                    data:formData,  //
                    // contentType:"application/json",
                    contentType:false,
                    // dataType:'json',
                    processData:false,
                    contentType:false,
                    success:function(result){
                        console.log(result);
                        res = eval(result);   //eval将json数据解析成json对象
                    }
                });
                console.log(res);
                if(res!=null){
                    ses.setItem("token",res.Name);
                    console.log(ses.getItem("token"));
                    var addr =Qs.parse(location.search.substring(1)).addr;
                    if(addr){
                        self.$router.push('/'+addr);
                    }else{
                        self.$router.push('/home');
                    }
                }
                // console.log(localStorage.getItem("data"));
            },
            backToHome(){
                
            }
        }
    }

</script>
<style lang="css">
    .login{margin:100px auto;}
    .login label{width:80px;text-align: right;}
</style>
