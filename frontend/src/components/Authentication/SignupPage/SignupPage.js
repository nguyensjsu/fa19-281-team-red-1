import React from 'react';
import { Component } from 'react';
import { Form, Button, Container } from 'react-bootstrap';

import styles from '../Authentication.module.css';

class SignupPage extends Component {

    state = {
        username : "",
        password: "",
        confirmPassword: ""
    }

    onChange = e => this.setState({ [e.target.name]: e.target.value })

    onLoginClickHandler = () => {
        // this.props.switchPage(LOGIN_PAGE)
        this.props.history.push('/login')
    }

    onSignupClickHandler = () => {

        // TODO: send request to server

        console.log("[Signup Page] Current State")
        console.log(this.state)

        this.props.history.push('/login')
    }

    render() {
        return (
            <Container className={styles.AuthPage}>
                <Form>
                    <Form.Group controlId="formUsername">
                        <Form.Label>Username</Form.Label>
                        <Form.Control type="text" placeholder="Enter username" name="username" value={this.state.username} onChange={this.onChange} />
                    </Form.Group>

                    <Form.Group controlId="formPassword">
                        <Form.Label>Password</Form.Label>
                        <Form.Control type="password" placeholder="Password" name="password" value={this.state.password} onChange={this.onChange} />
                    </Form.Group>

                    <Form.Group controlId="formConfirmPassword">
                        <Form.Label>Confirm Password</Form.Label>
                        <Form.Control type="password" placeholder="Confirm Password" name="confirmPassword" value={this.state.confirmPassword} onChange={this.onChange} />
                    </Form.Group>
                    <Button variant="success" className={styles.button} onClick={this.onSignupClickHandler}>
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

export default (SignupPage);
