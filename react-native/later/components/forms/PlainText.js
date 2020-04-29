import React, { useEffect, useState } from 'react'
import { TextInput, StyleSheet } from 'react-native'
import { colors } from '../../assets/colors'
import FormInputWrapper from './FormInputWrapper'

export default function PlainText(props) {
    const color = props.theme == 'light' ? colors.white : colors.black
    const [value, setValue] = useState(props.value ? props.value : '')

    useEffect(() => {
        let error = hasError(value)
        props.onChange(props.name, value, error)
    }, [value])

    const hasError = () => {
        if (props.hasError) {
            return props.hasError(value)
        } else {
            if (props.required && value == '') {
                return `${props.title} is required`
            } else {
                return null
            }
        }
    }

    return (
        <FormInputWrapper {...props}>
            <TextInput
                {...props.inputProps}
                autoCorrect={false}
                autoCapitalize='none'
                style={[styles.input, { color: color }]}
                onChangeText={text => setValue(text)}
                value={value}
                selectionColor={color}
            />
        </FormInputWrapper>
    )
}

const styles = StyleSheet.create({
    input: {
        fontWeight: '400',
        fontSize: 18,
    },
})