'use strict';

(function(){

  angular.module('pandemus', [
    'ngRoute',
    'pandemus.console',
    'pandemus.board-state',
    'pandemus.services.backend'
  ])

  .config(['$routeProvider', function($routeProvider){
    $routeProvider.otherwise({redirectTo: '/console'});
  }]);

})();
