import React from 'react'
import { View, Text, StyleSheet } from 'react-native'
import { colors } from '../../assets/colors'

export default function Message({ message }) {
    return (
        <View style={styles.container}>
            <Text style={styles.message}>{message}</Text>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        maxWidth: '60%',
        backgroundColor: colors.darkGray,
        margin: 5,
        marginTop: 10,
        padding: 10,
        borderRadius: 10,
    },
    message: {
        color: colors.white
    },
})