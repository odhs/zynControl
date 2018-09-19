package main
const scripts = 
`<script>
/* RENDER section */
window.addEventListener("hashchange", function(){
        /* On every hash change the render function is called with the new hash.
        This is how the navigation of our app happens.*/
        render(decodeURI(window.location.hash));
      }, false);

function render(url) {
  var temp = url.split('/')[0];

  var map = {
    /* The Homepage */
    '': function () {
        var x = document.querySelectorAll('section.page');
        var i;
        for (i = 0; i < x.length; i++) {
            x[i].classList.add('right');
        }
     },
    '#index': function() {
      document.querySelector('section.page').classList.add('right');
    },
    '#bank': function () {
      var id = url.split('#bank/')[1].trim();
      document.querySelector('.page-bank'+id).classList.remove('right');
    }
  };
  /* Execute the needed function depending on the url keyword (stored in temp). */
  if(map[temp]){
    map[temp]();
    return;
  }
  /* If the keyword isn't listed in the above - render the error page. */
  else {
    renderErrorPage();
  }  
}
</script>`