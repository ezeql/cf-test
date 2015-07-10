import React from 'react';
import { ReactBootstrap, Navbar, Nav } from 'react-bootstrap';
import { ReactRouterBootstrap, NavItemLink } from 'react-router-bootstrap';

export default class Menu extends React.Component {
    render() {
        return (
            <Navbar brand='CurrencyFair-Test' toggleNavKey={0}>
                <Nav eventKey={0}>
                    <NavItemLink to="/map" eventKey={1}>Map</NavItemLink>
                    <NavItemLink to="/grid" eventKey={2}>Grid</NavItemLink>
                    <NavItemLink to="/countries" eventKey={3}>Countries</NavItemLink>
                    <NavItemLink to="/graphs" eventKey={4}>Graphs</NavItemLink>
                </Nav>
            </Navbar>
        );
    }
}