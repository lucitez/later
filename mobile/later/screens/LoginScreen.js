import React from 'react'
import { Buffer } from 'buffer'
import { StyleSheet, View, Text } from 'react-native'
import { Button } from '../components/common'
import Network from '../util/Network'

function LoginScreen() {
    const login = (email, password) => {

        let headers = {
            Authorization: `Basic ${new Buffer(`${email}:${password}`).toString('base64')}`,
        }

        console.log(headers)
        Network.POST('/auth/login', {}, headers)
            .then(res => console.log(res))
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