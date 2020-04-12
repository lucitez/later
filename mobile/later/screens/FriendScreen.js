import React, { useState, useEffect } from 'react';
import { StyleSheet, View } from 'react-native';
import Header from '../components/Header';
import Icon from '../components/Icon';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import UserGroup from '../components/UserGroup';
import SearchBar from '../components/SearchBar';

function FriendScreen({ navigation }) {
    const [friends, setFriends] = useState([])
    const [search, setSearch] = useState('')
    const [offset, setOffset] = useState(0)

    useEffect(() => {
        getFriends('b6e05c09-0f62-4757-95f5-ea855adc4915', search, offset)
            .then(nextPage => setFriends(friends.concat(nextPage)))
            .catch(error => console.error(error))
    }, [offset])

    useEffect(() => {
        getFriends('b6e05c09-0f62-4757-95f5-ea855adc4915', search, offset)
            .then(friends => setFriends(friends))
            .catch(error => console.error(error))
    }, [search])

    return (
        <View style={styles.container}>
            <Header name="Friends" rightIcon={AddFriendsIcon(navigation)} />
            <SearchBar onChange={search => setSearch(search)} />
            <UserGroup users={friends} type='friend' />
        </View>
    );
}

function AddFriendsIcon(navigation) {
    return (
        <Icon type="add_friend" size={25} color={colors.white} onPress={() => navigation.navigate("Test")} />
    )
}

const getFriends = (userId, search, offset) => {
    let queryString = `/friends/for-user?user_id=${userId}&search=${search}&offset=${offset}&limit=10`

    return Network.GET(queryString)
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
});


export default FriendScreen;