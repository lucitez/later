import React, { useState, useEffect } from 'react';
import { StyleSheet, View, KeyboardAvoidingView, Text, SafeAreaView, ActivityIndicator } from 'react-native';
import { Header, SearchBar, Icon } from '../components/common';
import { ContentPreview } from '../components/content';
import { UserGroup } from '../components/user';
import Network from '../util/Network';
import { colors } from '../assets/colors';

function SendShareScreen({ navigation, route }) {

    let preview = route.params.contentPreview

    const [search, setSearch] = useState('')
    const [selectedFriends, setSelectedFriends] = useState({}) // map of userId to friendUser
    const [submitting, setSubmitting] = useState(false)
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
        setSubmitting(true)
        sendShares(preview.url, Object.keys(selectedFriends))
            .then(() => navigation.navigate(route.params.previousScreen, { success: true }))
            .catch(err => {
                console.error(err)
                setSubmitting(false)
            })
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
        <SafeAreaView style={styles.container}>
            <Header title='Share' leftIcon={backIcon} />
            <SearchBar
                onChange={value => setSearch(value)}
                iconName='friends'
                placeholder='Share with friends'
                onCancel={() => navigation.pop()}
            />
            <KeyboardAvoidingView behavior='padding' style={styles.contentContainer}>
                <View style={styles.contentPreviewContainer}>
                    <ContentPreview content={preview} />
                </View>
                <UserGroup users={friends} type='share' onSelectToggle={onSelectToggle} keyboardShouldPersistTaps='handled' />
                {
                    Object.values(selectedFriends).length > 0 ?
                        <View style={styles.selectedUsersContainer}>
                            <Text style={styles.selectedUsersText}>
                                {Object.values(selectedFriends).map(friend => friend.username).join(', ')}
                            </Text>
                            <View style={styles.sendIconContainer}>
                                {submitting
                                    ? <ActivityIndicator size='small' color={colors.white} />
                                    : <Icon type='next' size={30} color={colors.white} onPress={() => onSend()} />}
                            </View>
                        </View>
                        : null
                }
            </KeyboardAvoidingView>
        </SafeAreaView >
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
        search: search,
    }
    let queryString = `/friends/for-user`
    return Network.GET(queryString, params)
}

const sendShares = (url, userIds) => {
    let body = {
        recipientUserIds: userIds,
        url: url
    }
    return Network.POST('/shares/new', body)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.primary,
    },
    contentContainer: {
        flexGrow: 1,
        backgroundColor: colors.lightGray,
    },
    contentPreviewContainer: {
        backgroundColor: colors.white,
        margin: 5,
        borderRadius: 5,
    },
    selectedUsersContainer: {
        backgroundColor: colors.primary,
        height: 50,
        paddingRight: 10,
        paddingLeft: 10,
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