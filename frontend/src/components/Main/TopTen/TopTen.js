import React from 'react';
import { Container, Table } from 'react-bootstrap';
import axios from 'axios';
import { hostname } from '../../../config';
import cookie from 'react-cookies';

class TopTen extends React.Component {

    state = {
        top: []
    }

    componentDidMount() {
        axios.get(hostname + '/top').then(res => {
            console.log(res.data)
            this.setState({
                ...this.state,
                top: res.data
            })
        })
    }

    render() {
        return (
            <Container>
                <h1>Top Domains</h1>
                <Table striped bordered hover responsive>
                    <thead>
                        <tr>
                            <th>Domain</th>
                            <th>Count</th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            this.state.top.map((item, i) => {
                                return (
                                    <tr key={i}>
                                        <td>{item.Domain}</td>
                                        <td>{item.Counter}</td>
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

export default TopTen;