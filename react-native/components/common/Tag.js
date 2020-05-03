import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { colors } from '../../assets/colors';

function Tag({ name, size, theme }) {
    const backgroundColorFromTheme = theme == 'light' ? colors.white : colors.primary
    const textColorFromTheme = theme == 'light' ? colors.primary : colors.white

    return (
        <View style={styles.container}>
            <View style={[styles.tagContainer, { backgroundColor: backgroundColorFromTheme }, tagContainerFromSize(size)]}>
                <Text style={[tagFromSize(size), { color: textColorFromTheme }]}>{name}</Text>
            </View>
        </View>

    )
}

const tagContainerFromSize = size => {
    switch (size) {
        case 'large': return {
            borderRadius: 30,
            padding: 12,
            paddingTop: 8,
            paddingBottom: 8,
        }
        default: return {
            borderRadius: 15,
            padding: 7,
            paddingTop: 3,
            paddingBottom: 3,
        }
    }
}

const tagFromSize = size => {
    switch (size) {
        case 'large': return {
            fontSize: 16
        }
        default: return {
            fontSize: 14
        }
    }
}


const styles = StyleSheet.create({
    container: {
        flexDirection: 'column',
        alignItems: 'center',
    },
    tagContainer: {
        justifyContent: 'center',
        alignItems: 'center',

    },
})

export default Tag