
import React, { useState, useEffect } from 'react'
import { StyleSheet, View, Keyboard } from 'react-native'
import { colors } from '../assets/colors';

function BottomSheetContainer(props) {
    const [isKeyboardShowing, setKeyboardShowing] = useState(false)

    Keyboard.addListener('keyboardWillShow', function () {
        setKeyboardShowing(true)
    })
    Keyboard.addListener('keyboardWillHide', function () {
        setKeyboardShowing(false)
    })

    useEffect(() => {
        return function cleanup() {
            Keyboard.removeAllListeners('keyboardWillShow')
            Keyboard.removeAllListeners('keyboardWillHide')
        };
    })

    return (
        <View style={[
            styles.archiveBottomSheet,
            isKeyboardShowing ? { paddingBottom: 5 } : { paddingBottom: 30 },
            props.height ? { height: props.height } : null
        ]}>
            {props.children}
        </View>
    )
}

const styles = StyleSheet.create({
    archiveBottomSheet: {
        backgroundColor: colors.primary,
        justifyContent: 'flex-start'
    },
});

export default BottomSheetContainer