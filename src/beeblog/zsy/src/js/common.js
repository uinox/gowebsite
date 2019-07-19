function res1(){
    var name = 'any';
    var data = {
        "Name" : name,
        "Age"  : 20
    };
    $.ajax({
        type:"POST",
        url:"http://localhost:8080/topicat",
        async: true,
        data:JSON.stringify(data),
        contentType:"application/json",
        dataType:'json',
        processData:false,
        contentType:false,
        success:function(result){
            // console.log(result);
            var res1= result;
            // console.log(res1);
            // document.getElementById("interest").innerHTML=res.interest;
            // document.getElementById("major").innerHTML=res.major;
            // var res1 = JSON.stringify(res);
            // console.log(res1);
            return res1
        }
    })
}