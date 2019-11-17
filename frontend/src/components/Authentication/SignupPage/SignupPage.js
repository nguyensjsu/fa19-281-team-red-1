import React from 'react';
import { Component } from 'react';
import { Form, Button, Container } from 'react-bootstrap';
import { connect } from 'react-redux';
import { LOGIN_PAGE, switchPage } from '../../../redux/actions'

import styles from '../Authentication.module.css';

class SignupPage extends Component {

    onLoginClickHandler = () => {
        this.props.switchPage(LOGIN_PAGE)
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

                    <Form.Group controlId="formConfirmPassword">
                        <Form.Label>Confirm Password</Form.Label>
                        <Form.Control type="password" placeholder="Confirm Password" />
                    </Form.Group>
                    <Button variant="primary" className={styles.button}>
                        Submit
                    </Button>
                    <Button variant="primary" onClick={this.onLoginClickHandler}>
                        Login
                    </Button>
                </Form>
            </Container>
        )
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        switchPage: (page) => dispatch(switchPage(page))
    };
};

export default connect(null, mapDispatchToProps)(SignupPage);
