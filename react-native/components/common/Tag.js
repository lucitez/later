import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { colors } from '../../assets/colors';

function Tag(props) {
    return (
        <View style={styles.container}>
            <View style={styles.tagContainer}>
                <Text style={styles.tag}>{props.name}</Text>
            </View>
        </View>

    )
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'column',
        alignItems: 'center',
    },
    tagContainer: {
        backgroundColor: colors.primary,
        justifyContent: 'center',
        alignItems: 'center',
        borderRadius: 15,
        paddingLeft: 7,
        paddingRight: 7,
        paddingTop: 3,
        paddingBottom: 3,
    },
    tag: {
        color: colors.white
    },
})

export default Tag