import React, { useState, useEffect } from 'react';
import { StyleSheet, View } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { FriendRequest } from '../components/user';
import { userId } from '../util/constants';
import { Divider } from '../components/common';

function NotificationsScreen() {
    const [friendRequests, setFriendRequests] = useState([])

    useEffect(() => {
        getFriendRequests()
            .then(friendRequests => setFriendRequests(friendRequests))
            .catch(err => console.error(err))
    }, [])

    const onAccept = (friendRequest) => {
        // Preemptively update UI
        setFriendRequests(updateFriendRequests(friendRequests, friendRequest.id, 'accepted'))

        respondToFriendRequest(friendRequest.id, 'accept')
            .catch(() => {
                // Reset UI if request fails
                setFriendRequests(updateFriendRequests(friendRequests, friendRequest.id, undefined))
            })
    }

    const onDecline = (friendRequest) => {
        // Preemptively update UI
        setFriendRequests(updateFriendRequests(friendRequests, friendRequest.id, 'declinled'))

        respondToFriendRequest(friendRequest.id, 'decline')
            .catch(() => {
                // Reset UI if request fails
                setFriendRequests(updateFriendRequests(friendRequests, friendRequest.id, undefined))
            })
    }

    return (
        <View style={styles.container}>
            {friendRequests.map((friendRequest, index) =>
                <View key={index}>
                    <FriendRequest
                        request={friendRequest}
                        onAccept={() => onAccept(friendRequest)}
                        onDecline={() => onDecline(friendRequest)}
                    />
                    {index < friendRequests.length - 1 && <Divider />}
                </View>
            )}
        </View>
    );
}

const updateFriendRequests = (requests, requestId, status) => {
    return requests.map(request => requestId == request.id ? { ...request, ['status']: status } : request)
}

const getFriendRequests = () => {
    params = {
        userId: userId
    }
    return Network.GET('/friend-requests/pending', params)
}

const respondToFriendRequest = (id, action) => {
    params = {
        id: id
    }
    return Network.PUT(`/friend-requests/${action}`, params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
});


export default NotificationsScreen;