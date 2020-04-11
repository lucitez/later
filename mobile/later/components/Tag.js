import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import Colors from '../assets/colors';

function Tag(props) {
    return (
        <View style={styles.tagContainer}>
            <Text style={styles.tag}>{props.name}</Text>
        </View>
    )
}

const styles = StyleSheet.create({
    tagContainer: {
        backgroundColor: Colors.primary,
        justifyContent: 'center',
        alignItems: 'center',
        borderRadius: 5,
        padding: 5,
        paddingTop: 2,
        paddingBottom: 2,
    },
    tag: {
        color: Colors.white
    },
})

export default Tag