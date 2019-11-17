export const USER_LOGGEDIN = "USER_LOGGEDIN";
export const SWITCH_PAGE = "SWITCH_PAGE";
export const LOGIN_PAGE = "LOGIN_PAGE";
export const SIGNUP_PAGE = "SIGNUP_PAGE";


export const userLoginAction = () => (
    {
        type: USER_LOGGEDIN,
        payload: {
            success: true
        }
    }
)

export const switchPage = (page) => (
    {
        type: SWITCH_PAGE,
        payload: {
            page: page
        }
    }
)