import React, { useState } from 'react';
import { StyleSheet, View, ScrollView, Text } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { PlainText, Email, PhoneNumber } from '../components/forms';
import { Header, Button, BackIcon } from '../components/common';

function EditProfileScreen({ navigation, route }) {
    const user = route.params.user

    const [formData, setFormData] = useState({
        id: user.id,
        name: user.name,
        email: user.email,
        phoneNumber: user.phoneNumber
    })

    const [validationErrors, setValidationErrors] = useState({})
    const [error, setError] = useState(null)
    const [submitting, setSubmitting] = useState(false)

    const onFormDataChange = (name, value, valid) => {
        setFormData({
            ...formData,
            [name]: value
        })
        setValidationErrors({
            ...validationErrors,
            [name]: valid
        })
    }

    const validate = () => {
        setError(null)
        for (let [key, valid] of Object.entries(validationErrors)) {
            if (!valid) {
                setError(`Please provide a valid ${key}`)
                return false
            }
        }
        return true
    }

    const submitForm = () => {
        if (validate()) {
            Network.PUT("/users/update", formData)
                .then(() => navigation.navigate('Profile', { newUserData: formData }))
                .catch(err => setError(err))
                .finally(() => setSubmitting(false))
        } else {
            setSubmitting(false)
        }
    }

    return (
        <View style={styles.container}>
            <Header title='Edit Profile' leftIcon={<BackIcon navigation={navigation} />} />
            <View style={styles.formContainer}>
                <ScrollView style={styles.scrollView} contentContainerStyle={styles.scrollViewContainer} keyboardShouldPersistTaps='handled' >
                    <View style={styles.nameFormContainer}>
                        <View style={styles.nameFormContainer}>
                            <PlainText
                                theme='light'
                                name='name'
                                title='First Name'
                                value={formData.name}
                                onChange={onFormDataChange}
                            />
                        </View>
                    </View>
                    <View style={styles.emailFormContainer}>
                        <Email
                            theme='light'
                            name='email'
                            title='Email'
                            value={formData.email}
                            onChange={onFormDataChange}
                        />
                    </View>
                    <View style={styles.phoneFormContainer}>
                        <PhoneNumber
                            theme='light'
                            name='phoneNumber'
                            title='Phone Number'
                            value={formData.phoneNumber}
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
                </ScrollView>
            </View>
        </View >
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.primary,
        alignItems: 'center',
    },
    formContainer: {
        width: '100%',
        alignItems: 'center',
        justifyContent: 'flex-start'
    },
    scrollView: {
        width: '80%',
    },
    scrollViewContainer: {
        justifyContent: 'flex-start',
        alignItems: 'flex-start',
        paddingTop: 40,
        paddingBottom: 20,
    },
    usernameFormContainer: {
        width: '50%',
    },
    emailFormContainer: {
        minWidth: '75%',
    },
    nameFormContainer: {
        flexDirection: 'row',
        width: '100%',
    },
    nameFormContainer: {
        flex: 1,
    },
    phoneFormContainer: {
        width: '50%'
    },
    bottomContainer: {
        width: '100%',
        alignItems: 'flex-end',
        marginTop: 25,
    },
    submitButtonContainer: {
        width: '35%'
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