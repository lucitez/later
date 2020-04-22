export const SET_TOKENS = 'SET_TOKENS'
export const CLEAR_TOKENS = 'CLEAR_TOKENS'

export const setTokens = (tokens) => {
    return {
        type: SET_TOKENS,
        tokens,
    }
}

export const clearTokens = token => {
    return { type: SET_REFRESH_TOKEN, token }
}