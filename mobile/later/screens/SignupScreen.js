import React, { useContext } from 'react'
import { useDispatch } from 'react-redux'
import { Buffer } from 'buffer'
import { StyleSheet, View, Text, AsyncStorage } from 'react-native'
import { Button } from '../components/common'
import Network from '../util/Network'
import * as actions from '../actions'
import { AuthContext } from '../context'

const setRefreshToken = async (token) => {
    try {
        await AsyncStorage.setItem('refresh_token', token)
    } catch (e) {
        console.error(e)
    }
}

function SignupScreen() {
    const { signIn } = useContext(AuthContext)
    const dispatch = useDispatch()
    const login = (email, password) => {

        let headers = {
            Authorization: `Basic ${new Buffer(`${email}:${password}`).toString('base64')}`,
        }

        Network.POST('/auth/login', {}, headers)
            .then(res => {
                dispatch(actions.setTokens(res.accessToken, res.refreshToken))
                console.log(res)
                setRefreshToken(res.refreshToken)
                signIn()
            })
            .catch(err => console.error(err))
    }

    return (
        <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
            <Button name='Sign Up' size='large' theme='primary' onPress={() => login('test@test.com', 'pass')} />
            <Text>asdf</Text>
        </View>
    )
}

const styles = StyleSheet.create({

})

export default SignupScreen