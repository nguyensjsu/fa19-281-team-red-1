import React from 'react';
import { Component } from 'react';
import LoginPage from './LoginPage/LoginPage'
import SignupPage from './SignupPage/SignupPage'
import { connect } from 'react-redux';
import { LOGIN_PAGE } from '../../redux/actions'

class Authentication extends Component {

    render() {
        return (
            <React.Fragment>
                {this.props.page === LOGIN_PAGE ? <LoginPage /> : <SignupPage />}
            </React.Fragment>
        )
    }
}

const mapStateToProps = state => ({
    page: state.userLoggedin.page
})

export default connect(mapStateToProps)(Authentication);