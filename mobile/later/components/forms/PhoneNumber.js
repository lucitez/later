import React, { useEffect, useState } from 'react'
import { View, Text, TextInput, StyleSheet } from 'react-native'
import { colors } from '../../assets/colors'

export default function PhoneNumber(props) {

    const [rawValue, setRawValue] = useState(props.value)
    const [displayValue, setDisplayValue] = useState(toDisplayValue(props.value))

    const onChangeText = (text) => {
        let rawValue = toRawValue(text)
        setRawValue(rawValue)
        setDisplayValue(toDisplayValue(rawValue))
    }

    useEffect(() => {
        let valid = isValid(rawValue)
        props.onChange(props.name, rawValue, valid)
    }, [rawValue])

    return (
        <View style={styles.container}>
            <View style={styles.nameContainer}>
                <Text style={styles.name}>{props.title}:</Text>
            </View>
            <View style={styles.inputContainer}>
                <TextInput
                    keyboardType='phone-pad'
                    style={styles.input}
                    onChangeText={onChangeText}
                    value={displayValue}
                    selectionColor={colors.white}
                />
            </View>
            <View style={[styles.underline]} />
        </View>
    )
}

const toRawValue = (text) => {
    var raw = text.replace(/[^0-9]+/g, '')
    return raw
}

const toDisplayValue = (rawText) => {
    switch (true) {
        case rawText.length == 0:
            return ``
        case rawText.length <= 3:
            return `(${rawText}`
        case rawText.length <= 6:
            return `(${rawText.slice(0, 3)}) ${rawText.slice(3)}`
        default:
            return `(${rawText.slice(0, 3)}) ${rawText.slice(3, 6)} - ${rawText.slice(6)}`
    }
}

const isValid = (value) => {
    const pattern = /^[\d]{10}$/

    return pattern.test(value)
}

const styles = StyleSheet.create({
    container: {
        padding: 5,
    },
    nameContainer: {
        marginBottom: 5,
    },
    underline: {
        height: 1.5,
        marginTop: 5,
        backgroundColor: colors.white
    },
    name: {
        color: colors.white,
        fontWeight: '300',
        fontSize: 14,
    },
    input: {
        color: colors.white,
        fontWeight: '500',
        fontSize: 18,
    },
})