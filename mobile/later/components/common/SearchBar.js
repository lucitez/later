import React, { useState, useEffect } from 'react';
import { StyleSheet, Text, TextInput, View, Keyboard, TouchableOpacity } from 'react-native';
import { colors } from '../../assets/colors';
import Icon from './Icon';

function SearchBar(props) {
    const [search, setSearch] = useState(props.startingValue ? props.startingValue : '')
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
    }, [])

    useEffect(() => {
        props.onChange(search)
    }, [search])

    useEffect(() => {
        if (props.clear) {
            setSearch('')
        }
    }, [props.clear])

    const _cancelButton = () => {
        if (props.showCancelOnKeyboardActive && !isKeyboardShowing) {
            return null
        } else {
            return (
                <TouchableOpacity onPress={() => {
                    setSearch('')
                    Keyboard.dismiss()
                    if (props.onCancel) props.onCancel()
                }}>
                    <View style={styles.cancelContainer} >
                        <Text style={styles.cancel}>Cancel</Text>
                    </View>
                </TouchableOpacity>
            )
        }
    }

    const _clearButton = () => {
        if (search.length > 0) {
            return (
                <View style={styles.clearIconContainer}>
                    <Icon type='close' size={20} color={colors.darkGray} onPress={() => setSearch('')} />
                </View>
            )
        }
    }

    return (
        <View style={styles.container}>
            <View style={styles.searchBarContainer} >
                <View style={styles.searchIconContainer} >
                    <Icon type={props.iconName ? props.iconName : 'search'} size={20} color={colors.darkGray} />
                </View>
                <View style={styles.inputContainer} >
                    <TextInput
                        autoFocus={props.autoFocus}
                        autoCorrect={false}
                        returnKeyType={props.returnKeyType}
                        placeholder={props.placeholder ? props.placeholder : 'Search...'}
                        onChangeText={text => setSearch(text)}
                        value={search}
                    />
                </View>
                {_clearButton()}
            </View>
            {_cancelButton()}
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        backgroundColor: colors.primary,
        height: 50,
        width: '100%',
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: "center",
        padding: 5,
        paddingLeft: 10,
        paddingRight: 10
    },
    searchBarContainer: {
        flexGrow: 1,
        flexDirection: 'row',
        height: '100%',
        borderRadius: 10,
        backgroundColor: colors.white,
        alignItems: 'center',
        paddingLeft: 10,
    },
    searchIconContainer: {
        paddingTop: 3,
    },
    clearIconContainer: {
        paddingRight: 5,
        paddingTop: 2,
        marginLeft: 5,
    },
    inputContainer: {
        height: '100%',
        justifyContent: 'center',
        marginLeft: 5,
        flexGrow: 1,
        flexBasis: 0,
    },
    cancelContainer: {
        marginLeft: 10,
        marginRight: 5,
        justifyContent: 'center',
        height: '100%',
    },
    cancel: {
        color: colors.white,
        fontSize: 16,
    }
});

export default SearchBar