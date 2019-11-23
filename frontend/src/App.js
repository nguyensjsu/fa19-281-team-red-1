import React from 'react';
import {
  HashRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import { connect } from 'react-redux'
import cookie from 'react-cookies'

import Main from './components/Main/Main'
import Authentication from './components/Authentication/Authentication'

class App extends React.Component {
  render() {
    return (
      <Router>
        <div className="App">
          <Switch>
            <Route path={['/login', '/signup']} component={Authentication}></Route>
            <Route path='/' component={Main}></Route>
          </Switch>
        </div>
      </Router>
    )
  }
}


export default (App);
