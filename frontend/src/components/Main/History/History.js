import React from 'react';
import { Container, Table } from 'react-bootstrap';
import axios from 'axios';
import { hostname } from '../../../config';
import cookie from 'react-cookies';

class History extends React.Component {

    state = {
        history: []
    }

    componentDidMount() {
        axios.get(hostname + '/history', {
            params: {
                Username: cookie.load('username')
            }
        }).then(res => {
            console.log(res.data)
            this.setState({
                ...this.state,
                history: res.data.History
            })
        })
    }

    render() {
        return (
            <Container>
                <h1>History</h1>
                <Table striped bordered hover responsive>
                    <thead>
                        <tr>
                            <th>Url</th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            this.state.history.map((item, i) => {
                                return (
                                    <tr key={i}>
                                        <td>{item}</td>
                                    </tr>
                                )
                            })
                        }
                    </tbody>
                </Table>
            </Container>
        )
    }
}

export default History;