import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Text, Alert } from 'react-native';
import Header from '../components/Header';
import Colors from '../assets/colors';
import Network from '../util/Network';
import SearchBar from '../components/SearchBar';
import UserGroup from '../components/UserGroup';
import Icon from '../components/Icon';

const setRequestSent = (users, userId) => {
    return users.map(user =>
        user.id == userId ? { ...user, pending_request: true } : user
    )
}

const revertRequestSent = (users, userId) => {
    return users.map(user =>
        user.id == userId ? { ...user, pending_request: false } : user
    )
}

function AddFriendScreen({ navigation }) {
    const [users, setUsers] = useState([])
    const [search, setSearch] = useState('')

    useEffect(() => {
        if (search.length == 0 || search.length > 2) {
            getUsers('b6e05c09-0f62-4757-95f5-ea855adc4915', search)
                .then(users => setUsers(users))
                .catch(error => console.error(error))
        }
    }, [search])

    const onFriendRequestSent = userId => {
        setUsers(setRequestSent(users, userId))

        sendFriendRequest(userId)
            .then()
            .catch((err) => {
                setUsers(revertRequestSent(users, userId))
                Alert.alert(err)
            })
    }

    return (
        <View style={styles.container}>
            <Header name="Add Friends" leftIcon={BackIcon(navigation)} />
            <SearchBar onChange={search => setSearch(search)} />
            {
                users.length == 0 ?
                    <View style={styles.noFriendsContainer}><Text>Could not find any users that match your search</Text></View>
                    : <UserGroup users={users} onRequestSent={onFriendRequestSent} />
            }
        </View>
    );
}

function BackIcon(navigation) {
    return <Icon type='back' size={25} color={Colors.white} onPress={() => navigation.pop()} />
}

function sendFriendRequest(userId) {
    let queryString = `/friend-requests/send`
    let body = {
        sender_user_id: 'b6e05c09-0f62-4757-95f5-ea855adc491',
        recipient_user_id: userId
    }

    return Network.POST(queryString, body)
}

const getUsers = (userId, search) => {
    let queryString = `/users/add-friend-filter?user_id=${userId}&search=${search}`
    return Network.GET(queryString)
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