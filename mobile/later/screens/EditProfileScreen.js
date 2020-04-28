import React, { useState } from 'react';
import { StyleSheet, View, ScrollView, Text, SafeAreaView, Keyboard } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { PlainText, Email, PhoneNumber } from '../components/forms';
import { Header, Button, BackIcon } from '../components/common';
import { TouchableWithoutFeedback } from 'react-native-gesture-handler';

function EditProfileScreen({ navigation, route }) {
    const user = route.params.user

    const [formData, setFormData] = useState({
        id: { value: user.id, },
        name: { value: user.name },
        email: { value: user.email },
        phoneNumber: { value: user.phoneNumber }
    })

    const [validationErrors, setValidationErrors] = useState({})
    const [error, setError] = useState(null)
    const [submitting, setSubmitting] = useState(false)

    const onFormDataChange = (name, value, error) => {
        setFormData({
            ...formData,
            [name]: {
                value: value,
                error: error
            }
        })
    }

    const submitForm = () => {
        setError(null)
        setSubmitting(true)

        let body = {}

        for (let [key, field] of Object.entries(formData)) {
            if (field.error) {
                setError(field.error)
                setSubmitting(false)
                return
            }
            body[key] = field.value
        }

        console.log(body)

        Network.PUT("/users/update", body)
            .then(() => {
                navigation.navigate('Profile', { newUserData: body })
            })
            .catch(err => setError(err))
            .finally(() => setSubmitting(false))
    }

    return (
        <SafeAreaView style={styles.container}>
            <Header title='Edit Profile' leftIcon={<BackIcon navigation={navigation} />} />
            <ScrollView style={styles.scrollStyle} containerStyle={styles.contentContainer} keyboardShouldPersistTaps='handled'>
                <View style={styles.formContainer}>
                    <View style={styles.nameFormContainer}>
                        <PlainText
                            theme='light'
                            name='name'
                            title='First Name'
                            value={formData.name.value}
                            onChange={onFormDataChange}
                        />
                    </View>
                    <View style={styles.emailFormContainer}>
                        <Email
                            theme='light'
                            name='email'
                            title='Email'
                            value={formData.email.value}
                            onChange={onFormDataChange}
                        />
                    </View>
                    <View style={styles.phoneFormContainer}>
                        <PhoneNumber
                            theme='light'
                            name='phoneNumber'
                            title='Phone Number'
                            value={formData.phoneNumber.value}
                            onChange={onFormDataChange}
                        />
                    </View>
                    <View style={styles.bottomContainer}>
                        <View style={styles.submitButtonContainer}>
                            <Button name={submitting ? 'Submitting...' : 'Submit'} theme='light' size='medium' onPress={() => {
                                setSubmitting(true)
                                submitForm()
                            }} />
                        </View>
                        {error &&
                            <View style={styles.errorMessageContainer}>
                                <Text style={styles.errorMessage}>{error}</Text>
                            </View>}
                    </View>
                </View>
            </ScrollView>

        </SafeAreaView >
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.primary,
        justifyContent: 'flex-start',
    },
    scrollStyle: {
        paddingTop: 40,
        borderTopWidth: 1,
        borderTopColor: colors.white,
    },
    contentContainer: {
        flexGrow: 1,
        alignItems: 'center',
    },
    formContainer: {
        minWidth: '100%',
        maxWidth: '100%',
        paddingLeft: '15%',
        paddingRight: '15%',
        paddingBottom: 20,
        alignItems: 'flex-start'
    },
    emailFormContainer: {
        minWidth: '75%',
    },
    nameFormContainer: {
        minWidth: '75%',
    },
    phoneFormContainer: {
        minWidth: '50%',
    },
    bottomContainer: {
        minWidth: '100%',
        alignItems: 'flex-end',
        marginTop: 25,
    },
    submitButtonContainer: {
        width: '40%'
    },
    errorMessageContainer: {
        paddingTop: 10
    },
    errorMessage: {
        color: colors.white,
        fontSize: 18
    }
})

export default EditProfileScreen