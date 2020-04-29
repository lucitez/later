export const SET_TOKENS = 'SET_TOKENS'
export const SET_USER_ID = 'SET_USER_ID'
export const CLEAR_TOKENS = 'CLEAR_TOKENS'

export const setTokens = tokens => {
    return {
        type: SET_TOKENS,
        tokens,
    }
}

export const setUserId = userId => {
    return {
        type: SET_USER_ID,
        userId,
    }
}

export const clearTokens = () => {
    return { type: CLEAR_TOKENS }
}