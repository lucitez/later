import React from 'react'
import { View } from 'react-native'
import { Icon } from '../components/common'
import { colors } from '../assets/colors'

export default function SplashScreen() {
    return (
        <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
            <Icon type='share' size={40} color={colors.primary} />
        </View>
    )
}