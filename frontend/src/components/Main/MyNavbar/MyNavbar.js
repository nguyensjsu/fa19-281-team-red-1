import { Component } from 'react';
import React from 'react';
import { Navbar, Nav, NavItem } from 'react-bootstrap';
import cookie from 'react-cookies';
import { withRouter } from 'react-router'
import { NavLink, Link } from 'react-router-dom'
import { LinkContainer } from "react-router-bootstrap";

class MyNavbar extends Component {

    onSignoutClickHandler = (e) => {
        e.preventDefault();
        cookie.remove("username")
        this.props.history.push("/")
    }

    render() {
        return (
            <Navbar bg="light" expand="lg">
                <Navbar.Brand href="/">Yet Another Url Shortener</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="mr-auto">
                        <LinkContainer to="/" exact>
                            <Nav.Link>
                                Url Shortener
                            </Nav.Link>
                        </LinkContainer>
                        <LinkContainer to="/history">
                            <Nav.Link>History</Nav.Link>
                        </LinkContainer>
                        <LinkContainer to="/top">
                            <Nav.Link>Top 10</Nav.Link>
                        </LinkContainer>
                    </Nav>
                    <Nav className="ml-auto">
                        <Nav.Link onClick={this.onSignoutClickHandler}>
                            Signout
                        </Nav.Link>
                    </Nav>
                </Navbar.Collapse>
            </Navbar>
        )
    }
}

export default withRouter(MyNavbar)