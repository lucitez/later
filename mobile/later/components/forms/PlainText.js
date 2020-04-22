import React, { useEffect, useState } from 'react'
import { View, Text, TextInput, StyleSheet } from 'react-native'
import { colors } from '../../assets/colors'

export default function PlainText(props) {



    const color = props.theme == 'light' ? colors.white : colors.black

    const [value, setValue] = useState(props.value)


    useEffect(() => {
        let valid = isValid(value)
        props.onChange(props.name, value, valid)
    }, [value])

    const isValid = () => {
        if (props.isValid) {
            return props.isValid(value)
        } else {
            return true
        }
    }

    return (
        <View style={styles.container}>
            <View style={styles.nameContainer}>
                <Text style={[styles.name, { color: color }]}>{props.title}</Text>
            </View>
            <View style={styles.inputContainer}>
                <TextInput
                    {...props}
                    autoCapitalize='none'
                    style={[styles.input, { color: color }]}
                    onChangeText={text => setValue(text)}
                    value={value}
                    selectionColor={color}
                />
            </View>
            <View style={[styles.underline, { backgroundColor: color }]} />
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        padding: 5,
    },
    nameContainer: {
        marginBottom: 10,
    },
    underline: {
        height: 1,
        marginTop: 5,
        opacity: 0.5,
        backgroundColor: colors.white
    },
    name: {
        fontWeight: '300',
        fontSize: 14,
    },
    input: {
        fontWeight: '400',
        fontSize: 18,
    },
})