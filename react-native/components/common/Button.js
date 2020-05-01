import React from 'react';
import { StyleSheet, Text, TouchableOpacity, ActivityIndicator } from 'react-native';
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
    }
}

const buttonStyleFromTheme = theme => {
    switch (theme) {
        case 'primary': return styles.primaryButton
        case 'light': return styles.lightButton
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
    },
    primaryButtonContainer: {
        backgroundColor: colors.primary,
        borderWidth: 1.5,
        borderColor: colors.white
    },
    lightButtonContainer: {
        backgroundColor: colors.white,
        borderWidth: 1.5,
        borderColor: colors.primary
    },
    primaryButton: {
        color: colors.white
    },
    lightButton: {
        color: colors.primary
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