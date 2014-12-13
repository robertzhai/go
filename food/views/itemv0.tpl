{{range $key, $val := .Fooddata}}
        <div id="item" class="item">
		
            <div class="dish__wrap">
              <div class="pic">
                <a href="#">
                  <img src="/static{{$val.Pic}}" alt="{{$val.Dish}}" class="img-responsive">
                </a>
               
               
              </div>
              <div class="recipe__title">
                <a class="no_click" href="#">{{$val.Dish}}</a>
              </div>
              
            </div>
        </div> 
		
		{{end}} 
		