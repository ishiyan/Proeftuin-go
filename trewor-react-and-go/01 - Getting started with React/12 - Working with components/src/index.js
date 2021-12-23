import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import AppFooter from './AppFooter';

class App extends Component {

  render() {
    return (
      <div>
        <div>
          <h1>Hello, world!</h1>
        </div>
        <AppFooter />
      </div>
    );
  }
}


ReactDOM.render(<App />, document.getElementById('root'));
