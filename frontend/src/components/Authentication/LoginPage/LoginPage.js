import React from 'react';
import { Component } from 'react';
import { Form, Button, Container } from 'react-bootstrap';

import styles from './LoginPage.module.css';

class LoginPage extends Component {
    render() {
        return (
            <Container className={styles.LoginPage}>
                <Form>
                    <Form.Group controlId="formUsername">
                        <Form.Label>Username</Form.Label>
                        <Form.Control type="text" placeholder="Enter username" />
                    </Form.Group>

                    <Form.Group controlId="formPassword">
                        <Form.Label>Password</Form.Label>
                        <Form.Control type="password" placeholder="Password" />
                    </Form.Group>
                    <Button variant="success" type="submit" className={styles.button}>
                        Login
                    </Button>
                    <Button variant="primary" type="submit">
                        Signup
                    </Button>
                </Form>
            </Container>
        )
    }
}

export default LoginPage;
