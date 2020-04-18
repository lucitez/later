import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { colors } from '../../assets/colors';

function FriendPreview(props) {
    let user = props.user

    return (
        <View style={styles.container}>
            <View style={styles.imageContainer}>
                <View style={styles.thumb}></View>
            </View>
            <View style={{ flexGrow: 1 }}>
                <View style={styles.userInfoContainer}>
                    <View style={styles.nameContainer}>
                        <Text style={styles.name}>{user.firstName} {user.lastName}</Text>
                    </View>
                    <View style={styles.usernameContainer}>
                        <Text style={styles.username}>@{user.username}</Text>
                    </View>
                </View>
                <View style={styles.messageContainer}>
                    <Text style={styles.message}>You shared a link</Text>
                </View>
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
    userInfoContainer: {
        flexDirection: 'row',
        paddingTop: 5,
    },
    nameContainer: {
        marginLeft: 5,
        marginRight: 5
    },
    usernameContainer: {
        marginTop: 1.5,
    },
    name: {
        fontSize: 16,
        fontWeight: '500'
    },
    username: {
        fontSize: 14,
        fontWeight: '300'
    },
    messageContainer: {
        marginLeft: 5,
    },
    message: {
        opacity: 0.75,
        fontStyle: 'italic'
    }
});

export default FriendPreview