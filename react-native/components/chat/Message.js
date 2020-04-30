import React from 'react'
import { View, Text, StyleSheet } from 'react-native'
import { colors } from '../../assets/colors'

export default function Message({ message, fromMe }) {
    return (
        <View style={[styles.container, { backgroundColor: fromMe ? colors.blue : colors.darkGray }]}>
            <Text style={styles.message}>{message}</Text>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        maxWidth: '60%',
        margin: 5,
        marginTop: 10,
        padding: 10,
        borderRadius: 10,
    },
    message: {
        color: colors.white
    },
})