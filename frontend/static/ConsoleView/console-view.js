'use strict';

(function(){

  angular.module('pandemus.console', [
    'ngRoute'
  ])

  .config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/console', {
      templateUrl: '/ConsoleView/console-view.html',
      controller: 'ConsoleCtrl'
    });
  }])

  .controller('ConsoleCtrl', ['$scope', 'BoardState', function($scope, BoardState) {
      $scope.board = {
        data: BoardState.newGame(),
        committed: true,
        pending: [],
      };

      $scope.historyIdx = 0;
      $scope.history = [];
      $scope.commandText = 'Enter commands here...';

      var evalCommand = function(cmd){
        var factory = function(cmd) {
          var d = cmd.split(' ');
          var name = d[0];
          var args = d.slice(1, d.length);

          if(name == 'move'){
            return function(){
                var player = $scope.board.data.Players[args[0]];
                var city = "";
                for(var i=1 ; i<args.length ; i++)
                {
                  city = city + args[i] + " ";
                }
                player.Location = city.trim();
            };
          }
          else if(name == 'draw'){
            return function(){
              console.log("Draw cards command with args: ", args);
              var player = $scope.board.data.Players[args[0]];
              var color = args[1];
              var city = "";
              for(var i=2 ; i<args.length ; i++)
              {
                city = city + args[i] + " ";
              }
              player.Hand.push({
                Name: city,
                Color: color,
              });
            }
          }
          else if(name == 'discard'){
            return function(){
              var player = $scope.board.data.Players[args[0]];
              var cardName = args[1];
              var foundIx = -1;
              for(var i = 0 ; i < player.Hand.length ; i++)
              {
                if(player.Hand[i] == cardName){
                  foundIx = i;
                }
              }
              player.Hand.splice(foundIx, 1);

            };
          }
          else if(name == 'treat'){
            return function(){
              console.log("treat city command with args: ", args);
            };
          }
          else if(name == 'infect'){
            return function(){
              console.log("infect city command with args: ", args);
            };
          }
          else{
            return function(){
              console.log('Invalid Command name or format.');
            }
          }
        };

        var handler = factory(cmd);
        return handler();

      };

      $scope.processCommand = function(evt){
        evt = evt || window.event;
        var charCode = evt.keycode || evt.which;
        if(charCode == 13){ // Enter
          var cmd = $scope.commandText.trim();
          if(cmd == 'commit'){
            console.log("Committing board state: " + $scope.board.data);
            $scope.board.data.$save();
            $scope.board.committed = true;
          }
          else { // Process command
            console.log("Command applied");
            evalCommand(cmd);
            $scope.board.committed = false;


          }

          $scope.history.push($scope.commandText);
          $scope.commandText = "";
        }
        else if (charCode == 38) {  // Up key
          console.log("Up key pressed.");
        }


      };

      $scope.previousCommand = function() {

      };

  }]);
})();
