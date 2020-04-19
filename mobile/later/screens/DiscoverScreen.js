import React, { useState, useEffect } from 'react'
import { StyleSheet, View, Text } from 'react-native'
import { colors } from '../assets/colors'
import { SearchBar } from '../components/common'
import { userId } from '../util/constants';
import Network from '../util/Network'
import UserGroup from '../components/user/UserGroup'

function DiscoverScreen({ navigation }) {
    const [users, setUsers] = useState([])
    const [search, setSearch] = useState('')
    const [offset, setOffset] = useState(0)

    useEffect(() => {
        setOffset(0)
        if (search == '') {
            setUsers([])
        } else {
            searchUsers(search, 0)
                .then(users => setUsers(filterUsers(users)))
                .catch(err => console.error(err))
        }
    }, [search])

    useEffect(() => {
        if (offset > 0)
            searchUsers(search, offset)
                .then(nextUsers => setUsers(users.concat(filterUsers(nextUsers))))
                .catch(err => console.error(err))
    }, [offset])

    return (
        <View style={styles.container}>
            <View style={styles.header}>
                <SearchBar onChange={search => setSearch(search)} showCancelOnKeyboardActive={true} />
            </View>
            <UserGroup type='user' users={users} onPress={userId => navigation.navigate('User', { userId })} />
        </View>
    )
}

const searchUsers = (search, offset) => {
    let params = { search, offset }

    return Network.GET('/users/search', params)
}

const filterUsers = users => {
    return users.filter(user => user.id != userId)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray
    },
    header: {
        paddingTop: 50,
        paddingBottom: 10,
        backgroundColor: colors.primary,
    }
})

export default DiscoverScreen