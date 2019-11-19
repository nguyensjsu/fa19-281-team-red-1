import React from 'react';
import { Component } from 'react';
import { Form, Button, Container } from 'react-bootstrap';
import axios from 'axios';
import { hostname } from '../../../config'

import styles from '../Authentication.module.css';

class SignupPage extends Component {

    state = {
        username: "",
        password: "",
        confirmPassword: "",
        error: false
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

        let resolve = res => {
            console.log(res)
            if (res.status == 201) {
                this.props.history.push('/login')
            } else {
                console.log("Signup Error: " + res.data)
                this.setState({
                    ...this.state,
                    error: true
                })
            }
        }

        axios.post(hostname + '/signup', {
            username: this.state.username,
            password: this.state.password
        }).then(res => {
            this.props.history.push('/login')
        }).catch(err => {
            this.setState({
                ...this.state,
                error: true
            })
        })

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
                    {this.state.error ? <p style={{ color: 'red' }}>Signup Failed</p> : null}
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
