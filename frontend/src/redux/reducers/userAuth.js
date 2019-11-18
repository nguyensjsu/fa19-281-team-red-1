import * as actions from '../actions/index';
import cookie from 'react-cookies';

const defaultState = {
    username: cookie.load('username'),
}

const userAuth = (state = defaultState, action) => {
    switch (action.type) {
        case actions.USER_LOGGEDIN:
            return {
                ...state,
                username: action.payload.username
            }
        default:
            return state
    }
}

export default userAuth;