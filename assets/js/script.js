jQuery(function($) {
  var $bodyEl = $('body'),
      $sidedrawer_dbc = $('#sidedrawer_dbc');
  
  
  // ==========================================================================
  // Toggle Sidedrawer
  // ==========================================================================
  function showSidedrawer_dbc() {
    // show overlay
    var options = {
      onclose: function() {
        $sidedrawer_dbc
          .removeClass('active')
          .appendTo(document.body);
      }
    };
    
    var $overlayEl = $(mui.overlay('on', options));
    
    // show element
    $sidedrawer_dbc.appendTo($overlayEl);
    setTimeout(function() {
      $sidedrawer_dbc.addClass('active');
    }, 20);
  }
  
  
  function hideSidedrawer_dbc() {
    $bodyEl.toggleClass('hide-sidedrawer');
  }
  
  
  $('.js-show-sidedrawer').on('click', showSidedrawer_dbc);
  $('.js-hide-sidedrawer').on('click', hideSidedrawer_dbc);
  
  
  // ==========================================================================
  // Animate menu
  // ==========================================================================
  var $titleEls = $('strong', $sidedrawer_dbc);
  
  $titleEls
    .next()
    .hide();
  
  $titleEls.on('click', function() {
    $(this).next().slideToggle(200);
  });
});
