/*
 * This javascript needs serious changes.
 * The aproach was, first build a working feature then refactor.
 */
$(document).ready(function() {
	/**
	 * Opens a modal which asks for user name.
	 * @todo validation: only letters, max-length=8, unique in the game 
	 */
	var getUserName = function() {
		$('#modal-username').modal('show');

		$('button#btn-join').click(function() {
			user = $('input#username').val();
			
			startGame(user);

			$('#modal-username').modal('hide');
		});
	};
	
	// current user name
	var user = "";
	if (user == "") {
		getUserName();	
	}

	/**
	 * Adds estimation to game results
	 */
	var updateEstimation = function (user, estimate) {
		var p = $('#game').find('p#' + user).first();
		
		if (p.length === 0) {
			p = $('<p id="' + user + '">');
			$('#game').append(p);
		}

		span = $('<span class="label label-info">' + estimate + '</span>')
		p.text(user);
		p.append(span);
	};

	/**
	 * Adds players to the list of active users
	 */
	var updatePlayersList = function (playerObject) {
		var playerName = playerObject.Name;
		var player = $('<li />');
		player.text(playerName);
		if (user == playerName) {
			player.addClass("text-primary");
		}

		if (playerObject.Vote > 0) {
			updateEstimation(playerName, playerObject.Vote)
		}

		$('#players').append(player);
	};

	var startGame = function (user) {
		var socket = io();

		socket.emit('new_player', user);
		updatePlayersList(user);

		$('.estimates a').click(function() {
			var estimate = "" + $(this).data('val');
			updateEstimation(user, estimate);
			socket.emit('estimate', user + ":" + estimate);
		});

		socket.on('new_player', function(msg) {
			$('#players').empty()

			$.each(msg, function(i,e) {
				updatePlayersList(e);
			})
		});

		socket.on('estimate', function(msg) {
			var data = msg.split(':');
			updateEstimation(data[0], data[1]);
		});
	};
});

