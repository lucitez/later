import React from 'react';
import { StyleSheet, Text, View, TouchableOpacity, ActivityIndicator } from 'react-native';
import { colors } from '../../assets/colors';

function Button({ theme, size, name, loading, onPress }) {
    return (
        <TouchableOpacity
            style={[
                styles.buttonContainer,
                buttonContainerStyleFromTheme(theme),
                buttonContainerStyleFromSize(size)
            ]}
            onPress={onPress}
            disabled={loading}
        >
            {loading ?
                <ActivityIndicator size='small' color={colors.white} />
                :
                <Text style={[
                    buttonStyleFromTheme(theme),
                    buttonStyleFromSize(size)
                ]}>
                    {name}
                </Text>
            }

        </TouchableOpacity>
    )

}

const buttonContainerStyleFromTheme = theme => {
    switch (theme) {
        case 'primary': return styles.primaryButtonContainer
        case 'light': return styles.lightButtonContainer
        case 'danger': return styles.dangerButtonContainer
    }
}

const buttonStyleFromTheme = theme => {
    switch (theme) {
        case 'primary': return styles.primaryButton
        case 'light': return styles.lightButton
        case 'danger': return styles.dangerButton
    }
}

const buttonContainerStyleFromSize = size => {
    switch (size) {
        case 'small': return styles.smallButtonContainer
        case 'medium': return styles.mediumButtonContainer
        case 'large': return styles.largeButtonContainer
    }
}

const buttonStyleFromSize = size => {
    switch (size) {
        case 'small': return styles.smallButton
        case 'medium': return styles.mediumButton
        case 'large': return styles.largeButton
    }
}

const styles = StyleSheet.create({
    buttonContainer: {
        borderRadius: 5,
        alignItems: 'center',
        marginTop: 5,
        marginBottom: 5,
        borderWidth: 1.5,
    },
    primaryButtonContainer: {
        borderColor: colors.white,
        backgroundColor: colors.primary,
    },
    lightButtonContainer: {
        backgroundColor: colors.white,
        borderColor: colors.primary
    },
    dangerButtonContainer: {
        backgroundColor: colors.red,
        borderColor: colors.white
    },
    primaryButton: {
        color: colors.white
    },
    lightButton: {
        color: colors.primary
    },
    dangerButton: {
        color: colors.white
    },
    largeButtonContainer: {
        padding: 15,
    },
    mediumButtonContainer: {
        padding: 10,
    },
    smallButtonContainer: {
        padding: 5
    },
    smallButton: {
        fontSize: 12
    },
    mediumButton: {
        fontSize: 16
    },
    largeButton: {
        fontSize: 20
    }
})

export default Button