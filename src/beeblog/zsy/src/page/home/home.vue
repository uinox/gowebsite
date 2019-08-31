<template>
    <div>
        <head-top></head-top>
        <div id="main">
            <div style="overflow:hidden;">
                <div style="float:left;width:75%;min-height:100px;">
                    <ul>
                        <li class="topic" v-for="topic in topics" :key="topic.value">
                            <h1><a :href="'/article/article_view?id='+topic.Id">{{topic.Title}}</a></h1>
                            <h6 class="text-muted">This blog published by {{topic.Created}}, The total number of {{topic.Views}},{{topic.ReplyCount}} comments</h6>
                            <div>
                                {{topic.Content}}
                            </div>
                        </li>
                    </ul>
                    <p v-if="isNull" id="tip" ref="tip" style="margin-top:30px;color:red;">没有相关信息！</p>
                </div>
                <div class="topicCate" style="float:left;width:25%;text-align:left;">
                    <h3>Category</h3>
                    <ol>
                        <li v-for="cate in cates" :key="cate.value">
                            <a :href="'/home?cate='+cate.Title">{{cate.Title}}</a>
                        </li>
                    </ol>
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
    import Qs from 'qs'
    
    export default {
        data(){
            return {
                topics:null,
                cates:null,
                isNull:false
            }
        },
        created(){
            var res;
            var data={};
            var cate =Qs.parse(location.search.substring(1)).cate;
            var label = Qs.parse(location.search.substring(1)).label;
            var url;
            if(cate){
                url='http://www.tmubei.com:8080/homeat?cate='+cate;
            }else if(label){
                url='http://www.tmubei.com:8080/homeat?label='+label;
            }else{
                url='http://www.tmubei.com:8080/homeat';
            }
            $.ajax({
                type:"GET",
                url:url,
                async: false,
                data:JSON.stringify(data),  
                contentType:"application/json",
                dataType:'json',
                processData:false,
                contentType:false,
                success:function(result){
                    var resobj = eval(result)   //eval将json数据解析成json对象
                    res=resobj;
                }
            })
            this.topics = res.Topics;            
            this.cates = res.Categories;
        },
        mounted(){
            this.isNull=this.topics.length>0?false:true;
        },
        components:{
            headTop,
            footGuide,
        },
        methods:{
            
        }
    }

</script>
<style lang="css">
    ol{list-style-type: none;margin-left:-10px;text-align:center;}
    ol li {height:20px;line-height: 20px;}
    
    .topic{text-align: left;padding-left:10px;border-bottom: 1px solid #eee;padding-bottom: 10px;list-style-type: none;}
    .topic h1 a{font-size: 36px;font-weight: 500;}
    .topic h6{margin: 10px 0;color: #777;}

    .topicCate h3{font-size: 16px;text-align:center;}
    .topicCate a{font-size: 14px;}
</style>
