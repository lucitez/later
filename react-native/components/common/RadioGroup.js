import React, { useState, useEffect } from 'react'
import { StyleSheet, View } from 'react-native'
import Radio from './Radio'
import { colors } from '../../assets/colors'

export default function RadioGroup({ options, onChange }) {
    const [selectedOption, setSelectedOption] = useState(null)

    useEffect(() => onChange(selectedOption), [selectedOption])

    return (
        <View style={styles.container}>
            {options.map((option, index) => (
                <Radio
                    selected={option.value == selectedOption}
                    display={option.display}
                    icon={option.icon}
                    key={option.value}
                    first={index == 0}
                    last={index == options.length - 1}
                    onPress={() => selectedOption != option.value ? setSelectedOption(option.value) : setSelectedOption(null)}
                />
            ))}
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        flexGrow: 1,
        borderRadius: 10,
        borderWidth: 1,
        borderColor: colors.darkGray,
    },
})
