<template>
    <div class="login">
        <form role="form" ref="form">
            <div class="form-group">
                <label for="name">Account：</label>
                <input type="text" class="form-control" id="name" @keyup="checkInput('name')" @keyup.delete="checkInput('name')" placeholder="6-12位字母或数字" name="name" value="">
            </div>
            <div class="form-group">
                <label for="pwd">Password：</label>
                <input type="password" class="form-control" id="pwd" @keyup="checkInput('pwd')" @keyup.delete="checkInput('pwd')" placeholder="6-12位字母或数字" name="password" value="">
            </div>
            <div class="checkbox">
               <p style="line-height:30px;height:30px;">
                 <label>
                   <input type="checkbox" name="autoLogin" value="" style="vertical-align:middle;">
                 </label>&nbsp;Auto Login
               </p>
            </div>
            
            <button type="submit" class="btn btn-default" @click.prevent="login()">Login</button>
            <a class="btn btn-default" href="/register">Register</a>

        </form>
        <p ref="errShow" class="errShow" v-if="errShow">{{errMsg}}</p>

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
                topics:null,
                errShow:false,
                errMsg:""
            }
        },
        components:{
            headTop,
            footGuide,
        },
        created(){
            
        },
        methods:{
            login(){
                if(!(this.checkInput("name")&&this.checkInput("pwd"))){
                  return
                }
                var self=this;
                var formData=new FormData(self.$refs.form);
                var name=String(self.$refs.form.name.value);
                var password=self.$refs.form.password.value;
                var autoLogin=self.$refs.form.autoLogin.value;
                formData.append('name',name);
                formData.append('password',password);
                formData.append('autoLogin',autoLogin);
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
                    dataType:'json',
                    processData:false,
                    contentType:false,
                    success:function(result){
                        res = eval(result);   //eval将json数据解析成json对象

                    }
                });
                if(res!=null){
                  if(res.Status=="404"){
                    this.errShow=true;
                    this.errMsg="此用户名尚未注册，请注册";
                    return;
                  }
                  if(res.Status=="403"){
                    this.errShow=true;
                    this.errMsg="密码错误";
                    return;
                  }
                  if(res.Status=="200"){
                    ses.setItem("token",res.Name);
                    var addr =Qs.parse(location.search.substring(1)).addr;
                    if(addr){
                        self.$router.push('/'+addr);
                    }else{
                        self.$router.push('/home');
                    }
                  }

                }
            },
            checkInput(item){
              var self=this;
              if(item=="name"){
                var nameReg=/^\w{5,12}$/;
                var name=self.$refs.form.name;
                if(name.value.length==0){
                  this.errShow=true;
                  this.errMsg="姓名不能为空";
                  return false;
                }
                if(!nameReg.test(name.value)){
                  this.errShow=true;
                  this.errMsg="请输入6-12位数字或字母";
                  return false;

                }
                this.errShow=false;
                this.errMsg="";
              }

              if(item=="pwd"){
                var pwdReg=/^[a-zA-Z0-9]{5,12}$/;
                var pwd=self.$refs.form.password;
                if(pwd.value.length===0){
                  this.errShow=true;
                  this.errMsg="密码不能为空";
                  return false;
                }
                if(!pwdReg.test(pwd.value)){
                  this.errShow=true;
                  this.errMsg="请输入6-12位数字或字母";
                  return false;
                }
                this.errShow=false;
                this.errMsg="";
              }
              return true;
            }
        }
    }

</script>
<style lang="css">

</style>
