import React, { useState, useContext } from 'react'
import { StyleSheet, View, Text } from 'react-native'
import { Button, BackIcon } from '../../components/common'
import { AuthContext } from '../../context'
import { colors } from '../../assets/colors'
import { PlainText, Password } from '../../components/forms'
import OnboardingFormWrapper from './OnboardingFormWrapper'
import { logIn } from '../../util/auth'

function LoginScreen({ navigation }) {
    const { signIn } = useContext(AuthContext)

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

    const submit = () => {
        setError(null)
        logIn(formData)
            .then(() => signIn())
            .catch(err => {
                console.log(err)
                setError('Invalid username or password')
            })
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
            <Button name='Submit' size='medium' theme='primary' onPress={() => submit()} />
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