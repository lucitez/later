import React, { useEffect, useState } from 'react'
import { StyleSheet, View, Text, SafeAreaView, FlatList, Image, Keyboard, Dimensions } from 'react-native'
import { colors } from '../assets/colors'
import { Header, BackIcon } from '../components/common'
import Network from '../util/Network'
import { useSelector } from 'react-redux'
import { Message, ContentMessage } from '../components/chat'
import { TextInput } from 'react-native-gesture-handler'

const LIMIT = 20

export default ChatDisplayScreen = ({ navigation, route }) => {
    let chatDetails = route.params.chatDetails

    const [userId, setUserId] = useState(useSelector(state => state.auth.userId))
    const [messages, setMessages] = useState([])
    const [limitReached, setLimitReached] = useState(false)
    const [chatInput, setChatInput] = useState('')
    const [isKeyboardShowing, setKeyboardShowing] = useState(false)
    const [keyboardYPos, setKeyboardYPos] = useState(0)
    const [inputBottomPos, setInputBottomPos] = useState(0)

    const _keyboardDidShow = (e) => {
        setKeyboardYPos(e.endCoordinates.screenY)
        setKeyboardShowing(true)
    }

    const _keyboardDidHide = (e) => {
        setKeyboardShowing(false)
    }

    const replaceMessages = messages => {
        if (messages.length < LIMIT) {
            setLimitReached(true)
        }
        setMessages(messages.reverse())
    }

    const prependMessages = nextPage => {
        if (nextPage.length < LIMIT) {
            setLimitReached(true)
        }
        setMessages(nextPage.reverse().concat(messages))
    }

    const updateMessages = (offset = 0, messageUpdateFunc) => {
        loadMessages(chatDetails.chatId, offset)
            .then(messages => messageUpdateFunc(messages))
            .catch(err => console.error(err))
    }

    useEffect(() => {
        updateMessages(0, replaceMessages)
        Keyboard.addListener("keyboardDidShow", _keyboardDidShow);
        Keyboard.addListener("keyboardDidHide", _keyboardDidHide);

        _input.measureInWindow((x, y, width, height) => {
            setInputBottomPos(y + height)
        })

        return () => {
            Keyboard.removeListener("keyboardDidShow", _keyboardDidShow);
            Keyboard.removeListener("keyboardDidHide", _keyboardDidHide);
        }
    }, [])

    const ChatInput = () => {
        return (
            <View
                style={[styles.inputBarContainer, isKeyboardShowing ? { marginBottom: inputBottomPos - keyboardYPos } : null]}
                ref={component => _input = component}
            >
                <View style={styles.inputContainer}>
                    <TextInput
                        value={chatInput}
                        multiline={true}
                        style={styles.input}
                        onChangeText={text => setChatInput(text)}
                    />
                </View>
            </View>
        )
    }

    return (
        <SafeAreaView style={styles.container}>
            <Header leftIcon={<BackIcon navigation={navigation} />} titleComponent={headerTitleComponent(chatDetails)} />

            <View style={styles.contentContainer}>
                <FlatList
                    data={messages}
                    contentContainerStyle={styles.messagesContainer}
                    renderItem={({ item }) => renderMessage(item, userId)}
                />
                {ChatInput()}
            </View>

        </SafeAreaView>
    )
}

const renderMessage = (message, userId) => {
    let self = userId == message.sentBy
    return (
        <View style={[
            styles.messageContainer,
            { justifyContent: self ? 'flex-end' : 'flex-start' }
        ]} >
            {!self &&
                <View style={styles.imageContainer} >
                    <Image style={styles.thumb} source={{ uri: 'https://www.washingtonpost.com/resizer/uwlkeOwC_3JqSUXeH8ZP81cHx3I=/arc-anglerfish-washpost-prod-washpost/public/HB4AT3D3IMI6TMPTWIZ74WAR54.jpg' }} />
                </View>
            }
            {message.message ?
                <Message message={message.message} />
                :
                <ContentMessage {...message.content} />
            }
            {self &&
                <View style={styles.imageContainer} >
                    <Image style={styles.thumb} source={{ uri: 'https://www.washingtonpost.com/resizer/uwlkeOwC_3JqSUXeH8ZP81cHx3I=/arc-anglerfish-washpost-prod-washpost/public/HB4AT3D3IMI6TMPTWIZ74WAR54.jpg' }} />
                </View>
            }
        </View>
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
    inputBarContainer: {
        backgroundColor: colors.white,
        minHeight: 50,
        marginTop: 10,
        paddingLeft: 10,
        paddingTop: 5,
        paddingBottom: 5,
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center'
    },
    inputContainer: {
        width: '80%',
        borderRadius: 20,
        padding: 10,
        paddingTop: 5,
        paddingBottom: 5,
        backgroundColor: colors.lightGray,
        justifyContent: 'center',
    },
    input: {
        marginBottom: 3,
        fontSize: 16,
    },
    contentContainer: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
    messagesContainer: {
        flexBasis: 0,
        flexGrow: 1,
        justifyContent: 'flex-end'
    },
    messageContainer: {
        width: '100%',
        flexDirection: 'row',
        alignItems: 'flex-end',
    },
    imageContainer: {
        height: 40,
        width: 40,
        margin: 5,
    },
    thumb: {
        height: '100%',
        width: '100%',
        borderRadius: 20,
    },
})