import React, { useState, useEffect } from 'react'
import { StyleSheet, View, Text, Alert } from 'react-native'
import Network from '../util/Network'
import { colors } from '../assets/colors'
import { UserDetails } from '../components/user'
import { Button, BackIcon } from '../components/common'
import { userId } from '../util/constants';

function UserScreen({ navigation, route }) {
    const [user, setUser] = useState(null)

    useEffect(() => {
        Network.GET("/users/profile-by-id", { requestUserId: userId, id: route.params.userId })
            .then(user => setUser(user))
            .catch(error => console.error(error))
    }, [])

    const onSendFriendRequest = () => {
        let updatedUser = { ...user, ['friendStatus']: 'pending' }
        setUser(updatedUser)

        sendFriendRequest(user.id)
            .then(friendRequest => setUser({ ...updatedUser, ['friendRequestId']: friendRequest.id }))
            .catch(() => setUser(user))
    }

    const onCancelFriendRequest = () => {
        let updatedUser = { ...user, ['friendStatus']: null }
        setUser(updatedUser)

        cancelFriendRequest(user.friendRequestId)
            .then(() => setUser({ ...updatedUser, ['friendRequestId']: null }))
            .catch(() => setUser(user))
    }

    const onDeleteFriend = () => {
        Alert.alert(
            "",
            "Are you sure you want to remove this user as a friend?",
            [
                {
                    text: 'Cancel',
                    onPress: () => console.log("confirmed"),
                    style: 'cancel'
                },
                {
                    text: 'Confirm',
                    onPress: () => onDeleteConfirmed(),
                    style: 'destructive'
                }
            ],
        )

    }

    const onDeleteConfirmed = () => {
        setUser({ ...user, ['friendStatus']: null })
        deleteFriend(user.id)
            .catch(() => setUser(user))
    }

    const renderButton = () => {
        switch (user.friendStatus) {
            case 'friends':
                return <Button theme='light' size='small' name='Friends' onPress={onDeleteFriend} />
            case 'pending':
                return <Button theme='light' size='small' name='Request Pending' onPress={onCancelFriendRequest} />
            default:
                return <Button theme='primary' size='small' name='Add Friend' onPress={onSendFriendRequest} />
        }
    }

    return (
        <View style={styles.container}>
            <View style={styles.topContainer}>
                <View style={styles.backButtonContainer}>
                    <BackIcon navigation={navigation} />
                </View>
                <View style={styles.topContentContainer}>
                    <View style={styles.userInfoContainer}>
                        {user && <UserDetails user={user} />}
                    </View>
                    <View style={styles.friendStatusContainer}>
                        {user && renderButton()}
                    </View>
                </View>
            </View>
        </View>
    )
}



const sendFriendRequest = friendUserId => {
    let params = {
        senderUserId: userId,
        recipientUserId: friendUserId,
    }

    return Network.POST("/friend-requests/send", params)
}

const cancelFriendRequest = id => {
    return Network.PUT("/friend-requests/delete", { id })
}

const deleteFriend = friendUserId => {
    let params = { userId, friendUserId }

    return Network.PUT("/friends/delete-by-user-id", params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray
    },
    topContainer: {
        height: '25%',
        justifyContent: 'flex-end',
        backgroundColor: colors.primary,
        paddingBottom: 5,
    },
    backButtonContainer: {
        paddingLeft: 10,
    },
    topContentContainer: {
        flexDirection: 'row',
    },
    userInfoContainer: {
        flexGrow: 1,
        flexBasis: 0
    },
    friendStatusContainer: {
        height: '100%',
        paddingRight: 10,
        paddingTop: 10,
    }
})

export default UserScreen