import React, { useEffect, useState } from 'react'
import { TextInput, StyleSheet } from 'react-native'
import { colors } from '../../assets/colors'
import FormInputWrapper from './FormInputWrapper'

export default function PhoneNumber(props) {
    const color = props.theme == 'light' ? colors.white : colors.black

    const [rawValue, setRawValue] = useState(props.value ? props.value : '')
    const [displayValue, setDisplayValue] = useState(toDisplayValue(props.value ? props.value : ''))

    const onChangeText = (text) => {
        let rawValue = toRawValue(text)
        setRawValue(rawValue)
        setDisplayValue(toDisplayValue(rawValue))
    }

    useEffect(() => {
        let error = hasError(props.required, rawValue)
        props.onChange(props.name, rawValue, error)
    }, [rawValue])

    return (
        <FormInputWrapper {...props}>
            <TextInput
                keyboardType='phone-pad'
                style={[styles.input, { color: color }]}
                onChangeText={onChangeText}
                value={displayValue}
                selectionColor={color}
            />
        </FormInputWrapper>
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

const hasError = (required, value) => {
    if (required && value == '') {
        return 'phone number is required'
    }

    const pattern = /^[\d]{10}$/

    let valid = pattern.test(value)

    if (valid) {
        return null
    }

    return 'Invalid phone number format'
}

const styles = StyleSheet.create({
    input: {
        fontWeight: '400',
        fontSize: 18,
    },
})