import * as actions from '../actions/index';

const defaultState = {
    userLoggedin: false,
    page: actions.LOGIN_PAGE,
    username: null
}

const userAuth = (state = defaultState, action) => {
    switch (action.type) {
        case actions.USER_LOGGEDIN:
            return {
                ...state,
                userLoggedin: action.payload.success,
                username: action.payload.username
            }
        case actions.SWITCH_PAGE: 
            return {
                ...state,
                page: action.payload.page
            }
        default:
            return state
    }
}

export default userAuth;