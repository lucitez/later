import React from 'react';
import { View, ScrollView, TouchableHighlight } from 'react-native';
import AddFriendPreview from './AddFriendPreview'
import FriendPreview from './FriendPreview'
import ShareWithFriendPreview from './ShareWithFriendPreview'
import Divider from './Divider'

function UserGroup(props) {
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
        case 'share':
            return (
                <TouchableHighlight key={user.id} onPress={() => props.onSelectToggle(user)}>
                    <ShareWithFriendPreview user={user} />
                </TouchableHighlight>
            )
        default:
            null
    }
}

export default UserGroup