import React from 'react';
import { Component } from 'react';
import { Form, Button, Container } from 'react-bootstrap';

import styles from '../Authentication.module.css';
import cookie from 'react-cookies';

class LoginPage extends Component {
    state = {
        username: "",
        password: ""
    }

    onChange = e => this.setState({ [e.target.name]: e.target.value })

    onSignupClickHandler = () => {
        this.props.history.push('/signup')
    }

    onLoginClickHandler = () => {

        // TODO: send request to server

        console.log("[Login Page] Current State")
        console.log(this.state)

        cookie.save('username', 'test')
        this.props.history.push('/')
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

export default (LoginPage);
