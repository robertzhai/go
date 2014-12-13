<?php
include "./config.php";
include "./Database.php";
define("DEBUG", true);

$db = new Database($db_food);
//var_dump($db);exit;
$total_pages = 100;
for($page = 1; $page < $total_pages; $page ++) {
	$str = file_get_contents('http://m.wecook.cn/recipe/cate/ename/kuaishoucai/p/' . $page . '.html');
		
	preg_match_all("/<div class=\"recipe__cover text-center\">(.*?)<\/div>/is", $str, $temp);
	
	if(isset($temp[1]) && !empty($temp[1])) {
		
		foreach($temp[1] as $item) {
		
			preg_match("/<img src=\"(.*?)\" alt=\"(.*?)\".*?/is", $item, $food);
			print_r($food);
			$img = $food[1];
			$food_name =$food[2];
		    $contents = file_get_contents($img);
		    $index = strpos($img, "?");
		    echo "index : " . $index . PHP_EOL;
		    if($index) {
		        $filename = substr($img, 0, $index);
		    }
			$pic_name =  "/images/food/" . array_pop(explode("/" , $filename));
			 
			$sql = "replace into food set pic='{$pic_name}', dish='{$food_name}'";
			$result = $db->query($sql);
			var_dump($result);
		    $pic = '.' . $pic_name;
		    file_put_contents($pic, $contents);
		    echo $pic . " download finished " . PHP_EOL;
			
		}
		
	}		
}
