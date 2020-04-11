import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import Header from '../components/Header';
import Colors from '../assets/colors';
import Network from '../util';
import SearchBar from '../components/SearchBar';
import UserGroup from '../components/UserGroup';
import Icon from '../components/Icon';

function AddFriendScreen({ navigation }) {
    const [users, setUsers] = useState([])
    const [search, setSearch] = useState('')

    useEffect(() => {
        if (search.length == 0 || search.length > 2) {
            getUsers(search)
                .then(response => response.json())
                .then(users => setUsers(users))
                .catch(error => console.error(error))
        }
    }, [search])

    return (
        <View style={styles.container}>
            <Header name="Add Friends" leftIcon={BackIcon(navigation)} />
            <SearchBar onChange={search => setSearch(search)} />
            {
                users.length == 0 ?
                    <View style={styles.noFriendsContainer}><Text>Could not find any users that match your search</Text></View>
                    : <UserGroup users={users} />
            }
        </View>
    );
}

function BackIcon(navigation) {
    return (
        <Icon type='back' size={25} color={Colors.white} onPress={() => navigation.pop()} />
    )
}

const getUsers = (search) => {
    let queryString = `${Network.local}/users/filter?search=${search}`
    return fetch(queryString)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: Colors.lightGray,
    },
    searchContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 10
    },
    noFriendsContainer: {
        marginTop: 10,
        width: '100%',
        alignItems: 'center',
    }
});


export default AddFriendScreen;