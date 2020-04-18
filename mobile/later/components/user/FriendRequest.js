import React from 'react'
import { StyleSheet, View, Text } from 'react-native'
import { colors } from '../../assets/colors'
import { Button } from '../common'

function FriendRequest({ request, onAccept, onDecline }) {
    return (
        <View style={styles.container}>
            <View style={styles.imageContainer}>
                <View style={styles.thumb}></View>
            </View>
            <View style={styles.userInfoContainer}>
                <View style={styles.nameContainer}>
                    <Text style={styles.name}>{request.firstName} {request.lastName}</Text>
                </View>
                <View style={styles.usernameContainer}>
                    <Text style={styles.username}>@{request.username}</Text>
                </View>
            </View>
            <View style={styles.actionContainer}>
                {
                    request.status ?
                        <Text>{request.status}</Text>
                        :
                        <>
                            <View style={styles.buttonContainer}>
                                <Button theme='light' size='small' name='Decline' onPress={() => onDecline()} />
                            </View>
                            <View style={styles.buttonContainer}>
                                <Button theme='primary' size='small' name='Accept' onPress={() => onAccept()} />
                            </View>
                        </>
                }

            </View>
        </View>
    )
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
        padding: 5,
        justifyContent: 'center',
        flexGrow: 1,
    },
    nameContainer: {
        marginLeft: 5,
    },
    usernameContainer: {
        marginLeft: 5,
    },
    name: {
        fontSize: 16,
        fontWeight: '500'
    },
    username: {
        fontSize: 14,
        fontWeight: '300'
    },
    actionContainer: {
        flexDirection: 'row',
        width: '40%',
        justifyContent: 'flex-end',
        alignItems: 'center'
    },
    buttonContainer: {
        flexGrow: 1,
        padding: 3,
        justifyContent: 'center'
    }
})

export default FriendRequest
