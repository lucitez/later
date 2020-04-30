import React, { useState } from 'react'
import { StyleSheet, View, Text, TouchableOpacity } from 'react-native'
import { Button } from '../../components/common'
import { colors } from '../../assets/colors'
import { PlainText, PhoneNumber } from '../../components/forms'
import OnboardingFormWrapper from './OnboardingFormWrapper'
import Network from '../../util/Network'

function SignupScreen({ navigation }) {

    const [formData, setFormData] = useState({
        name: {
            value: '',
            username: 'Name is required'
        },
        username: {
            value: '',
            error: 'Username is required'
        },
        phoneNumber: {
            value: '',
            error: 'Phone Number is required'
        },
    })

    const [submitting, setSubmitting] = useState(false)
    const [error, setError] = useState(null)

    const onFormDataChange = (name, value, error) => {
        setFormData({
            ...formData,
            [name]: {
                value: value,
                error: error
            }
        })
    }

    const signUp = () => {
        setSubmitting(true)
        setError(null)

        for (let [_, field] of Object.entries(formData)) {
            if (field.error) {
                setError(field.error)
                setSubmitting(false)
                return
            }
        }

        Network.GET('/auth/sign-up/check-conflicts', {
            phoneNumber: formData.phoneNumber.value,
            username: formData.username.value
        })
            .then(() => {
                setSubmitting(false)
                navigation.navigate('SMS', { formData })
            })
            .catch(err => {
                setError(err)
                setSubmitting(false)
            })

    }

    const onLoginPressed = () => {
        navigation.navigate('Login')
    }

    return (
        <OnboardingFormWrapper
            title='Welcome to Later!'
            description='create an account below'
            rightIcon={
                <View style={{ flexDirection: 'row' }}>
                    <Text>Have an account?</Text>
                    <TouchableOpacity style={styles.loginContainer} onPress={onLoginPressed}>
                        <Text style={styles.login}>Log in</Text>
                    </TouchableOpacity>
                </View>
            }
        >
            <PlainText
                required
                name='name'
                title='Name'
                onChange={onFormDataChange}
            />
            <PlainText
                required
                pattern='/[a-z0-9]{5,20}'
                name='username'
                title="Username"
                subtitle="this is permanent"
                onChange={(name, value, error) => onFormDataChange(name, value, error)}
            />
            <PhoneNumber
                required
                name='phoneNumber'
                title='Phone Number'
                onChange={(name, value, error) => onFormDataChange(name, value, error)}
            />
            <Button name='Next' size='medium' theme='primary' onPress={() => signUp()} loading={submitting} />
            {error &&
                <View style={styles.errorMessageContainer}>
                    <Text style={styles.errorMessage}>{error}</Text>
                </View>}

        </OnboardingFormWrapper >
    )
}

const styles = StyleSheet.create({
    loginContainer: {
        marginLeft: 5
    },
    login: {
        color: colors.primary
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

export default SignupScreen