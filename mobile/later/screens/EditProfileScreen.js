import React, { useState, useEffect } from 'react';
import { StyleSheet, View, ScrollView, TouchableOpacity, Text, Alert, Keyboard } from 'react-native';
import { colors } from '../assets/colors';
import Header from '../components/Header';
import PlainText from '../components/forms/PlainText'
import Network from '../util/Network';
import Button from '../components/Button';
import Email from '../components/forms/Email';
import Icon from '../components/Icon';
import BackIcon from '../components/BackIcon';

function EditProfileScreen({ navigation, route }) {
    const user = route.params.user

    const [formData, setFormData] = useState({
        id: user.id,
        username: user.username,
        firstName: user.firstName,
        lastName: user.lastName,
        email: user.email,
        phoneNumber: user.phoneNumber
    })

    const [validationErrors, setValidationErrors] = useState({})
    const [validationError, setValidationError] = useState(null)
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
        setValidationError(null)
        for (let [key, valid] of Object.entries(validationErrors)) {
            if (!valid) {
                setValidationError(`Please provide a valid ${key}`)
                return false
            }
        }
        return true
    }

    const submitForm = () => {
        if (validate()) {
            Network.PUT("/users/update", formData)
                .then(() => navigation.navigate('Profile', { newUserData: formData }))
                .catch(err => Alert.alert(err))
                .finally(() => setSubmitting(false))
        } else {
            setSubmitting(false)
        }
    }

    console.log(submitting)

    return (
        <View style={styles.container}>
            <Header name='Edit Profile' leftIcon={<BackIcon navigation={navigation} />} />
            <View style={styles.formContainer}>
                <ScrollView style={styles.scrollView} contentContainerStyle={styles.scrollViewContainer} keyboardShouldPersistTaps='handled' >
                    <View style={styles.usernameFormContainer}>
                        <PlainText
                            name='username'
                            title='Username'
                            value={formData.username}
                            onChange={onFormDataChange}
                        />
                    </View>
                    <View style={styles.nameFormContainer}>
                        <View style={styles.firstNameFormContainer}>
                            <PlainText
                                name='firstName'
                                title='First Name'
                                value={formData.firstName}
                                onChange={onFormDataChange}
                            />
                        </View>
                        <View style={styles.firstNameFormContainer}>
                            <PlainText
                                name='lastName'
                                title='Last Name'
                                value={formData.lastName}
                                onChange={onFormDataChange}
                            />
                        </View>
                    </View>
                    <View style={styles.emailFormContainer}>
                        <Email
                            name='email'
                            title='Email'
                            value={formData.email}
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
                        {validationError &&
                            <View style={styles.errorMessageContainer}>
                                <Text style={styles.errorMessage}>{validationError}</Text>
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
    firstNameFormContainer: {
        flex: 1,
    },
    lastNameFormContainer: {
        flex: 1,
    },
    bottomContainer: {
        width: '100%',
        alignItems: 'flex-end',
        marginTop: 25,
    },
    submitButtonContainer: {
        width: '33%'
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