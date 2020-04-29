import React from 'react';
import { StyleSheet, View } from 'react-native';
import { colors } from '../../assets/colors';
import UserDetailsPreview from './UserDetailsPreview';
import { TouchableOpacity } from 'react-native-gesture-handler';

function UserPreview({ user, onPress }) {
    return (
        <TouchableOpacity style={styles.container} onPress={() => onPress(user.id)}>
            <View style={styles.userDetailsContainer}>
                <UserDetailsPreview user={user} />
            </View>
        </TouchableOpacity>
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