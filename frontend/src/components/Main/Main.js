import React from 'react'
import MyNavbar from './MyNavbar/MyNavbar'
import cookie from 'react-cookies'
import { Redirect, Switch, Route } from 'react-router-dom';
import UrlShortener from './UrlShortener/UrlShortener'
import TopTen from './TopTen/TopTen'
import History from './History/History'

class Main extends React.Component {
    render() {

        let redirect = cookie.load('username') ? null : <Redirect to="/login" />;
        return (
            <React.Fragment>
                {redirect}
                <MyNavbar></MyNavbar>
                <Switch>
                    <Route path="/" exact component={UrlShortener}></Route>
                    <Route path="/history" exact component={History}></Route>
                    <Route path="/top" exact component={TopTen}></Route>
                </Switch>
            </React.Fragment>
        )
    }
}

export default Main;