import React, { useState, useContext } from 'react'
import { useDispatch } from 'react-redux'
import { Buffer } from 'buffer'
import { StyleSheet, View, Text, AsyncStorage } from 'react-native'
import { Button, BackIcon } from '../../components/common'
import Network from '../../util/Network'
import * as actions from '../../actions'
import { AuthContext } from '../../context'
import { colors } from '../../assets/colors'
import { PlainText, Password } from '../../components/forms'
import { authHeader } from '../../util/headers'
import OnboardingFormWrapper from './OnboardingFormWrapper'
import jwtDecode from 'jwt-decode'

const setRefreshToken = async (token) => {
    await AsyncStorage.setItem('refresh_token', token)
}

function LoginScreen({ navigation }) {
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
            .then(tokens => {
                try {
                    let userId = jwtDecode(tokens.accessToken).sub
                    dispatch(actions.setTokens(tokens))
                    dispatch(actions.setUserId(userId))
                    setRefreshToken(tokens.refreshToken)
                    signIn()
                }
                catch (e) {
                    console.error(e)
                }
            })
            .catch(() => setError('Invalid username or password'))
    }

    return (
        <OnboardingFormWrapper
            title='Log In to Later'
            leftIcon={<BackIcon navigation={navigation} color={colors.primary} />}
        >
            <PlainText
                name='identifier'
                title='Email, Username, or Phone Number'
                value={formData.identifier}
                onChange={onFormDataChange}
            />
            <Password
                name='password'
                title='Password'
                value={formData.password}
                onChange={onFormDataChange}
            />
            <Button name='Submit' size='medium' theme='primary' onPress={() => login()} />
            {error &&
                <View style={styles.errorMessageContainer}>
                    <Text style={styles.errorMessage}>{error}</Text>
                </View>}
        </OnboardingFormWrapper>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
    header: {
        height: '7%',
        alignItems: 'flex-end',
        justifyContent: 'center',
    },
    signUp: {
        color: colors.primary
    },
    details: {
        alignItems: 'center',
        marginBottom: 30,
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