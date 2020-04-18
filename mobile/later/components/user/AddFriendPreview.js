import React from 'react';
import { StyleSheet, Text, View, TouchableOpacity } from 'react-native';
import { colors } from '../../assets/colors';

function AddFriendPreview(props) {
    let user = props.user

    return (
        <View style={styles.container}>
            <View style={styles.imageContainer}>
                <View style={styles.thumb}></View>
            </View>
            <View style={styles.detailsContainer}>
                <Text style={styles.name}>{user.firstName} {user.lastName}</Text>
                <Text style={styles.username}>@{user.username}</Text>
            </View>
            <View style={styles.addFriendContainer}>
                {
                    user.pendingRequest ?
                        <View style={styles.requestSentContainer}>
                            <Text style={{ color: colors.white, fontSize: 12 }}>Request Sent</Text>
                        </View>
                        :
                        <TouchableOpacity onPress={() => props.onRequestSent()}>
                            <View style={styles.addFriendButton}>
                                <Text style={{ color: colors.green, fontSize: 12 }}>Send Request</Text>
                            </View>
                        </TouchableOpacity>
                }
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        height: 50,
        width: '100%',
        backgroundColor: colors.white
    },
    imageContainer: {
        width: 50,
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
    addFriendButton: {
        padding: 6,
        borderWidth: 1.5,
        borderRadius: 5,
        borderColor: colors.green,
        justifyContent: 'center'
    },
    requestSentContainer: {
        padding: 7,
        borderRadius: 5,
        backgroundColor: colors.green,
        justifyContent: 'center'
    },

});

export default AddFriendPreview