$(function(){
  $('.down_link').click(function(e){
    var link = $(this).data('link');
    var episode = $(this).data('episode');
    var prefix = $(this).data('prefix');
    $.post('/download',
           {link: link, episode: episode, prefix: prefix}, 
           function(){console.log('ok')});
    e.preventDefault();
  });
})
