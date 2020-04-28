import React from 'react'
import { View } from 'react-native'

export default Message = ({ message }) => {
    return (
        <View>
            <Text>{message}</Text>
        </View>
    )
}