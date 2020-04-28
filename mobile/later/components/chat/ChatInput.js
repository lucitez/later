import React, { useState, useEffect } from 'react'
import { StyleSheet, View, Text, TextInput, Keyboard } from "react-native"
import { colors } from '../../assets/colors'
import { TouchableOpacity } from 'react-native-gesture-handler'

export default function ChatInput({ onSend }) {
    const [chatInput, setChatInput] = useState('')
    const [isKeyboardShowing, setKeyboardShowing] = useState(false)
    const [keyboardYPos, setKeyboardYPos] = useState(0)
    const [inputBottomPos, setInputBottomPos] = useState(0)

    const _keyboardWillShow = (e) => {
        setKeyboardYPos(e.endCoordinates.screenY)
        setKeyboardShowing(true)
    }

    const _keyboardWillHide = (e) => setKeyboardShowing(false)

    useEffect(() => {
        Keyboard.addListener("keyboardWillShow", _keyboardWillShow);
        Keyboard.addListener("keyboardWillHide", _keyboardWillHide);

        _input.measureInWindow((x, y, width, height) => {
            setInputBottomPos(y + height)
        })

        return () => {
            Keyboard.removeListener("keyboardWillShow", _keyboardWillShow);
            Keyboard.removeListener("keyboardWillHide", _keyboardWillHide);
        }
    }, [])

    return (
        <View
            style={[styles.inputBarContainer, isKeyboardShowing ? { marginBottom: inputBottomPos - keyboardYPos } : null]}
            ref={component => _input = component}
        >
            <View style={styles.inputContainer}>
                <TextInput
                    value={chatInput}
                    multiline={true}
                    style={styles.input}
                    onChangeText={text => setChatInput(text)}
                />
            </View>
            <TouchableOpacity style={styles.sendContainer} onPress={() => {
                if (chatInput != '') {
                    onSend(chatInput)
                    setChatInput('')
                }
            }}>
                <Text style={styles.send}>Send</Text>
            </TouchableOpacity>
        </View>
    )
}

const styles = StyleSheet.create({
    inputBarContainer: {
        backgroundColor: colors.white,
        minHeight: 50,
        padding: 10,
        paddingTop: 5,
        paddingBottom: 5,
        flexDirection: 'row',
        justifyContent: 'flex-start',
        alignItems: 'center',
    },
    inputContainer: {
        flexGrow: 1,
        borderRadius: 20,
        padding: 10,
        paddingTop: 5,
        paddingBottom: 5,
        backgroundColor: colors.lightGray,
        justifyContent: 'center',
    },
    input: {
        marginBottom: 3,
        fontSize: 16,
    },
    sendContainer: {
        flexGrow: 1,
        padding: 5,
        paddingLeft: 10,
        paddingRight: 0,
        marginBottom: 3,
        justifyContent: 'flex-end',
    },
    send: {
        color: colors.primary,
        fontSize: 16,
        fontWeight: '600'
    },
})