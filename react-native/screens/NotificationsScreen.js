import React, { useState, useEffect } from 'react';
import { StyleSheet, View, FlatList } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { FriendRequest } from '../components/user';
import { Divider } from '../components/common';

function NotificationsScreen() {
    const [friendRequests, setFriendRequests] = useState([])
    const [refreshing, setRefreshing] = useState(false)

    const loadFriendRequests = () => {
        setRefreshing(true)
        Network.GET('/friend-requests/pending')
            .then(fr => {
                setFriendRequests(fr)
                setRefreshing(false)
            })
            .catch(err => setRefreshing(false))
    }

    useEffect(() => {
        loadFriendRequests()
    }, [])

    const onAction = (friendRequest, action) => {
        // Preemptively update UI
        setFriendRequests(updateFriendRequests(friendRequests, friendRequest.id, action))

        respondToFriendRequest(friendRequest.id, action)
            .catch(() => {
                // Reset UI if request fails
                setFriendRequests(updateFriendRequests(friendRequests, friendRequest.id, undefined))
            })
    }

    const renderFriendRequest = ({ item }) => {
        <FriendRequest
            request={item}
            onAction={action => onAction(item, action)}
        />
    }

    return (
        <View style={styles.container}>
            <FlatList
                data={friendRequests}
                renderItem={renderFriendRequest}
                onRefresh={loadFriendRequests}
                refreshing={refreshing}
                ItemSeparatorComponent={<Divider />}
            />
        </View>
    );
}

const updateFriendRequests = (requests, requestId, status) => {
    return requests.map(request => requestId == request.id ? { ...request, ['status']: status } : request)
}

const respondToFriendRequest = (id, action) => {
    actions = {
        'Declined': 'decline',
        'Accepted': 'accept'
    }
    params = {
        id: id
    }
    return Network.PUT(`/friend-requests/${actions[action]}`, params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
});


export default NotificationsScreen;