<template>
    <div class="register">
        <form method="get" action="/login" ref="form">
            <div class="form-group">
                <div>
                    <label for="name">Account：</label>
                    <input type="text" @input="change" @change="change" class="form-control" id="name" placeholder="Enter account" name="name" value="">    
                </div>
                <span ref="nametip" v-if="nameError">格式不正确</span>
            </div>
            <div class="form-group">
                <div>
                    <label for="email">Email：</label>
                    <input type="email" @input="change" @change="change" class="form-control" id="email" pattern="" placeholder="Email" name="email" value="">    
                </div>
                <span ref="emailtip" v-if="emailError">格式不正确</span>
            </div>
            <div class="form-group">
                <div>
                    <label for="password">Password：</label>
                    <input type="password" @input="change" @change="change" class="form-control" id="password" pattern="\d{5,12}" placeholder="password" name="password" value="">    
                </div>
                <span ref="pwdtip" v-if="pwdError">格式不正确</span>
            </div>
            <button type="submit" class="btn btn-default" @click.prevent="checkInput()">Register</button>
            <a class="btn btn-default" style="margin-left:30px;" href="/home">Back</a>
        </form>
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
                nameError: false,
                emailError: false,
                pwdError:false
            }
        },
        components:{
            headTop,
            footGuide,
        },
        methods:{
            checkInput(){
                var self=this;

                var formData=new FormData(self.$refs.form);
                formData.append('name',self.$refs.form.name.value);
                formData.append('email',self.$refs.form.email.value);
                formData.append('password',self.$refs.form.password.value);
                var res2=[];
                var url;
                url='http://www.tmubei.com:8080/register';
    
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
                        var ss = result;
                        var resobj = eval(ss);   //eval将json数据解析成json对象
                        // console.log(resobj);
                        if(resobj && resobj[0].status=="successfull"){
                            self.$router.push('/register1');
                        }
                    }
                });
            },
            backToHome(){
                
            },
            change(event){
                var self=event.target;

                var rename = /\w{6,}/;
                var reemail = /\w+@\w+.\w+(.\w+){1,3}/;
                var repwd =/\w{6,12}/;

                self.value=self.value.replace(" ","");
                switch (self.name) {
                    case "name":
                        if(!rename.test(self.value)){
                            this.nameError = true
                        }else{
                            this.nameError = false
                        }
                        break;
                    case "email":
                        if(!reemail.test(self.value)){
                            this.emailError = true
                        }else{
                            this.emailError = false
                        }
                        break;
                    case "password":
                        if(!repwd.test(self.value)){
                            this.pwdError = true
                        }else{
                            this.pwdError = false
                        }
                        break;
                    default:
                        break;
                }
            }
        }
    }

</script>
<style lang="css">
    .register{margin:50px auto;}
    .register label{width:100px;text-align: right;}
    .register input{width: 300px;}
</style>
