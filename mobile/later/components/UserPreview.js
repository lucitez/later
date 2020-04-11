import React, { useState } from 'react';
import { StyleSheet, Text, View, TouchableOpacity } from 'react-native';
import Colors from '../assets/colors';

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
                {
                    user.pending_request ?
                        <View style={styles.requestSentContainer}><Text style={{ color: Colors.white }}>Request Sent</Text></View>
                        :
                        <TouchableOpacity onPress={() => props.onRequestSent()}>
                            <View style={styles.addFriendButton}><Text style={{ color: Colors.green }}>Send Request</Text></View>
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
        backgroundColor: Colors.white
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
        padding: 7,
        borderWidth: 2,
        borderRadius: 5,
        borderColor: Colors.green,
        justifyContent: 'center'
    },
    requestSentContainer: {
        padding: 7,
        borderRadius: 5,
        backgroundColor: Colors.green,
        justifyContent: 'center'
    },

});

export default UserPreview