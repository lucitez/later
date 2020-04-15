import React, { useState, useEffect } from 'react';
import { StyleSheet, Text, TextInput, View, Keyboard, TouchableOpacity } from 'react-native';
import { colors } from '../assets/colors';
import Icon from '../components/Icon';

function SearchBar(props) {
    const [search, setSearch] = useState(props.value ? props.value : '')

    useEffect(() => {
        props.onChange(search)
    }, [search])

    useEffect(() => {
        if (props.clear) {
            setSearch('')
        }
    }, [props.clear])

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
                {
                    search.length > 0 ?
                        <View style={styles.clearIconContainer}>
                            <Icon type='close' size={20} color={colors.darkGray} onPress={() => setSearch('')} />
                        </View>
                        : null
                }
            </View>
            <TouchableOpacity onPress={() => {
                setSearch('')
                Keyboard.dismiss()
                if (props.onCancel) props.onCancel()
            }}>
                <View style={styles.cancelContainer} >
                    <Text style={{ color: colors.white }}>Cancel</Text>
                </View>
            </TouchableOpacity>
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
        paddingLeft: 7,
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
    }
});

export default SearchBar