import React, { useState, useContext } from 'react'
import { StyleSheet, View, Text } from 'react-native'
import OnboardingFormWrapper from './OnboardingFormWrapper'
import { Password } from '../../components/forms'
import { Button, BackIcon } from '../../components/common';
import { AuthContext } from '../../context'
import { colors } from '../../assets/colors';
import { signUp } from '../../util/auth'

export default function CreatePasswordScreen({ navigation, route }) {
    const { signIn } = useContext(AuthContext)

    const [submitting, setSubmitting] = useState(false)
    const [formData, setFormData] = useState(route.params.formData)
    const [error, setError] = useState(null)

    const onSubmit = () => {
        setSubmitting(true)
        setError(null)

        if (formData.password.error) {
            setError(formData.password.error)
            return
        }

        signUp(Object.fromEntries(Object.entries(formData).map(([field, data]) => [field, data.value])))
            .then(() => {
                setSubmitting(false)
                signIn()
            })
            .catch(err => {
                setError(err)
                setSubmitting(false)
            })
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
            <Button name='Submit' theme='primary' size='medium' onPress={onSubmit} loading={submitting} />
            {error &&
                <View style={styles.errorMessageContainer}>
                    <Text style={styles.errorMessage}>{error}</Text>
                </View>}
            {submitting &&
                <View style={styles.errorMessageContainer}>
                    <Text style={styles.contextMessage}>Creating your account</Text>
                </View>}
        </OnboardingFormWrapper>
    )
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
    },
    contextMessage: {
        color: colors.black,
        fontWeight: '300',
    },
})

