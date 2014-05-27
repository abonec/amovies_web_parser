$(function(){
  console.log('hello');
  $('.down_link').click(function(e){
    var link = $(this).data('link');
    var episode = $(this).data('episode');
    var prefix = $(this).data('prefix');
    $.post('/download',
           {link: link, episode: episode, prefix: prefix}, 
           function(data){
             alert(data);
           });
    e.preventDefault();
  });
  $('.remove_download').click(function(e){
    var id = $(this).data('id');
    $.ajax({url: '/remove/'+id, type: 'DELETE'});
    e.preventDefault();
    location.reload();
  });
});
