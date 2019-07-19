<template>
    <div>
        <head-top></head-top>
        <div id="article_add">
            <div class="container">
                <h1>Add Topic</h1>
                <form ref="form" enctype="multipart/form-data">
                    <div class="form-group">
                        <label>Topic Title:</label>
                        <input type="text" ref="title" name="title" class="cateName" id="" value="">
                    </div>
                    <div class="form-group">
                        <label>Topic Category</label>
                        <input type="text" ref="category" name="category" id="" class="cateName" value="">
                    </div>
                    <div class="form-group">
                        <label>Topic Label:</label>
                        <input type="text" ref="label" name="label" class="cateName" id="" value="">
                    </div>
                    <div class="form-group">
                        <label>Topic Content:</label>
                        <textarea ref="content" name="content" id="" cols="30" rows="10" class="cateName textarea" value=""></textarea>
                    </div>
                    <div class="form-group">
                        <label>Topic File:</label>
                        <input type="file" ref="attachment" name="attachment" class="cateName" id="" style="height:auto;">
                    </div>
                    <input type="submit" @click.prevent="addTopic" class="btn btn-default" value="Add Topic" />
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
    
    export default {
        name:'article_add',
        data(){
            return {

            }
        },
        components:{
            headTop,
            footGuide,
        },
        methods:{
            addTopic(){
                var self=this;
                var formData = new FormData(self.$refs.form);
                formData.append('file',self.$refs.attachment.files[0]);
                formData.append('title',self.$refs.title.value);
                formData.append('category',self.$refs.category.value);
                formData.append('content',self.$refs.content.value);
                formData.append('label',self.$refs.label.value);

                var res2=[];
                var url;
                url='http://www.tmubei.com:8080/topicat/add';
    
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
                        if(resobj && resobj[0].status=="successfull"){
                            self.$router.push('../article');
                        }
                    }
                });
            }
        }
    }

</script>