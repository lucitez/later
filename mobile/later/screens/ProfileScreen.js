import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import Header from '../components/Header';
import { userId } from '../util/constants';
import Network from '../util/Network';
import { colors } from '../assets/colors';

function ProfileScreen() {
    const [user, setUser] = useState(null)

    useEffect(() => {
        getUser()
            .then(u => setUser(u))
            .catch(err => console.error(err))
    })

    return (
        <View style={styles.container}>
            <Header name="Profile" />
            <Text>todo</Text>
        </View>
    )
}

const getUser = () => {
    params = {
        userId: userId
    }
    Network.GET('/users/by-id', params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
})

export default ProfileScreen