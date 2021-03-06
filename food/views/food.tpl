<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0, minimal-ui" />
  <title>快手菜</title>
<link rel="stylesheet" href="/static/css/reset.css" />
<style>
.box {
	width:1200px;
	margin:0 auto;
}
.box_img {
	width:280px;
}
.dish__wrap {
	
}
.item{
		width:290px;
		float:left;
		margin-right:5px;
	
}

.more {
	margin:0 auto;
	clear:both;
	width:1200px;
	height:60px;
	text-align:center;
	display:block;
	border:1px solid silver;
	margin-top:100px;
	
}
.more a {
	display: inline-block;
margin-top: 2%;
}
a {
	text-decoration: none
}
</style>

</head>


<body>
<div class="box">
{{range $key, $val := .Fooddata}}
        <div id="item" class="item">
		
            <div class="dish__wrap">
              <div class="pic">
                <a href="#">
                  <img class="box_img" src="/static{{$val.Pic}}" alt="{{$val.Dish}}" class="img-responsive">
                </a>
               
               
              </div>
              <div class="recipe__title">
                <a class="no_click" href="#">{{$val.Dish}}</a>
              </div>
              
            </div>
        </div> 
		
		{{end}} 
		
	
		
</div>
<div class="more">
	 <a class="more_a" href="#" onclick="more()">
          查看更多
        </a>
	
	</div>
<script src="/static/js/jquery-1.7.1.min.js"></script>

<script>

var page = 1;

function more() {
 page = page + 1;
 var 	html =$.ajax({url:"/food/ajax?page=" + page ,async:false});
  $(".box").append(html.responseText);
}

</script>


</body>
</html>	