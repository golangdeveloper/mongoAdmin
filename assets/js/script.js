jQuery(function($) {
  var $bodyEl = $('body'),
      $sidedrawer_dbc = $('#sidedrawer_dbc'),
      $sidedrawer_cmd = $('#sidedrawer_cmd');
  
  
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
  function showSidedrawer_cmd() {
    // show overlay
    var options = {
      onclose: function() {
        $sidedrawer_cmd
            .removeClass('active')
            .appendTo(document.body);
      }
    };

    var $overlayEl = $(mui.overlay('on', options));

    // show element
    $sidedrawer_cmd.appendTo($overlayEl);
    setTimeout(function() {
      $sidedrawer_cmd.addClass('active');
    }, 20);
  }
  
  function pcSidedrawer_dbc() {
    $bodyEl.toggleClass('hide-sidedrawer');
  }
  function pcSidedrawer_cmd() {
    $bodyEl.toggleClass('hide-sidedrawer');
  }

  // mobile
  $('.js-show-DBC').on('click', showSidedrawer_dbc);
  $('.js-show-CMD').on('click', showSidedrawer_cmd);

  // pc
  $('.js-hide-DBC').on('click', pcSidedrawer_dbc);
  $('.js-hide-CMD').on('click', pcSidedrawer_cmd);



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
