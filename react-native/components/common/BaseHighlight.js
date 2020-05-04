import React from 'react'
import { TouchableHighlight } from 'react-native'

export default function BaseHighlight({ onPress, style, children }) {
    return (
        <TouchableHighlight activeOpacity={0.50} underlayColor='#FFFFFF' onPress={onPress} style={style}>
            {children}
        </TouchableHighlight>
    )
}