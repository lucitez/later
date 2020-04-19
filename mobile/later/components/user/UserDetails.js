import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { colors } from '../../assets/colors';

function UserDetails({ user }) {
    return (
        <View style={styles.container}>
            <View style={styles.imageContainer}>
                <View style={styles.thumb}></View>
            </View>

            <View style={styles.userInfoContainer}>
                <Text style={styles.name}>{user.firstName} {user.lastName}</Text>
                <Text style={styles.username}>@{user.username}</Text>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        flex: 1,
        backgroundColor: colors.white
    },
    imageContainer: {
        aspectRatio: 1,
        padding: 5,
        justifyContent: 'center',
        alignItems: 'center',
    },
    thumb: {
        height: '100%',
        width: '100%',
        borderRadius: 100,
        backgroundColor: 'coral',
    },
    userInfoContainer: {
        height: '100%',
        flexGrow: 1,
        padding: 5,
        justifyContent: 'center',
    },
    name: {
        fontSize: 16,
        fontWeight: '500'
    },
    username: {
        fontSize: 14,
        fontWeight: '300'
    },
});

export default UserDetails