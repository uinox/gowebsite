<template>
    <div>
        <head-top></head-top>
        <div id="article_view">
            <div class="container">
                <div>
                    <h1>{{topic.Title}}&nbsp;&nbsp;<small>{{topic.Category}}</small></h1>
                    <h5><a v-bind:href="['/home?label='+topic.Labels.trim()]">{{topic.Labels}}</a></h5>
                    {{topic.Content}}
                    <h5>Topic File：<a v-bind:href="['http://www.tmubei.com:8080/attachment/'+topic.Attachment]">{{topic.Attachment}}</a></h5>
                </div>
                <div style="border-bottom:1px solid #ededed;">
                    <p style="border-bottom:1px solid #ededed;color:red;padding-left:20px;margin-top:20px;">Replies：</p>
                    <div v-for="reply in replies" style="padding-left:20px;">
                        <h3>{{reply.Content}}</h3>
                        <h5 style="margin-left:30px;">{{reply.Name}} <small>{{reply.Created}}</small> &nbsp;&nbsp; <small><a @click="deleteReply(topic.Id,reply.Id)">Delete</a></small></h5>
                    </div>

                </div>
                <h3>Add Reply</h3>
                <form method="post" ref="form">
                    <input type="hidden" name="tid" :value="topic.Id">
                    <div class="form-group">
                        <label>NickName</label>
                        <input type="text" name="nickname" id="" class="form-control cateName">
                    </div>
                    <div class="form-group">
                        <label>Content</label>
                        <textarea name="content" id="" cols="30" rows="10" class="form-control cateName textarea"></textarea>
                    </div>
                    <button type="submit" class="btn btn-default" @click.prevent="addReply">Submit</button>
                </form>
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
        inject:['reload'],
        data(){
            return {
                topic:null,
                replies:null
            }
        },
        created(){
            var res;
            var url;
            var id =Qs.parse(location.search.substring(1)).id;
            url='http://www.tmubei.com:8080/topicat/view/'+id;

            $.ajax({
                type:"get",
                url:url,
                async: false,
                // data:JSON.stringify(data),  //
                // contentType:"application/json",
                dataType:'json',
                processData:false,
                contentType:false,
                success:function(result){
                    res = eval(result)   //eval将json数据解析成json对象
                }
            });
            this.topic=res.Topic;
            this.replies=res.Replies;
            // console.log(this.topic)
        },
        components:{
            headTop,
            footGuide,
        },
        methods:{
            addReply(){
                var self=this;
                var res;
                var url;
                var formData = new FormData(self.$refs.form);
                formData.append('tid',self.$refs.form.tid.value);
                formData.append('nickname',self.$refs.form.nickname.value);
                formData.append('content',self.$refs.form.content.value);

                var id =Qs.parse(location.search.substring(1)).id;
                url='http://www.tmubei.com:8080/replyat/add';

                $.ajax({
                    type:"post",
                    url:url,
                    async: false,
                    data:formData,  //
                    // contentType:"application/json",
                    // dataType:'json',
                    processData:false,
                    contentType:false,
                    success:function(result){
                        res = eval(result)   //eval将json数据解析成json对象
                        self.reload();
                    }
                });
            },
            deleteReply(tid,rid){
                var self=this;
                var res;
                var url;
                var data={};
                url='http://www.tmubei.com:8080/replyat/delete?tid='+tid+'&rid='+rid;

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
                        res = eval(result);   //eval将json数据解析成json对象
                        self.reload();
                    }
                });
            }
        }
    }

</script>
