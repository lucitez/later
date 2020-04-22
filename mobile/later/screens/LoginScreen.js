import React, { useState, useContext } from 'react'
import { useDispatch } from 'react-redux'
import { Buffer } from 'buffer'
import { StyleSheet, View, Text, AsyncStorage } from 'react-native'
import { Button, Icon } from '../components/common'
import Network from '../util/Network'
import * as actions from '../actions'
import { AuthContext } from '../context'
import { colors } from '../assets/colors'
import { PlainText, Password } from '../components/forms'
import { authHeader } from '../util/headers'

const setRefreshToken = async (token) => {
    await AsyncStorage.setItem('refresh_token', token)
}

function LoginScreen() {
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

    const login = () => {
        setError(null)
        let token = new Buffer(`${formData.identifier}:${formData.password}`).toString('base64')

        Network.POST('/auth/login', {}, authHeader(token))
            .then(newTokens => {
                dispatch(actions.setTokens(newTokens))
                setRefreshToken(newTokens.refreshToken)
                signIn()
            })
            .catch(err => {
                console.log(err)
                setError('Invalid username or password')
            })
    }

    return (
        <View style={styles.container}>
            <View style={styles.header}>
            </View>
            <View style={styles.details}>
                <View style={styles.logoContainer}>
                    <Icon type='share' size={50} color={colors.primary} />
                </View>
                <View style={styles.titleContainer}>
                    <Text style={styles.title}>Log In To Later</Text>
                </View>
            </View>
            <View style={styles.formContainer}>
                <View style={styles.formContentsContainer}>
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
                        <Button name='Submit' size='medium' theme='primary' onPress={() => login()} />
                    </View>
                    {error &&
                        <View style={styles.errorMessageContainer}>
                            <Text style={styles.errorMessage}>{error}</Text>
                        </View>}
                </View>

            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray
    },
    header: {
        height: '10%',
    },
    details: {
        alignItems: 'center',
        marginBottom: 30,
    },
    logoContainer: {
        padding: 10,
    },
    titleContainer: {
        padding: 10,
    },
    title: {
        fontSize: 24,
        fontWeight: '300'
    },
    formContainer: {
        alignItems: 'center'
    },
    formContentsContainer: {
        width: '80%',
        backgroundColor: colors.white,
        borderRadius: 10,
        padding: 20,
    },
    identifierContainer: {
        marginBottom: 5,
    },
    passwordContianer: {
        marginTop: 5,
        marginBottom: 10,
    },
    errorMessageContainer: {
        paddingTop: 5,
        alignItems: 'center'
    },
    errorMessage: {
        color: 'red',
        fontWeight: '300',
    }
})

export default LoginScreen