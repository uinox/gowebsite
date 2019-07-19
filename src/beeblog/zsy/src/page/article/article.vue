<template>
    <div>
        <head-top></head-top>
        <div id="article">
            <div class="container">
                <h1>Topic List</h1>
                <!-- <a href="/article_add" class="btn btn-default">Add Topic</a> -->
                <router-link :to="{path:'/article/article_add'}" tag="a" class="btn btn-default" replace>
                    Add Topic
                </router-link>
                <div style="width:100%;overflow:auto;">
                <table class="table table-striped" style="min-width:600px;">
                    <thead>
                        <tr>
                            <td width="2%" align="center">#</td>
                            <td align="center">Title</td>
                            <td align="center">LastUpdate</td>
                            <td align="center">Views</td>
                            <td align="center">Reply</td>
                            <td align="center">LastReply</td>
                            <td align="center">Operation</td>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="topic in topics" :key="topic.Id">
                            <td align="center">{{ topic.Id}}</td>
                            <td align="center"><a :href="['/article/article_view?id='+topic.Id]">{{ topic.Title}}</a></td><!-- href="/topic/view/{{.Uid}}" -->
                            <td align="center">{{ topic.Updated}}</td>
                            <td align="center">{{ topic.Views}}</td>
                            <td align="center">{{ topic.ReplyCount}}</td>
                            <td align="center">{{ topic.ReplyTime}}</td>
                            <td align="center">
                                <router-link :to="{path:'/article/article_edit',query:{tid:topic.Id}}" tag="a">
                                    Edit
                                </router-link>
                                <!-- <a href="#">Edit</a> -->
                                <!-- /topic/modify/?tid={{topic.Id}} -->
                                <a @click="deleteTopic(topic.Id)" style="padding-left:20px;cursor:pointer;">Delete</a><!-- /topic/delete?tid={{topic.Id}} -->
                            </td>
                        </tr>
                    </tbody>
                </table>
                </div>
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
    
    export default {
        inject:["reload"],
        data(){
            
            return {
                topics:null
            }
        },
        created(){
            var data = {};
            var res2=[];
            var url;
            url='http://www.tmubei.com:8080/topicat';

            $.ajax({
                type:"get",
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
                    res2 = resobj.reverse();
                }
            });
            this.topics=res2;
        },
        components:{
            headTop,
            footGuide,
        },
        methods:{
            deleteTopic(id){
                var self=this;
                var data={tid:id};
                var url="http://www.tmubei.com:8080/topicat/delete?tid="+id;
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
                    self.reload();
                }
            });
            }
        }
    }

</script>