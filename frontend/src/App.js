import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import { connect } from 'react-redux'

import MyNavbar from './components/MyNavbar/MyNavbar'
import Authentication from './components/Authentication/Authentication'

class App extends React.Component {
  render() {
    return (
      <Router>
        <div className="App">
          {
            this.props.userAuth.userLoggedin ? <MyNavbar /> : <Authentication />
          }
          {/* <MyNavbar></MyNavbar>
          <LoginPage></LoginPage>
          <SignupPage></SignupPage> */}
        </div>
      </Router>
    )
  }
}
const mapStateToProps = state => ({
  userAuth: state.userLoggedin
})

export default connect(mapStateToProps)(App);
