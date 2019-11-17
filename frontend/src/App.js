import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

import MyNavbar from './components/MyNavbar/MyNavbar'
import LoginPage from './components/Authentication/LoginPage/LoginPage'

function App() {
  return (
    <Router>
      <div className="App">
        <MyNavbar></MyNavbar>
        <LoginPage></LoginPage>
      </div>
    </Router>
  );
}

export default App;
