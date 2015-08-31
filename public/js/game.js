$(document).ready(function() {
	var getUserName = function() {
		$('#modal-username').modal('show');

		$('button#btn-join').click(function() {
			user = $('input#username').val();
			
			startGame(user);

			$('#modal-username').modal('hide');
		});
	};
	
	getUserName();

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

	var startGame = function (user) {
		var socket = io();

		$('.estimates a').click(function() {
			var estimate = "" + $(this).data('val');
			updateEstimation(user, estimate);
			socket.emit('estimate', user + ":" + estimate);
		});

		socket.on('estimate', function(msg) {
			var data = msg.split(':');
			updateEstimation(data[0], data[1]);
		});
	};
});

