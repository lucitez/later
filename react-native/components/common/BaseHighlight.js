import React from 'react'
import { TouchableHighlight } from 'react-native'

export default function BaseHighlight({ onPress, style, children }) {
    return (
        <TouchableHighlight activeOpacity={0.75} underlayColor='#FFFFFF' onPress={onPress} style={style}>
            {children}
        </TouchableHighlight>
    )
}