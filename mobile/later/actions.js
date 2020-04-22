export const SET_TOKENS = 'SET_TOKENS'
export const CLEAR_TOKENS = 'CLEAR_TOKENS'

export const setTokens = (accessToken, refreshToken) => {
    return {
        type: SET_TOKENS,
        tokens: {
            accessToken: accessToken,
            refreshToken: refreshToken
        }
    }
}

export const clearTokens = token => {
    return { type: SET_REFRESH_TOKEN, token }
}