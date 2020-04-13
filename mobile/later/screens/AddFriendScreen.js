import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Text, Alert } from 'react-native';
import Header from '../components/Header';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import SearchBar from '../components/SearchBar';
import UserGroup from '../components/UserGroup';
import Icon from '../components/Icon';

const setRequestSent = (users, userId) => {
    return users.map(user =>
        user.id == userId ? { ...user, pendingRequest: true } : user
    )
}

const revertRequestSent = (users, userId) => {
    return users.map(user =>
        user.id == userId ? { ...user, pendingRequest: false } : user
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
                    <View style={styles.noUserContainer}>
                        <Text>Could not find any users that match your search</Text>
                    </View>
                    : <UserGroup type='add_friend' users={users} onRequestSent={onFriendRequestSent} />
            }
        </View>
    );
}

function BackIcon(navigation) {
    return <Icon type='back' size={25} color={colors.white} onPress={() => navigation.pop()} />
}

function sendFriendRequest(userId) {
    let queryString = `/friend-requests/send`
    let body = {
        senderUserId: 'b6e05c09-0f62-4757-95f5-ea855adc4915',
        recipientUserId: userId
    }

    return Network.POST(queryString, body)
}

const getUsers = (userId, search) => {
    params = {
        userId: userId,
        search: search
    }
    let queryString = `/users/add-friend-filter`
    return Network.GET(queryString, params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
    searchContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 10
    },
    noUserContainer: {
        marginTop: 10,
        width: '100%',
        alignItems: 'center',
    }
});


export default AddFriendScreen;