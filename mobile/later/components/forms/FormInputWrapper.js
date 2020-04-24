import React from 'react'
import { View, Text, StyleSheet } from 'react-native'
import { colors } from '../../assets/colors'

export default function FormInputWrapper({ theme, title, children }) {
    const color = theme == 'light' ? colors.white : colors.black

    return (
        <View style={styles.container}>
            <View style={styles.nameContainer}>
                <Text style={[styles.name, { color: color }]}>{title}</Text>
            </View>
            <View style={styles.inputContainer}>
                {children}
            </View>
            <View style={[styles.underline, { backgroundColor: color }]} />
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        marginBottom: 15,
    },
    nameContainer: {
        marginBottom: 15,
    },
    underline: {
        height: 1,
        marginTop: 5,
        opacity: 0.25,
    },
    name: {
        fontWeight: '300',
        fontSize: 14,
    }
})