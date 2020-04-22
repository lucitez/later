import React, { useState, useContext } from 'react'
import { useDispatch } from 'react-redux'
import { Buffer } from 'buffer'
import { StyleSheet, View, Text, AsyncStorage } from 'react-native'
import { Button } from '../components/common'
import Network from '../util/Network'
import * as actions from '../actions'
import { AuthContext } from '../context'
import { colors } from '../assets/colors'
import { PlainText, Password } from '../components/forms'

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

    const [formData, setFormData] = useState({
        identifier: '',
        password: ''
    })

    const [error, setError] = useState(null)

    const onFormDataChange = (name, value) => {
        setFormData({
            ...formData,
            [name]: value
        })
    }

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
        <View style={styles.container}>
            <View style={styles.header} >

            </View>
            <View style={styles.details}>
                <View style={styles.logoContainer}>

                </View>
                <View style={styles.titleContainer}>
                    <Text style={style.title}>Sign Up For Later</Text>
                </View>
            </View>
            <View style={styles.formContainer}>
                <View style={styles.identifierContainer}>
                    <PlainText
                        name='identifier'
                        title='Email, Username, or Phone Number'
                        value={formData.identifier}
                        onChange={onFormDataChange}
                    />
                </View>
                <View style={styles.passwordContianer}>
                    <Password
                        name='password'
                        title='Password'
                        value={formData.password}
                        onChange={onFormDataChange}
                    />
                </View>
                <View style={styles.submitButtonContainer}>
                    <Button name='Submit' />
                </View>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.primary
    }
})

export default SignupScreen