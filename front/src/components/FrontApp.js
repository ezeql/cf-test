import React from 'react';
import {RouteHandler} from 'react-router';
import Menu from './Menu';
// var ReactTransitionGroup = React.addons.TransitionGroup;

export default class FrontApp extends React.Component {
    render(){
        return (
        <div>
            <Menu></Menu>
             <RouteHandler></RouteHandler>
        </div>);
    }
}


