import React from 'react';
import { View, ScrollView } from 'react-native';
import UserPreview from './UserPreview'
import Divider from './Divider'

function UserGroup(props) {
    return (
        <ScrollView>
            {
                props.users.map((user, index) => (
                    <View key={index}>
                        <UserPreview user={user} key={user.id} onRequestSent={() => props.onRequestSent(user.id)} />
                        {index < props.users.length - 1 ? <Divider key={props.index} /> : null}
                    </View>
                ))
            }
        </ScrollView>
    );
}

export default UserGroup