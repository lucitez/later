import { AsyncStorage } from 'react-native'
import store from "../store"
import Network from './Network'
import { authHeader } from './headers'
import jwtDecode from 'jwt-decode'
import * as actions from '../actions'
import { Buffer } from 'buffer'

const SESSION_TIMEOUT_THRESHOLD = 300 // Will refresh the access token 5 minutes before it expires

let sessionTimeout = null

const setSessionTimeout = (duration) => {
    clearTimeout(sessionTimeout);
    sessionTimeout = setTimeout(
        refreshTokens,
        duration - SESSION_TIMEOUT_THRESHOLD
    );
};

const onSuccess = ({ accessToken, refreshToken }) => {
    let decodedAT = jwtDecode(accessToken)
    let userId = decodedAT.sub
    let expiresAt = decodedAT.exp

    store.dispatch(actions.setTokens({ accessToken, refreshToken }))
    store.dispatch(actions.setUserId(userId))

    AsyncStorage.setItem('refresh_token', refreshToken)

    setSessionTimeout(expiresAt);
}

const onFailure = (exception) => {
    store.dispatch(actions.clearTokens())
    AsyncStorage.removeItem('refresh_token')

    throw exception
}

export const refreshTokens = () => {
    let refreshToken = store.getState().auth.tokens.refreshToken

    if (!refreshToken) {
        return Promise.reject("No refresh token")
    }

    let refreshAuthHeader = authHeader(refreshToken)
    return Network.POST('/auth/refresh', {}, refreshAuthHeader)
        .then(onSuccess)
        .catch(onFailure)

}

export const logIn = ({ identifier, password }) => {
    let loginToken = new Buffer(`${identifier}:${password}`).toString('base64')

    return Network.POST('/auth/login', {}, authHeader(loginToken))
        .then(onSuccess)
        .catch(onFailure)
}

export const signUp = ({ name, username, phoneNumber, password }) => {
    let token = new Buffer(`${phoneNumber}:${password}`).toString('base64')
    let header = authHeader(token)
    let body = { name, username }

    return Network.POST('/auth/sign-up', body, header)
        .then(onSuccess)
        .catch(onFailure)
}

