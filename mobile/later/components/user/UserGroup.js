import React from 'react';
import { View, ScrollView, TouchableHighlight } from 'react-native';
import { Divider } from '../common';
import UserPreview from './UserPreview'
import FriendPreview from './FriendPreview'
import ShareWithFriendPreview from './ShareWithFriendPreview'
import ConversationPreview from './ConversationPreview'

function UserGroup(props) {
    return (
        <ScrollView keyboardShouldPersistTaps={props.keyboardShouldPersistTaps}>
            {
                props.users.map((user, index) => (
                    <View key={index}>
                        {Preview(props, user)}
                        {index < props.users.length - 1 ? <Divider /> : null}
                    </View>
                ))
            }
        </ScrollView>
    );
}

function Preview(props, user) {
    switch (props.type) {
        case 'user':
            return <UserPreview user={user} onPress={() => null} />
        case 'friend':
            return <FriendPreview user={user} />
        case 'convo':
            return <ConversationPreview user={user} />
        case 'share':
            return (
                <TouchableHighlight onPress={() => props.onSelectToggle(user)}>
                    <ShareWithFriendPreview user={user} />
                </TouchableHighlight>
            )
        default:
            null
    }
}

export default UserGroup