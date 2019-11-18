import React from 'react';
import { Component } from 'react';
import LoginPage from './LoginPage/LoginPage'
import SignupPage from './SignupPage/SignupPage'
import { connect } from 'react-redux';
import { Switch, Route } from 'react-router-dom'

class Authentication extends Component {

    render() {
        return (
            <React.Fragment>
                <Switch>
                    <Route path='/login' component={LoginPage}></Route>
                    <Route path='/signup' component={SignupPage}></Route>
                </Switch>
            </React.Fragment>
        )
    }
}

const mapStateToProps = state => ({
    page: state.userLoggedin.page
})

export default connect(mapStateToProps)(Authentication);