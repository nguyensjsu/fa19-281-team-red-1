import React from 'react';
import { InputGroup, FormControl, Container, Button } from 'react-bootstrap';

class UrlShortener extends React.Component {

    state = {
        url : "",
        short_url: ""
    }

    onChange = e => this.setState({ [e.target.name]: e.target.value })

    onGenerateClickHandler = () => {
        console.log("[UrlShortener] Current State:")
        console.log(this.state)

        // TODO: send request to backend
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
                    <Button onClick={this.onGenerateClickHandler}>Generate</Button>
                </Container>
            </React.Fragment>
        )
    }
}

export default UrlShortener;