import React, { useEffect, useState } from 'react'
import { View, Text, TextInput, StyleSheet } from 'react-native'
import { colors } from '../../assets/colors'

export default function PlainText(props) {

    const [value, setValue] = useState(props.value)

    useEffect(() => props.onChange(props.name, value), [value])

    return (
        <View style={styles.container}>
            <View style={styles.nameContainer}>
                <Text style={styles.name}>{props.title}:</Text>
            </View>
            <View style={styles.inputContainer}>
                <TextInput
                    style={styles.input}
                    onChangeText={text => setValue(text)}
                    value={value}
                    selectionColor={colors.white}
                />
            </View>
            <View style={styles.underline} />
        </View>
    )
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
        marginTop: 1,
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