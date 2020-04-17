import React, { useEffect, useState } from 'react'
import { View, Text, TextInput, StyleSheet } from 'react-native'
import { colors } from '../../assets/colors'

export default function Email(props) {

    const [value, setValue] = useState(props.value)

    useEffect(() => {
        let valid = isValid(value)
        props.onChange(props.name, value, valid)
    }, [value])

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
            <View style={[styles.underline]} />
        </View>
    )
}

const isValid = (value) => {
    const pattern = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/

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