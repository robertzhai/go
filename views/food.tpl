
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0, minimal-ui" />
  <title>快手菜</title>
  <link rel="stylesheet" href="http://cdn.staticfile.org/twitter-bootstrap/3.2.0/css/bootstrap.min.css">
  <link rel="stylesheet" href="http://c1.wecook.cn/Public/Mobile/components/header/header.css?v=2014092601">
  <link rel="stylesheet" href="http://c1.wecook.cn/Public/Mobile/components/gotop/gotop.css?v=2014102001">
<link rel="stylesheet" href="http://c1.wecook.cn/Public/Mobile/components/download-modal/download-modal.css?v=2014102001">
  <link rel="stylesheet" href="http://c1.wecook.cn/Public/Mobile/css/recipe/featured.css?v=2014102001">

</head>
<body>
  <div class="container" id="header">
    <div class="row">
      <div class="col-xs-2">
        <div class="back">
          <a href="javascript:void(0)">返回</a>
        </div>
      </div>
      <div class="col-xs-8">
        <div class="text-center title-wrap">
          <span class="title icon">快手菜</span>
        </div>
      </div>
      <div class="col-xs-2">
        <div class="home">
          <a href="/">首页</a>
        </div>
      </div>
    </div>
  </div>
  <div class="content has_header">
        
    <div id="recipes" class="container">
      <div class="row">
	
       <link rel="stylesheet" href="http://c1.wecook.cn/Public/Mobile/components/download-modal/download-modal.css?v=2014102001">
  <link rel="stylesheet" href="http://c1.wecook.cn/Public/Mobile/css/recipe/featured.css?v=2014102001"> 

      {{range $key, $val := .Fooddata}}
        <div class="col-xs-6 col-md-4 col-lg-3 recipe_item">
		
            <div class="recipe__wrap">
              <div class="recipe__cover text-center">
                <a href="/caipu-4377.html">
                  <img src="/static{{$val.Pic}}" alt="{{$val.Dish}}" class="img-responsive">
                </a>
               
              </div>
              <div class="recipe__title">
                <a class="no_click" href="#">{{$val.Dish}}</a>
              </div>
              
            </div>
        </div> 
		
		{{end}} 	

    <div id="loadmore" style="display:none;">
      <div class="row">
        <a class="col-xs-12 text-center loadstate" href="/recipe/featured/p/2.html">
          查看更多
        </a>
      </div>
      <div class="hidden" id="loadtmp"></div>
    </div>

    <div class="download" id="download-modal">
      <div class="download-wrap">
        <div class="download-close">&times;</div>
        <a class="download-btn" href="http://wecook.cn/download?src=recipe"></a>
      </div>
    </div>
    <div class="page hidden">
        <div> <a class="prev" href="/recipe/cate/ename/kuaishoucai/p/1.html"><<</a> <a class="num" href="/recipe/cate/ename/kuaishoucai/p/1.html">1</a><span class="current">2</span><a class="num" href="/recipe/cate/ename/kuaishoucai/p/3.html">3</a><a class="num" href="/recipe/cate/ename/kuaishoucai/p/4.html">4</a> <a class="next" href="/recipe/cate/ename/kuaishoucai/p/3.html">>></a> <a class="end" href="/recipe/cate/ename/kuaishoucai/p/207.html">207</a> </div>    </div>
  </div>
  <div id="gotop"></div>
  <script src="http://cdn.staticfile.org/jquery/2.1.1/jquery.min.js"></script>
  <script src="http://c1.wecook.cn/Public/Mobile/components/header/header.js?v=2014092601"></script>
  <script src="http://c1.wecook.cn/Public/Mobile/components/gotop/gotop.js?v=2014102001"></script>
  <script src="http://c1.wecook.cn/Public/Mobile/components/download-modal/download-modal.js?v=2014102001"></script>
  <script src="http://c1.wecook.cn/Public/Mobile/plugins/imagesloaded/imagesloaded.pkgd.min.js"></script>
  <script src="http://c1.wecook.cn/Public/Mobile/components/wxshare/wxshare.js?v=2014102201"></script>
  <script>
    jQuery(function(){
      wecook.header.setIcon(false).setBack(true).setHome(true);
      $('.no_click').on('click', false);
      $('.recipe_action').on('click', function(){
        wecook.downloadModal.show();
        return false;
      });

      imagesLoaded('#recipes', function(){
        $('.img-loading').hide();
        $('.recipe_item').fadeIn();
        $('#loadmore').show();
      });

      var lastPage, nextPage;
      if($('.page a.next').length){
        nextPage = $('.page a.next').attr('href');
      }else{
        $('.loadstate').text('已经全部加载');
      }
      $('.loadstate').on('click', function(){
        var self = $(this);
        if($(this).hasClass('loading')){
          return false;
        }
        if(lastPage === nextPage){
          return false;
        }

        $(this).addClass('loading')

        $('#loadtmp').load(nextPage + ' #recipes .row .recipe_item', function(responseText, textStatus, XMLHttpRequest){

          imagesLoaded('#loadtmp', function(){
            self.removeClass('loading');
            lastPage = nextPage;
            var match = responseText.match(/<a class="next" href="([^\"]*)">>><\/a>/);
            if(match){
              $('.loadstate').text('查看更多');
              nextPage = match[1];
            }else{
              $('.loadstate').text('已经全部加载');
            }
            $('#loadtmp').children().fadeOut(0).appendTo('#recipes .row').fadeIn();
            $('.no_click').on('click', false);
            $('.recipe_action').on('click', function(){
              wecook.downloadModal.show();
              return false;
            });
          });
        });

        return false;
      });
    });
  </script>
  <script>
    (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
    (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
    m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
    })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
    ga('create', 'UA-47304460-2', 'm.wecook.cn');
    ga('send', 'pageview');
  </script>
</body>
</html>