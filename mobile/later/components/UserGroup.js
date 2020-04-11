import React from 'react';
import { View, ScrollView, Text } from 'react-native';
import AddFriendPreview from './AddFriendPreview'
import FriendPreview from './FriendPreview'
import Divider from './Divider'

function UserGroup(props) {
    console.log(props)
    return (
        <ScrollView>
            {
                props.users.map((user, index) => (
                    <View key={index}>
                        {Preview(props, user)}
                        {index < props.users.length - 1 ? <Divider key={props.index} /> : null}
                    </View>
                ))
            }
        </ScrollView>
    );
}

function Preview(props, user) {
    switch (props.type) {
        case 'add_friend':
            return <AddFriendPreview user={user} key={user.id} onRequestSent={() => props.onRequestSent(user.id)} />
        case 'friend':
            return <FriendPreview user={user} key={user.id} />
        default:
            null
    }
}

export default UserGroup