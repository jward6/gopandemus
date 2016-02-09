'use strict';

(function(){
  angular.module('pandemus.board-state', [])

  .directive('boardState', function(){
    return {
      templateUrl: '/Components/Board/board-state.html',
      scope: {
        state: '=state'
      }
    };
  });

})();
