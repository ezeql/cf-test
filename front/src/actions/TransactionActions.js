'use strict';
var AppDispatcher = require('../dispatcher/FrontAppDispatcher');
var TransactionConstants = require('../constants/TransactionConstants');

var TransactionActions = {
    create: function(transaction) {
        AppDispatcher.dispatch({
            actionType: TransactionConstants.TODO_CREATE,
            transaction: transaction
        });
    },

    startFetching: function() {
        AppDispatcher.dispatch({
            actionType: TransactionConstants.START,
            data: []
        });
    }
};

module.exports = TransactionActions;