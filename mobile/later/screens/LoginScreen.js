import React from 'react'
import { useDispatch } from 'react-redux'
import { Buffer } from 'buffer'
import { StyleSheet, View, Text, AsyncStorage } from 'react-native'
import { Button } from '../components/common'
import Network from '../util/Network'
import * as actions from '../actions'

const setRefreshToken = async (token) => {
    await AsyncStorage.setItem('@refresh_token', token)
}

function LoginScreen() {
    const dispatch = useDispatch()
    const login = (email, password) => {
        let headers = {
            Authorization: `Basic ${new Buffer(`${email}:${password}`).toString('base64')}`,
        }

        Network.POST('/auth/login', {}, headers)
            .then(newTokens => {
                dispatch(actions.setTokens(newTokens))
                try {
                    setRefreshToken(newTokens.refreshToken)
                } catch (err) { // This should never happen
                    console.error(err)
                }
            })
            .catch(err => console.error(err))
    }

    return (
        <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
            <Button name='login' size='large' theme='primary' onPress={() => login('test@test.com', 'pass')} />
            <Text>asdf</Text>
        </View>
    )
}

const styles = StyleSheet.create({

})

export default LoginScreen