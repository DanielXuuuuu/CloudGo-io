$(document).ready(function(){
    $.ajax({
        url: "/api/test"
    }).then(function(data){
        $('.author-id').append(data.id);
        $('.author-name').append(data.content);
    });
})