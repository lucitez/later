import React from 'react'
import { View, Text, StyleSheet } from 'react-native'
import { colors } from '../../assets/colors'

export default function FormInputWrapper({ theme, title, subtitle, children }) {
    const color = theme == 'light' ? colors.white : colors.black

    return (
        <View style={styles.container}>
            <View style={styles.titleContainer}>
                <Text style={[styles.title, { color: color }]}>{title}</Text>
                {subtitle && <Text style={[styles.subtitle, { color }]}> - {subtitle}</Text>}
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
    titleContainer: {
        marginBottom: 15,
        flexDirection: 'row',
        alignItems: 'flex-end',
    },
    underline: {
        height: 1,
        marginTop: 5,
        opacity: 0.25,
    },
    title: {
        fontWeight: '300',
        fontSize: 14,
    },
    subtitle: {
        fontWeight: '300',
        fontSize: 13,
    }
})