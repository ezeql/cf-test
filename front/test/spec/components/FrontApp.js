'use strict';

describe('FrontApp', function () {
  var React = require('react/addons');
  var FrontApp, component;

  beforeEach(function () {
    var container = document.createElement('div');
    container.id = 'content';
    document.body.appendChild(container);

    FrontApp = require('components/FrontApp.js');
    component = React.createElement(FrontApp);
  });

  it('should create a new instance of FrontApp', function () {
    expect(component).toBeDefined();
  });
});
