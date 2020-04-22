import {
    SET_TOKENS,
    CLEAR_TOKENS
} from './actions'
import { combineReducers } from 'redux'

const initialState = {
    auth: {
        tokens: {
            accessToken: '',
            refreshToken: ''
        },
        somethingElse: []
    }
}

const tokens = (state = {}, action) => {
    switch (action.type) {
        case SET_TOKENS:
            return action.tokens
        case CLEAR_TOKENS:
            return {}
        default:
            return state
    }
}

const auth = combineReducers({
    tokens,
})


const laterApp = combineReducers({
    auth,
})

export default laterApp