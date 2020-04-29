import React, { useState, useEffect } from 'react'
import { StyleSheet, View, Text, Alert } from 'react-native'
import OnboardingFormWrapper from './OnboardingFormWrapper'
import { PlainText } from '../../components/forms'
import Network from '../../util/Network';
import { Button, BackIcon } from '../../components/common';
import { TouchableOpacity } from 'react-native-gesture-handler';
import { colors } from '../../assets/colors';

export default function SMSCodeScreen({ navigation, route }) {
    let formData = route.params.formData
    let phoneNumber = formData.phoneNumber.value

    const [confirmationCode, setConfirmationCode] = useState('000000')
    const [confirmationCodeInput, setConfirmationCodeInput] = useState('')
    const [error, setError] = useState(null)

    const fetchCode = () => {
        // Network.POST('/auth/sms-confirmation', { phoneNumber: phoneNumber })
        //     .then(confirmationCode => setConfirmationCode(confirmationCode))
        //     .catch(err => console.error(err))
    }

    useEffect(() => {
        fetchCode()
    }, [])

    const onResendPress = () => {
        Alert.alert(
            "Confirm Resend SMS",
            "This action will invalidate the previously sent code.",
            [
                {
                    text: 'Cancel',
                    style: 'cancel'
                },
                {
                    text: 'Confirm',
                    onPress: () => fetchCode(),
                    style: 'default'
                }
            ],
        )
    }

    const onNext = () => {
        setError(null)
        if (confirmationCodeInput != confirmationCode) {
            setError('Entered code is incorrect.')
            return
        }

        navigation.navigate('Password', { formData })
    }

    return (
        <OnboardingFormWrapper
            title='Verify Phone Number'
            description='Input the verification code we sent to your phone'
            leftIcon={<BackIcon navigation={navigation} color={colors.primary} />}
        >
            <PlainText
                inputProps={{ keyboardType: 'phone-pad', autoFocus: true }}
                name='confirmationCode'
                title='Confirmation Code'
                onChange={(_, value) => setConfirmationCodeInput(value)}
            />
            <Button name='Next' theme='primary' size='medium' onPress={onNext} />
            {error &&
                <View style={styles.errorMessageContainer}>
                    <Text style={styles.errorMessage}>{error}</Text>
                </View>}
            <View style={styles.resendContainer}>
                <TouchableOpacity onPress={onResendPress}>
                    <Text style={styles.resend}>Resend SMS</Text>
                </TouchableOpacity>
            </View>
        </OnboardingFormWrapper>
    )
}

const styles = StyleSheet.create({
    resendContainer: {
        marginTop: 5,
        paddingTop: 5,
        alignItems: 'flex-end'
    },
    resend: {
        color: colors.primary,
        fontWeight: '500'
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

