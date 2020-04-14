import React from 'react';
import { StyleSheet, View } from 'react-native';
import Button from './Button';

function ButtonGroup(props) {
    return (
        <View style={styles.buttonGroupContainer}>
            {props.buttonProps.map(buttonProp => (
                <Button {...buttonProp} key={buttonProp.name} />
            ))}
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