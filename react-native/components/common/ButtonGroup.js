import React from 'react';
import { StyleSheet, View } from 'react-native';

function ButtonGroup(props) {
    return (
        <View style={styles.buttonGroupContainer}>
            {props.children}
        </View>
    )
}

const styles = StyleSheet.create({
    buttonGroupContainer: {
        flexDirection: 'column',
        flexGrow: 1,
        padding: 10,
    }
})

export default ButtonGroup