'use strict';
import AppDispatcher from '../dispatcher/FrontAppDispatcher';
import TransactionConstants from '../constants/TransactionConstants';
import TransactionActions from '../actions/TransactionActions';

function startService() {
	let conn;
	if (window.WebSocket) {
		conn = new WebSocket("ws://" + window.location.hostname + ":8090/ws");
		conn.onclose = function(evt) {
			//AppActions.disconnected(); //hypothetical
		};
		conn.onmessage = function(evt) {
			try {
				var whole = JSON.parse(evt.data);
				TransactionActions.create(whole);
			} catch (e) {
				 console.log("error");
				 console.log(evt.data);
			}
		};
	} else {
		console.log("Your browser does not support WebSockets");
	}
}

var TransactionService = {
	start: startService
};

module.exports = TransactionService;