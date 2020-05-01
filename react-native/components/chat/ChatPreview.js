import React from 'react';
import { StyleSheet, Text, View, Image } from 'react-native';
import { colors } from '../../assets/colors';
import { timeSince } from '../../util/time';

function ChatPreview({ chat }) {
    return (
        <View style={styles.container}>
            <View style={styles.imageContainer}>
                <Image style={styles.thumb} source={{ uri: 'https://www.washingtonpost.com/resizer/uwlkeOwC_3JqSUXeH8ZP81cHx3I=/arc-anglerfish-washpost-prod-washpost/public/HB4AT3D3IMI6TMPTWIZ74WAR54.jpg' }} />
            </View>
            <View style={{ flexGrow: 1 }}>
                <View style={styles.userInfoContainer}>
                    <View style={styles.nameContainer}>
                        <Text style={[styles.name, chat.hasUnread && { fontWeight: 'bold' }]}>{chat.otherUserName}</Text>
                    </View>
                    <View style={styles.usernameContainer}>
                        <Text style={[styles.username, chat.hasUnread && { fontWeight: '500' }]}>@{chat.otherUserUsername}</Text>
                    </View>
                </View>
                <View style={styles.messageContainer}>
                    <Text style={styles.message}>{chat.activity}{chat.lastMessageSentAt && ` - ${timeSince(Date.parse(chat.lastMessageSentAt))}`}</Text>
                </View>
            </View>
            {chat.hasUnread &&
                <View style={styles.hasUnreadContainer}>
                    <View style={styles.hasUnread} />
                </View>
            }
            <View>
                <Text>{chat.hasUnread}</Text>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        alignItems: 'center',
        height: 70,
        width: '100%',
        backgroundColor: colors.white
    },
    imageContainer: {
        aspectRatio: 1,
        height: '70%',
        marginLeft: 5,
        justifyContent: 'center',
        alignItems: 'center',
    },
    thumb: {
        height: '100%',
        width: '100%',
        borderRadius: 100,
        backgroundColor: 'coral',
    },
    userInfoContainer: {
        flexDirection: 'row',
        paddingTop: 5,
        marginBottom: 2,
    },
    nameContainer: {
        marginLeft: 5,
        marginRight: 5
    },
    usernameContainer: {
        marginTop: 1.5,
    },
    name: {
        fontSize: 16,
        fontWeight: '500'
    },
    username: {
        fontSize: 14,
        fontWeight: '300'
    },
    messageContainer: {
        marginLeft: 5,
    },
    message: {
        opacity: 0.75,
        fontStyle: 'italic'
    },
    hasUnreadContainer: {
        height: '100%',
        justifyContent: 'center',
        padding: 15,
    },
    hasUnread: {
        height: 7.5,
        width: 7.5,
        borderRadius: 5,
        backgroundColor: colors.primary
    }
});

export default ChatPreview