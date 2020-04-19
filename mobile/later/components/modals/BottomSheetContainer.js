
import React, { useState, useEffect } from 'react'
import { StyleSheet, View, Keyboard } from 'react-native'
import { colors } from '../../assets/colors';

function BottomSheetContainer(props) {
    const [isKeyboardShowing, setKeyboardShowing] = useState(false)

    const _keyboardWillShow = () => setKeyboardShowing(true)
    const _keyboardWillHide = () => setKeyboardShowing(false)

    useEffect(() => {
        Keyboard.addListener("keyboardWillShow", _keyboardWillShow);
        Keyboard.addListener("keyboardWillHide", _keyboardWillHide);
        return () => {
            Keyboard.removeListener("keyboardWillShow", _keyboardWillShow);
            Keyboard.removeListener("keyboardWillHide", _keyboardWillHide);
        }
    })

    return (
        <View style={[
            styles.saveBottomSheet,
            isKeyboardShowing ? { paddingBottom: 5 } : { paddingBottom: 30 },
            props.height ? { height: props.height } : null
        ]}>
            {props.children}
        </View>
    )
}

const styles = StyleSheet.create({
    saveBottomSheet: {
        backgroundColor: colors.primary,
        justifyContent: 'flex-start'
    },
});

export default BottomSheetContainer