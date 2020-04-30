import React, { useRef } from 'react'
import { Animated, StyleSheet, View, Image } from 'react-native'
import Message from './Message'
import ContentMessage from './ContentMessage'

export default function MessageContainer({ message, userId }) {
    let fromMe = userId == message.sentBy

    const slideUpAnim = useRef(new Animated.Value(-10)).current

    React.useEffect(() => {
        Animated.timing(
            slideUpAnim,
            {
                toValue: 0,
                duration: 200,
            }
        ).start();
    }, [])

    return (
        <Animated.View style={[
            styles.messageContainer,
            { justifyContent: fromMe ? 'flex-end' : 'flex-start' },
            { marginBottom: slideUpAnim }
        ]} >
            {!fromMe && <UserProfile />}
            {message.message ? <Message message={message.message} fromMe={fromMe} /> : <ContentMessage {...message.content} fromMe={fromMe} />}
            {fromMe && <UserProfile />}
        </Animated.View>
    )
}

const UserProfile = () => {
    return (
        <View style={styles.imageContainer} >
            <Image style={styles.thumb} source={{ uri: 'https://www.washingtonpost.com/resizer/uwlkeOwC_3JqSUXeH8ZP81cHx3I=/arc-anglerfish-washpost-prod-washpost/public/HB4AT3D3IMI6TMPTWIZ74WAR54.jpg' }} />
        </View>
    )
}

const styles = StyleSheet.create({
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