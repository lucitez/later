import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Text, ActivityIndicator, TouchableHighlight } from 'react-native';
import { SearchBar, Icon } from '../common';
import { ShareWithFriendPreview } from '../user'
import Network from '../../util/Network';
import { colors } from '../../assets/colors';
import { FlatList } from 'react-native-gesture-handler';
import ContentPreview from '../content/ContentPreview';

export default function ShareSendSelector({ onSend, contentPreview, onSearchCancel }) {
    const [search, setSearch] = useState('')

    const [loading, setLoading] = useState(false)
    const [submitting, setSubmitting] = useState(false)

    const [friends, setFriends] = useState([])
    const [selectedFriends, setSelectedFriends] = useState({}) // map of userId to friendUser

    useEffect(() => {
        setLoading(true)
        searchFriends(search)
            .then(friends => {
                setFriends(transformFriends(friends, selectedFriends))
                setLoading(false)
            })
            .catch(error => {
                console.error(error)
                setLoading(false)
            })
    }, [search])

    const onSelectToggle = friendUser => {
        let newSelectedFriends = updatedSelectedFriends(friendUser, selectedFriends)
        setSelectedFriends(newSelectedFriends)
        setFriends(transformFriends(friends, newSelectedFriends))
    }

    const _renderContentPreview = () => (
        <View style={styles.contentPreviewContainer}>
            <ContentPreview content={contentPreview} onDotPress={() => null} />
        </View>
    )

    const _renderFriend = ({ item }) => (
        <TouchableHighlight activeOpacity={0.75} underlayColor='#FFFFFF' onPress={() => onSelectToggle(item)}>
            <ShareWithFriendPreview user={item} />
        </TouchableHighlight>
    )

    const _renderFooter = () => (
        <View style={styles.selectedUsersContainer}>
            <Text style={styles.selectedUsersText}>
                {Object.values(selectedFriends).map(friend => friend.username).join(', ')}
            </Text>
            <View>
                {submitting ? <ActivityIndicator size='small' color={colors.white} />
                    : <Icon
                        type='next'
                        size={30}
                        color={colors.white}
                        onPress={() => {
                            setSubmitting(true)
                            onSend(Object.keys(selectedFriends), () => setSubmitting(false))
                        }}
                    />}
            </View>
        </View>
    )

    return (
        <View style={styles.container}>
            <SearchBar
                onChange={value => setSearch(value)}
                iconName='friends'
                placeholder='Share with friends'
                onCancel={() => {
                    if (onSearchCancel) onSearchCancel()
                }}
            />

            <View style={styles.contentContainer}>
                {contentPreview && _renderContentPreview()}
                <FlatList
                    data={friends}
                    refreshing={loading}
                    renderItem={_renderFriend}
                    keyboardShouldPersistTaps='handled'
                />
                {Object.values(selectedFriends).length > 0 && _renderFooter()}
            </View>
        </View>
    );
}

function transformFriends(friends, selectedFriends) {
    return friends.map(friend => (
        { ...friend, ['selected']: selectedFriends.hasOwnProperty(friend.userId) }
    ))
}

const updatedSelectedFriends = (friendUser, selectedFriends) => {
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

const styles = StyleSheet.create({
    container: {
        flexBasis: 0,
        flexGrow: 1,
        backgroundColor: colors.lightGray,
    },
    contentContainer: {
        flexBasis: 0,
        flexGrow: 1,
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
        alignItems: 'center',
        borderBottomWidth: 0.5,
        borderBottomColor: colors.white
    },
    selectedUsersText: {
        fontSize: 20,
        fontWeight: '500',
        color: colors.white,
    }
});