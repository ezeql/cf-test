'use strict';
var AppDispatcher = require('../dispatcher/FrontAppDispatcher');
var EventEmitter = require('events').EventEmitter;
var TransactionConstants = require('../constants/TransactionConstants.js');
var assign = require('object-assign');
import moment from 'moment';

const CHANGE_EVENT = 'change';

var _transactions = [];
var _countryCount = {};

/**
 * Create a TODO item.
 * @param  {string} text The content of the TODO
 */
function create(transaction) {
    _transactions.push({
        "id": _transactions.length,
        "data": transaction
    });

    if (_countryCount[transaction.originatingCountry]) {
        _countryCount[transaction.originatingCountry] = 
        _countryCount[transaction.originatingCountry] + 1;
    } else {
        _countryCount[transaction.originatingCountry] = 1;
    }
}

var TransactionStore = assign({}, EventEmitter.prototype, {

    getAll: function() {
        return _transactions;
    },

    countriesDataArray: function() {
        return Object.keys(_countryCount).map(
            key => [ key,_countryCount[key] ] );
    },

    getCountryCount: function() {
        return Object.keys(_countryCount).map(
            key => ({ "label" : key, "value" : _countryCount[key] }));
    },
    getTransactionByDate:function(from,to){
        const filtered = this.getAll().filter( 
                            t => t.data.currencyFrom === from && 
                                 t.data.currencyTo   === to);

        const parse = moment.bind('D-MMM-YY H:m:s');

        const points = filtered.map( t =>( { date:parse(t.data.timePlaced).valueOf(), value:t.data.rate } ) );

        return {
            name: from + '/' + to,
            values: points
        };
        
    },
    emitChange: function() {
        this.emit(CHANGE_EVENT);
    },

    addChangeListener: function(callback) {
        this.on(CHANGE_EVENT, callback);
    },

    removeChangeListener: function(callback) {
        this.removeListener(CHANGE_EVENT, callback);
    }
});

AppDispatcher.register(function(action) {

    switch (action.actionType) {
        case TransactionConstants.TODO_CREATE:
            create(action.transaction);
            TransactionStore.emitChange();
            break;

        case TransactionConstants.TRANSACTION_START:

            TransactionStore.emitChange();
            break;
        default:
            // no op
    }
});

module.exports = TransactionStore;