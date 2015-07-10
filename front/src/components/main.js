'use strict';

import React from 'react/addons';
import Router from 'react-router';  
import { DefaultRoute, Route } from 'react-router';

import CFApp from './FrontApp';
import GeoMap from './GeoMap';
import Grid from './Grid';
import CountriesTreeMap from './CountriesTreeMap';
import Graphs from './Graphs';

import TransactionService from '../services/TransactionService';

const content = document.getElementById('content');

const Routes = (
    <Route name='/' handler={CFApp}>
        <Route name="grid" handler={Grid}/>
        <Route name="map" handler={GeoMap}/>
        <Route name="countries" handler={CountriesTreeMap}/>
        <Route name="graphs" handler={Graphs}/>
    </Route>
);

TransactionService.start();

Router.run(Routes, function(Handler, state) {
    React.render(<Handler/>, content);
});

