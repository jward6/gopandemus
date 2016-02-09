(function(){

  angular.module('pandemus.services.backend', [
    'ngResource'
  ])

  //.factory('NewGame', ['$resource',
  //  function($resource){
  //      return $resource('api/new-game', {}, {
  //          query: {method: 'GET', isArray: false}
  //      });
  //  }])

    .factory('BoardState', ['$resource',
      function($resource){
        return $resource('api/board-state', {}, {
          newGame: {
            url: '/api/new-game',
            method: 'GET',
            isArray: false
          }
        });
      }]);


    //.factory("Command", ['$resource',
    //  function($resource){
    //    return $resource('api/command', {})
    //}]);

})();
