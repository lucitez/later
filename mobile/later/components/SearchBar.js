import React, { useState, useEffect } from 'react';
import { StyleSheet, Text, TextInput, View, Keyboard, TouchableOpacity } from 'react-native';
import Colors from '../assets/colors';
import Icon from '../components/Icon';

function SearchBar(props) {
    const [search, setSearch] = useState('')

    useEffect(() => {
        props.onChange(search)
    }, [search])

    return (
        <View style={styles.container}>
            <View style={styles.searchBarContainer} >
                <Icon type='search' size={20} color={Colors.black} />
                <View style={styles.inputContainer} >
                    <TextInput autoCorrect={false} placeholder='Search...' onChangeText={text => setSearch(text)} value={search} />
                </View>
            </View>
            <TouchableOpacity onPress={() => {
                setSearch('')
                Keyboard.dismiss()
            }}>
                <View style={styles.cancelContainer} >

                    <Text style={{ color: Colors.white }}>Cancel</Text>
                </View>
            </TouchableOpacity>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        backgroundColor: Colors.gray,
        height: 40,
        width: '100%',
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: "center",
        backgroundColor: '#D1D1D1',
        padding: 5,
        paddingLeft: 7,
    },
    searchBarContainer: {
        flexGrow: 1,
        flexDirection: 'row',
        height: '100%',
        borderRadius: 10,
        backgroundColor: Colors.white,
        alignItems: 'center',
        paddingLeft: 10,
    },
    inputContainer: {
        height: '100%',
        justifyContent: 'center',
        marginLeft: 1,
        flexGrow: 1,
    },
    cancelContainer: {
        marginLeft: 10,
        marginRight: 5,
        justifyContent: 'center',
        height: '100%',
    }
});

export default SearchBar