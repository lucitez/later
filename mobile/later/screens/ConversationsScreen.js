import React, { useState, useEffect } from 'react';
import { StyleSheet, View } from 'react-native';
import { SearchBar } from '../components/common';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { UserGroup } from '../components/user';

function ConversationScreen() {
    const [conversations, setConversations] = useState([])
    const [search, setSearch] = useState(null)
    const [offset, setOffset] = useState(0)

    useEffect(() => {
        getConversations(search, offset)
            .then(conversations => setConversations(conversations))
            .catch(err => console.error(err))
    }, [])

    useEffect(() => {
        if (offset > 0)
            getConversations(search, offset)
                .then(conversations => setConversations(conversations))
                .catch(err => console.error(err))
    }, [offset])

    useEffect(() => {
        if (search)
            getConversations(search, offset)
                .then(conversations => setConversations(conversations))
                .catch(error => console.error(error))
    }, [search])

    return (
        <View style={styles.container}>
            <SearchBar onChange={search => setSearch(search)} />
            <UserGroup users={conversations} type='convo' />
        </View>
    );
}

const getConversations = (search, offset) => {
    params = {
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


export default ConversationScreen;