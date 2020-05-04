import React, { useState, useEffect } from 'react';
import { StyleSheet, View, FlatList } from 'react-native';
import { Divider, SearchBar, Icon } from '../components/common';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { ChatPreview } from '../components/chat/index'
import { TouchableOpacity } from 'react-native-gesture-handler';

const LIMIT = 20

export default function ChatsScreen({ navigation, route }) {
    const [chats, setChats] = useState([])
    const [refreshing, setRefreshing] = useState(false)
    const [limitReached, setLimitReached] = useState(false)
    const [search, setSearch] = useState('')

    const replaceChats = chats => {
        if (chats.length < LIMIT) {
            setLimitReached(true)
        }
        setChats(chats)
    }

    const appendChats = nextPage => {
        if (nextPage.length < LIMIT) {
            setLimitReached(true)
        }
        setChats(chats.concat(nextPage))
    }

    const updateChats = (offset = 0, chatUpdateFunc) => {
        setRefreshing(true)
        getChats(offset)
            .then(chats => chatUpdateFunc(chats))
            .catch(err => console.error(err))
            .finally(() => setRefreshing(false))
    }

    useEffect(() => {
        updateChats(0, replaceChats)
    }, [])

    const onEndReached = offset => {
        if (!limitReached) {
            updateChats(offset, appendChats)
        }
    }

    const renderChat = (chat, navigation) => (
        <TouchableOpacity onPress={() => {
            setChats(chats.map(c => c.chatId == chat.chatId ? { ...c, hasUnread: false } : c))
            navigation.navigate('Chat', { chatDetails: chat })
        }}>
            <ChatPreview chat={chat} />
        </TouchableOpacity>
    )

    // todo start chat with any friend
    return (
        <View style={styles.container}>
            <SearchBar onChange={text => setSearch(text)} />
            <FlatList
                onRefresh={() => updateChats(0, replaceChats)}
                onEndReached={onEndReached}
                refreshing={refreshing}
                keyExtractor={item => item.chatId}
                data={chats}
                renderItem={({ item }) => renderChat(item, navigation)}
                ItemSeparatorComponent={Divider}
            />
        </View>
    );
}



const getChats = (offset) => {
    let params = {
        offset: offset,
        limit: LIMIT,
    }
    let queryString = `/chats/for-user`

    return Network.GET(queryString, params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
});
