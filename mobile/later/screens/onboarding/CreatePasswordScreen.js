import React, { useState, useContext } from 'react'
import { useDispatch } from 'react-redux'
import { Buffer } from 'buffer'
import { StyleSheet, View, Text, AsyncStorage } from 'react-native'
import OnboardingFormWrapper from './OnboardingFormWrapper'
import { Password } from '../../components/forms'
import { Button, BackIcon } from '../../components/common';
import Network from '../../util/Network';
import { authHeader } from '../../util/headers'
import { AuthContext } from '../../context'
import { colors } from '../../assets/colors';
import * as actions from '../../actions'

const setRefreshToken = async (token) => {
    await AsyncStorage.setItem('refresh_token', token)
}

export default function CreatePasswordScreen({ navigation, route }) {
    const { signIn } = useContext(AuthContext)
    const dispatch = useDispatch()

    const [formData, setFormData] = useState(route.params.formData)
    const [error, setError] = useState(null)

    const onSubmit = () => {
        setError(null)

        if (formData.password.error) {
            setError(formData.password.error)
            return
        }

        signUp(formData)
            .then(newTokens => {
                dispatch(actions.setTokens(newTokens))
                setRefreshToken(newTokens.refreshToken)
                signIn()
            })
            .catch(err => console.error(err))
    }

    const onFormDataChange = (name, value, error) => {
        setFormData({
            ...formData,
            [name]: {
                value: value,
                error: error
            }
        })
    }

    return (
        <OnboardingFormWrapper
            title='Create Password'
            leftIcon={<BackIcon navigation={navigation} color={colors.primary} />}
        >
            <Password
                inputProps={{ autoFocus: true }}
                name='password'
                title='Password'
                onChange={onFormDataChange}
            />
            <Button name='Submit' theme='primary' size='medium' onPress={onSubmit} />
            {error &&
                <View style={styles.errorMessageContainer}>
                    <Text style={styles.errorMessage}>{error}</Text>
                </View>}
        </OnboardingFormWrapper>
    )
}

const signUp = formData => {
    let token = new Buffer(`${formData.phoneNumber.value}:${formData.password.value}`).toString('base64')
    let header = authHeader(token)
    let body = {
        name: formData.name.value,
        username: formData.username.value
    }

    return Network.POST('/auth/sign-up', body, header)
}

const styles = StyleSheet.create({
    passwordInputContainer: {
        flex: 1,
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

