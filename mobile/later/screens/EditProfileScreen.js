import React, { useState } from 'react';
import { StyleSheet, View, ScrollView, TouchableOpacity, Text, Alert } from 'react-native';
import { colors } from '../assets/colors';
import Header from '../components/Header';
import PlainText from '../components/forms/PlainText'
import Network from '../util/Network';
import Button from '../components/Button';

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

    const onFormDataChange = (name, value) => {
        setFormData({
            ...formData,
            [name]: value
        })
    }

    const submitForm = () => {
        Network.PUT("/users/update", formData)
            .then(() => navigation.navigate('Profile', { newUserData: formData }))
            .catch(err => Alert.alert(err))
    }

    console.log(formData)

    return (
        <View style={styles.container}>
            <Header name='Edit Profile' />
            <View style={styles.formContainer}>
                <ScrollView style={styles.scrollView} contentContainerStyle={styles.scrollViewContainer} >
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
                        <PlainText
                            name='email'
                            title='Email'
                            value={formData.email}
                            onChange={onFormDataChange}
                        />
                    </View>
                    <View onPress={submitForm} style={{ width: '100%', alignItems: 'flex-end', paddingTop: 10, }}>
                        <Button name='Submit' theme='light' size='medium' />
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
})

export default EditProfileScreen