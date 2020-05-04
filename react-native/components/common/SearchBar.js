import React, { useState, useEffect } from 'react';
import { StyleSheet, Text, TextInput, View, Keyboard, TouchableOpacity, FlatList, ActivityIndicator } from 'react-native';
import { colors } from '../../assets/colors';
import Icon from './Icon';
import Tag from './Tag';

function SearchBar(props) {
    const [search, setSearch] = useState(props.value ? props.value : '')
    const [isKeyboardShowing, setKeyboardShowing] = useState(false)
    const [autocompleteOptions, setAutocompleteOptions] = useState([])
    const [loading, setLoading] = useState(false)

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
        console.log('HAPPENING IN SEARCH')
        props.onChange(search)
        if (props.autocompleteFunc) {
            setLoading(true)
            props.autocompleteFunc(search)
                .then(options => {
                    setAutocompleteOptions(options)
                    setLoading(false)
                })
                .catch(err => {
                    console.error(err)
                    setLoading(false)
                })
        }
    }, [search])

    useEffect(() => {
        console.log('HAPPENING IN CLEAR')
        if (props.clear) {
            setSearch('')
        }
    }, [props.clear])

    const _rightIcon = () => {
        if (props.rightIcon) {
            return (
                <View style={styles.rightIconContainer} >
                    {props.rightIcon}
                </View>
            )
        }

        if (props.showCancelOnKeyboardActive && !isKeyboardShowing) {
            return null
        } else {
            return (
                <TouchableOpacity onPress={() => {
                    setSearch('')
                    Keyboard.dismiss()
                    if (props.onCancel) props.onCancel()
                }}>
                    <View style={styles.rightIconContainer} >
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

    const _autocompleteContent = () => {
        const renderAutocompleteData = ({ item }) => (
            <TouchableOpacity style={styles.optionContainer} onPress={() => setSearch(item)}>
                <Tag name={item} theme='light' size='large' />
            </TouchableOpacity>
        )

        if (props.autocompleteFunc) {
            return (
                <View style={styles.autoCompleteContainer}>
                    {loading ? (
                        <View style={styles.autocompleteContainer}>
                            <View style={{ alignItems: 'center', marginBottom: 10 }}>
                                <Text style={{ color: colors.white }}>Loading your similar tags</Text>
                            </View>
                            <ActivityIndicator color={colors.white} />
                        </View>
                    ) : (
                            <FlatList
                                keyboardShouldPersistTaps='handled'
                                data={autocompleteOptions}
                                keyExtractor={(_, index) => index.toString()}
                                renderItem={renderAutocompleteData}
                                persistentScrollbar={true}
                            />
                        )

                    }

                </View>
            )
        } else return null
    }

    return (
        <View style={props.autocompleteFunc && { flexBasis: 0, flexGrow: 1 }}>
            <View style={styles.topContainer}>
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
                {_rightIcon()}
            </View>
            {_autocompleteContent()}
        </ View>
    );
}

const styles = StyleSheet.create({
    topContainer: {
        height: 55,
        width: '100%',
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: "center",
        backgroundColor: colors.primary,
        padding: 10,
    },
    autoCompleteContainer: {
        flexBasis: 0,
        flexGrow: 1,
        padding: 5,
        paddingLeft: 10,
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
    rightIconContainer: {
        marginLeft: 10,
        marginRight: 5,
        justifyContent: 'center',
        height: '100%',
    },
    cancel: {
        color: colors.white,
        fontSize: 16,
    },
    autocompleteContainer: {
        flexBasis: 0,
        flexGrow: 1,
        padding: 15
    },
    optionContainer: {
        padding: 5,
        alignItems: 'flex-start'
    },
});

export default SearchBar