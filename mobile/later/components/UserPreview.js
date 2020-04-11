import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import Color from '../assets/colors';
import Icon from '../components/Icon';

function UserPreview(props) {
    let user = props.user

    return (
        <View style={styles.container}>
            <View style={styles.imageContainer}>
                <View style={styles.thumb}></View>
            </View>
            <View style={styles.detailsContainer}>
                <Text style={styles.name}>{user.first_name} {user.last_name}</Text>
                <Text style={styles.username}>{user.username}</Text>
            </View>
            <View style={styles.addFriendContainer}>
                <Icon type='plus' size={25} />
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        height: 60,
        width: '100%',
        backgroundColor: Color.white
    },
    imageContainer: {
        width: 60,
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
    detailsContainer: {
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
    addFriendContainer: {
        height: '100%',
        padding: 5,
        marginRight: 10,
        justifyContent: 'center',
        alignItems: 'center',
    },
});

export default UserPreview