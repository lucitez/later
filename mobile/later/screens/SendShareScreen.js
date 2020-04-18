import React, { useState, useEffect } from 'react';
import { StyleSheet, View, KeyboardAvoidingView, Text } from 'react-native';
import { Header, SearchBar, Icon } from '../components/common';
import { ContentPreview } from '../components/content';
import { UserGroup } from '../components/user';
import Network from '../util/Network';
import { colors } from '../assets/colors';
import { userId } from '../util/constants';

function SendShareScreen({ navigation, route }) {

    let preview = route.params.contentPreview

    const [search, setSearch] = useState('')
    const [selectedFriends, setSelectedFriends] = useState({}) // map of userId to friendUser
    const [friends, setFriends] = useState([])

    useEffect(() => {
        searchFriends(search)
            .then(friends => setFriends(transformFriends(friends, selectedFriends)))
            .catch(error => console.error(error))
    }, [search])

    useEffect(() => {
        setFriends(transformFriends(friends, selectedFriends))
    }, [selectedFriends])

    const onSelectToggle = friendUser => {
        setSelectedFriends(updateSelectedFriends(friendUser, selectedFriends))
    }

    const onSend = () => {
        sendShares(preview.url, Object.keys(selectedFriends))
        navigation.navigate(route.params.previousScreen, { success: true })
    }

    const backIcon = (
        <Icon
            type='back'
            size={25}
            color={colors.white}
            onPress={() => navigation.pop()}
        />
    )

    return (
        <View style={styles.container}>
            <Header name='Share' leftIcon={backIcon} />
            <SearchBar
                onChange={value => setSearch(value)}
                iconName='friends'
                placeholder='Share with friends'
                onCancel={() => navigation.pop()}
            />
            <View style={styles.contentPreviewContainer}>
                <ContentPreview content={preview} />
            </View>
            <KeyboardAvoidingView
                behavior='padding'
                style={{ flex: 1 }}
            >
                <UserGroup users={friends} type='share' onSelectToggle={onSelectToggle} keyboardShouldPersistTaps='handled' />
                {
                    Object.values(selectedFriends).length > 0 ?
                        <View style={styles.selectedUsersContainer}>
                            <Text style={styles.selectedUsersText}>
                                {Object.values(selectedFriends).map(friend => friend.username).join(', ')}
                            </Text>
                            <View style={styles.sendIconContainer}>
                                <Icon type='next' size={30} color={colors.white} onPress={() => onSend()} />
                            </View>
                        </View>
                        : null
                }
            </KeyboardAvoidingView>
        </View >
    );
}

function transformFriends(friends, selectedFriends) {
    return friends.map(friend => (
        { ...friend, ['selected']: selectedFriends.hasOwnProperty(friend.userId) }
    ))
}

const updateSelectedFriends = (friendUser, selectedFriends) => {
    if (friendUser.selected) {
        let { [friendUser.userId]: userId, ...rest } = selectedFriends
        return rest
    } else {
        return { ...selectedFriends, [friendUser.userId]: friendUser }
    }
}

const searchFriends = search => {
    let params = {
        userId: userId,
        search: search,
    }
    let queryString = `/friends/for-user`
    return Network.GET(queryString, params)
}

const sendShares = (url, userIds) => {
    body = {
        senderUserId: userId,
        recipientUserIds: userIds,
        url: url
    }
    return Network.POST('/shares/new', body)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
    noPreviewContainer: {
        marginTop: 10,
        width: '100%',
        alignItems: 'center',
    },
    contentPreviewContainer: {
        backgroundColor: colors.white,
        margin: 5,
        borderRadius: 5,
    },
    selectedUsersContainer: {
        backgroundColor: colors.primary,
        padding: 10,
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center'
    },
    selectedUsersText: {
        fontSize: 20,
        fontWeight: '500',
        color: colors.white,
    }
});

export default SendShareScreen;