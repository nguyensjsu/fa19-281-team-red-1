import React from 'react'
import MyNavbar from './MyNavbar/MyNavbar'
import cookie from 'react-cookies'
import { Redirect } from 'react-router-dom';

class Main extends React.Component {
    render() {

        let redirect = cookie.load('username') ? null : <Redirect to="/login" />;
        return (
            <React.Fragment>
                {redirect}
                <MyNavbar></MyNavbar>
            </React.Fragment>
        )
    }
}

export default Main;