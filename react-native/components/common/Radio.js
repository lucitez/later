import React from 'react'
import { StyleSheet, TouchableOpacity } from 'react-native'
import { colors, contentTypes } from '../../assets/colors'
import Icon from './Icon'

export default function Radio({ selected, display, icon, onPress, first, last }) {
    return (
        <TouchableOpacity
            style={[styles.container, selected && selectedContainer(icon), first && styles.firstOptionContainer, last && styles.lastOptionContainer]}
            activeOpacity={0.7}
            underlayColor={colors.white}
            onPress={() => onPress()}
        >
            <>
                {display && <Text style={{ color: colorFromSelected(null, selected) }}>{display}</Text>}
                {icon && <Icon type={icon} size={25} color={colorFromSelected(icon, selected)} />}
            </>
        </TouchableOpacity>
    )
}

const selectedContainer = icon => {
    switch (icon) {
        case undefined:
            return { backgroundColor: colors.darkGray }
        default:
            return { backgroundColor: contentTypes[icon].color }
    }
}

const colorFromSelected = (icon, selected) => {
    if (selected) return colors.white
    if (!selected && !icon) return colors.darkGray
    if (!selected && icon) return contentTypes[icon].color
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center'
    },
    firstOptionContainer: {
        borderTopLeftRadius: 9,
        borderBottomLeftRadius: 9,
    },
    lastOptionContainer: {
        borderTopRightRadius: 9,
        borderBottomRightRadius: 9,
    }
})