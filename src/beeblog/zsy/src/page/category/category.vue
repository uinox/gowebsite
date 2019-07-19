<template>
    <div>
        <head-top></head-top>
        <div id="category">
            <div class="category">
                <h1>Category List</h1>
                <form method="post" action="/"  ref="form">
                    <div class="form-group">
                        <label>Category Name:</label>
                        <input type="text" id="name" value="" class="cateName" placeholder="Enter account" name="name" ref="name">
                    </div>
                    <input type="hidden" name="op" value="add" ref="op">
                    <button type="button" class="btn btn-default" @click="checkInput()">Add</button>
                </form>
                <table class="table table-striped" width="100%">
                    <thead>
                        <tr>
                            <td width="10%" align="center">#</td>
                            <td width="30%" align="center">Name</td>
                            <td width="30%" align="center">Nmber</td>
                            <td width="30%" align="center">Operation</td>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="cate in cates" :key="cate.Id">
                            <td align="center">{{ cate.Id}}</td>
                            <td align="center">{{ cate.Title}}</td>
                            <td align="center">{{ cate.TopicCount}}</td>
                            <td align="center"><a class="delete" @click="deleteCate(cate.Id)" >Delete</a></td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
        <foot-guide></foot-guide>
    </div>
</template>
<script type="text/javascript">
    import headTop from '@/components/header/head'
    import footGuide from '@/components/footer/footGuide'
    import '@/js/jquery-1.8.3.min.js'
    import '@/js/common.js'
    import Qs from 'qs'
    
    export default {
        inject:["reload"],
        data(){
            var data = {};
            
            var res1=[];
            var res2=[];
            var op =Qs.parse(location.search.substring(1)).op;
            
            var name=Qs.parse(location.search.substring(1)).name;
            // console.log(name);
            var url;
            // if(op && name){
            //     url='http://localhost:8080/categoryAt?op='+op+'&name='+name;
            // }else{
            //     url='http://localhost:8080/categoryAt';
            // }
            url='http://www.tmubei.com:8080/categoryAt';

            $.ajax({
                type:"post",
                url:url,
                async: false,
                data:JSON.stringify(data),  //
                contentType:"application/json",
                dataType:'json',
                processData:false,
                contentType:false,
                success:function(result){
                    var ss = result
                    var resobj = eval(ss)   //eval将json数据解析成json对象
                    
                    res2 = resobj

                }
            });
            return {
                // topics:res1,
                cates:res2
            }
        },
        components:{
            headTop,
            footGuide,
        },
        methods:{
            checkInput(){
                var name = this.$refs.name.value;
                if (name.length == 0){
                    alert("Please enter a category name");
                    return false;
                }

                var data = {
                    Name : name,
                    Op  : "add"
                };
                var url='http://www.tmubei.com:8080/categoryAt';
                $.ajax({
                    type:"post",
                    url:url,
                    async: false,
                    data:JSON.stringify(data),  //
                    contentType:"application/json",
                    dataType:'json',
                    processData:false,
                    contentType:false,
                    success:function(result){
                        var ss = result
                        var resobj = eval(ss)   //eval将json数据解析成json对象
                        
                        this.cates = resobj;
                    }
                });
                this.reload();

            },
            deleteCate(a){
                var data = {
                    Id : a+"",
                    Op : "del"
                };
                var url='http://www.tmubei.com:8080/categoryAt';
                $.ajax({
                    type:"post",
                    url:url,
                    async: false,
                    data:JSON.stringify(data),  //
                    contentType:"application/json",
                    dataType:'json',
                    processData:false,
                    contentType:false,
                    success:function(result){
                        var ss = result;
                        var resobj = eval(ss);   //eval将json数据解析成json对象
                        
                        this.cates = resobj;
                    }
                })
                this.reload();
            }
        }
    }

</script>
<style lang="css">
    .category{text-align:left;padding:0px 30px;}
    .category h1{font-size: 18px;text-align:center;}
    .delete{cursor: pointer;color: #428bca;}
</style>