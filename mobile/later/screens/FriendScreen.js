import React, { useState, useEffect } from 'react';
import { StyleSheet, View } from 'react-native';
import { Icon, SearchBar } from '../components/common';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { UserGroup } from '../components/user';

function FriendScreen({ navigation }) {
    const [friends, setFriends] = useState([])
    const [search, setSearch] = useState(null)
    const [offset, setOffset] = useState(0)

    useEffect(() => {
        getFriends(search, offset)
            .then(nextPage => setFriends(friends.concat(nextPage)))
            .catch(error => console.error(error))
    }, [])

    useEffect(() => {
        if (offset > 0) {
            getFriends(search, offset)
                .then(nextPage => setFriends(friends.concat(nextPage)))
                .catch(error => console.error(error))
        }

    }, [offset])

    useEffect(() => {
        if (search)
            getFriends(userId, search, offset)
                .then(friends => setFriends(friends))
                .catch(error => console.error(error))
    }, [search])

    return (
        <View style={styles.container}>
            <SearchBar onChange={search => setSearch(search)} />
            <UserGroup users={friends} type='friend' />
        </View>
    );
}

const getFriends = (search, offset) => {
    params = {
        userId: userId,
        search: search,
        offset: offset,
        limit: 10
    }
    let queryString = `/friends/for-user`

    return Network.GET(queryString, params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
    searchContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 10
    },
});


export default FriendScreen;