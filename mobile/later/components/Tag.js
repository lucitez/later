import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { colors } from '../assets/colors';

function Tag(props) {
    return (
        <View style={styles.container}>

            <View style={styles.tagContainer}>
                <Text style={styles.tag}>{props.name}</Text>
            </View>
            {/* <View style={styles.triangle} /> */}
        </View>

    )
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'column',
        alignItems: 'center',
    },
    triangle: {
        borderTopWidth: 7,
        borderLeftWidth: 10,
        borderRightWidth: 10,
        borderLeftColor: 'transparent',
        borderRightColor: 'transparent',
        borderColor: colors.primary
    },
    tagContainer: {
        backgroundColor: colors.primary,
        justifyContent: 'center',
        alignItems: 'center',
        borderRadius: 3,
        padding: 5,
        paddingTop: 3,
        paddingBottom: 3,
    },
    tag: {
        color: colors.white
    },
})

export default Tag