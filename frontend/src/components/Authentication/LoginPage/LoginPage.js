import React from 'react';
import { Component } from 'react';
import { Form, Button, Container } from 'react-bootstrap';
import { connect } from 'react-redux';
import { userLoginAction, SIGNUP_PAGE } from '../../../redux/actions'

import styles from '../Authentication.module.css';

class LoginPage extends Component {

    onSignupClickHandler = () => {
        // this.props.switchPage(SIGNUP_PAGE)
    }

    onLoginClickHandler = () => {
        this.props.login("test");
    }

    render() {
        return (
            <Container className={styles.AuthPage}>
                <Form>
                    <Form.Group controlId="formUsername">
                        <Form.Label>Username</Form.Label>
                        <Form.Control type="text" placeholder="Enter username" />
                    </Form.Group>

                    <Form.Group controlId="formPassword">
                        <Form.Label>Password</Form.Label>
                        <Form.Control type="password" placeholder="Password" />
                    </Form.Group>
                    <Button variant="success" className={styles.button} onClick={this.onLoginClickHandler}>
                        Login
                    </Button>
                    <Button variant="primary" onClick={this.onSignupClickHandler}>
                        Signup
                    </Button>
                </Form>
            </Container>
        )
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        login: (username) => dispatch(userLoginAction(username))
    };
};

export default connect(null, mapDispatchToProps)(LoginPage);
