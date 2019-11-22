import React from 'react';
import { InputGroup, FormControl, Container, Button } from 'react-bootstrap';
import axios from 'axios'
import { hostname } from '../../../config'
import cookie from 'react-cookies';

class UrlShortener extends React.Component {

    state = {
        url: "",
        short_url: "",
        message: ""
    }

    onChange = e => this.setState({ [e.target.name]: e.target.value })

    onGenerateClickHandler = () => {
        console.log("[UrlShortener] Current State:")
        console.log(this.state)

        if (this.state.url.trim() === "") {
            this.setState({
                ...this.state,
                message: "Url cannot be empty"
            })
            return
        }

        // TODO: send request to backend
        axios.post(hostname + '/shorten', {
            url: this.state.url,
            Username: cookie.load("username")
        }).then(res => {
            this.setState({
                ...this.state,
                message: "",
                short_url: res.data.ShortUrl
            })
        },
            err => {
                this.setState({
                    ...this.state,
                    message: err.response.data
                })
            }
        )
    }

    render() {
        return (
            <React.Fragment>
                <Container>

                    <h1>URL Shortener</h1>
                    <InputGroup className="mb-3">
                        <InputGroup.Prepend>
                            <InputGroup.Text id="url">Url: </InputGroup.Text>
                        </InputGroup.Prepend>
                        <FormControl
                            placeholder="Url"
                            aria-label="Url"
                            aria-describedby="url"
                            value={this.state.url}
                            name="url"
                            onChange={this.onChange}
                        />
                    </InputGroup>
                    <InputGroup className="mb-3">
                        <InputGroup.Prepend>
                            <InputGroup.Text id="shorturl">Short Url: </InputGroup.Text>
                        </InputGroup.Prepend>
                        <FormControl
                            disabled
                            placeholder="Short Url"
                            aria-label="Short Url"
                            aria-describedby="shorturl"
                            value={this.state.short_url}
                            name="short_url"
                            onChange={this.onChange}
                        />
                    </InputGroup>
                    <p style={{ color: "red" }}>{this.state.message}</p>
                    <Button onClick={this.onGenerateClickHandler}>Generate</Button>
                </Container>
            </React.Fragment>
        )
    }
}

export default UrlShortener;