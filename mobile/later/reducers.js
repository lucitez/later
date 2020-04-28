import {
    SET_TOKENS,
    CLEAR_TOKENS,
    SET_USER_ID,
} from './actions'
import { combineReducers } from 'redux'

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

const userId = (state = '', action) => {
    switch (action.type) {
        case SET_USER_ID:
            return action.userId
        default:
            return state
    }
}

const auth = combineReducers({
    tokens,
    userId,
})


const laterApp = combineReducers({
    auth,
})

export default laterApp