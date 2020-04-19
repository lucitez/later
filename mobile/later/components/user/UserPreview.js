import React from 'react';
import { StyleSheet, View } from 'react-native';
import { colors } from '../../assets/colors';
import UserDetails from './UserDetails';

function UserPreview({ user }) {
    return (
        <View style={styles.container}>
            <View style={styles.userDetailsContainer}>
                <UserDetails user={user} />
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        height: 60,
        width: '100%',
        backgroundColor: colors.white
    },
    userDetailsContainer: {
        flexGrow: 1,
    }
});

export default UserPreview