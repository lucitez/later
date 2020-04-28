import React, { useEffect, useState } from 'react'
import { StyleSheet, View, Text, SafeAreaView, FlatList, Image, Keyboard } from 'react-native'
import { colors } from '../assets/colors'
import { Header, BackIcon } from '../components/common'
import Network from '../util/Network'
import { useSelector } from 'react-redux'
import { MessageContainer, ChatInput } from '../components/chat'
import { v4 as uuidv4 } from 'uuid';

const LIMIT = 15

export default function ChatDisplayScreen({ navigation, route }) {
    let chatDetails = route.params.chatDetails
    let userId = useSelector(state => state.auth.userId)

    const [offset, setOffset] = useState(0)
    const [messages, setMessages] = useState([])
    const [limitReached, setLimitReached] = useState(false)

    const updateMessages = (messageUpdateFunc) => {
        loadMessages(chatDetails.chatId, offset)
            .then(messages => {
                if (messages.length < LIMIT) {
                    setLimitReached(true)
                }
                setOffset(offset + messages.length)
                messageUpdateFunc(messages)
            })
            .catch(err => console.error(err))
    }

    const onEndReached = () => {
        if (!limitReached) {
            updateMessages(nextPage => setMessages(messages.concat(nextPage)))
        }
    }

    useEffect(() => {
        updateMessages(messages => setMessages(messages))
    }, [])

    const sendMessage = (message) => {
        let newId = Math.floor(Math.random() * 100).toString()
        let tempMessage = {
            id: newId,
            chatId: chatDetails.chatId,
            message: message,
            sentBy: userId,
            sentAt: Date.now()
        }

        setMessages([tempMessage, ...messages])

        Network.POST('/messages/send', { chatId: chatDetails.chatId, message })
            .then(newMessage => {
                setMessages([newMessage, ...messages])
            })
            .catch(err => {
                console.error(err)
                setMessages(messages.filter(m => m.id != newId))
            })
    }

    return (
        <SafeAreaView style={styles.container}>
            <Header leftIcon={<BackIcon navigation={navigation} />} titleComponent={headerTitleComponent(chatDetails)} />

            <View style={styles.contentContainer}>
                <FlatList
                    inverted
                    onScrollBeginDrag={() => Keyboard.dismiss()}
                    data={messages}
                    contentContainerStyle={styles.messagesContainer}
                    onEndReached={onEndReached}
                    onEndReachedThreshold={0.1}
                    renderItem={({ item }) => <MessageContainer message={item} userId={userId} />}
                />
                <ChatInput onSend={sendMessage} />
            </View>

        </SafeAreaView>
    )
}

const loadMessages = (chatId, offset) => {
    let params = {
        chatId,
        limit: LIMIT,
        offset: offset
    }

    return Network.GET('/messages/by-chat-id', params)
}

const headerTitleComponent = chatDetails => {
    if (chatDetails.groupName) {
        return (
            <View style={styles.headerTitleContainer}>
                <Text style={styles.headerTitle}>{chatDetails.groupName}</Text>
            </View>
        )
    }
    return (
        <View style={styles.headerTitleContainer}>
            <Text style={styles.headerTitle}>{chatDetails.otherUserName}</Text>
            <Text style={styles.headerSubtitle}>@{chatDetails.otherUserUsername}</Text>
        </View>
    )
}


const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.primary,
    },
    headerTitleContainer: {
        alignItems: 'center'
    },
    headerTitle: {
        fontWeight: 'bold',
        color: colors.white,
        fontSize: 16,
    },
    headerSubtitle: {
        opacity: 0.75,
        color: colors.white,
        fontSize: 16,
    },
    contentContainer: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
    messagesContainer: {
        justifyContent: 'flex-end'
    },
})