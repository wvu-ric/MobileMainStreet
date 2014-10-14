/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */


 // set global image counter
    var counter = 0;
    var current_img_array;

    // start rotating in random images
    function fadeImage(img, count, len){
        // timing based on alt_img_array length
        var divName = "#img"+count;
        var min = 2000, max = 10000;
        var time = Math.floor(Math.random() * (max - min) + min);

        setTimeout(function() { 
            var $image = $('<img src=img/filmstrip/'+img+'.jpg style="opacity:0;" class="image" />');
            $(divName).append($image); // Add second image to element
            $($(divName).find( "img:eq(0)")).fadeOut(2000);
            $($(divName).find( "img:eq(1)")).fadeTo( "slow" , 1, function() {
                   $($(divName).find( "img:eq(0)")).remove();
                    counter=counter+1;
                    if(counter >= len){
                        // call another rotation
                        nextRotation();
                    }
            });
           
       }, time);
    }
    function startRotator(array) {
        current_img_array = array;
        var tmp_array = array.slice(0);// clone array
        var len = tmp_array.length;
      
        for ( var i = 0; i < len; i++ ) {
          
             if(current_img_array == img_array){
                  var randomElementIndex = i;
                  var image = tmp_array[i];
                 //var randomElementIndex = Math.floor( Math.random() * tmp_array.length );
            } else {
                var randomElementIndex = Math.floor( Math.random() * tmp_array.length );
                var image = tmp_array.splice(randomElementIndex, 1);
            }
           
           
            fadeImage(image, i+1, len);
        }
    }
    function nextRotation(){
        counter = 0;
        if(current_img_array == img_array){
            startRotator(alt_img_array);
        } else {
            startRotator(img_array);
        }
    }