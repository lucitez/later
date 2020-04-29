import React, { useState, useEffect } from 'react'
import { StyleSheet, View, SafeAreaView, TouchableOpacity } from 'react-native'
import { colors } from '../assets/colors'
import { SearchBar, Divider } from '../components/common'
import Network from '../util/Network'
import { UserDetailsPreview } from '../components/user'
import { FlatList } from 'react-native-gesture-handler'

const LIMIT = 20

function DiscoverScreen({ navigation }) {
    const [users, setUsers] = useState([])
    const [search, setSearch] = useState('')
    const [offset, setOffset] = useState(0)
    const [refreshing, setRefreshing] = useState(false)
    const [limitReached, setLimitReached] = useState(false)

    const updateUsers = (setFunc, offset = 0) => {
        setRefreshing(true)
        searchUsers(search, offset)
            .then(users => {
                if (users.length < LIMIT) {
                    setLimitReached(true)
                }
                setOffset(offset + users.length)
                setFunc(users)
                setRefreshing(false)
            })
            .catch(err => {
                console.error(err)
                setRefreshing(false)
            })
    }

    useEffect(() => {
        setOffset(0)
        setLimitReached(false)
        updateUsers(users => setUsers(users))
    }, [search])

    const onEndReached = () => {
        if (!limitReached) {
            updateUsers(nextPage => setUsers(users.concat(nextPage)), offset)
        }
    }

    const onRefresh = () => {
        setOffset(0)
        setLimitReached(false)
        updateUsers(users => setUsers(users))
    }

    const renderUser = ({ item }) => (
        <TouchableOpacity onPress={() => navigation.navigate('User', { userId: item.id })}>
            <UserDetailsPreview user={item} />
        </TouchableOpacity>
    )

    return (
        <SafeAreaView style={styles.container}>
            <View style={styles.header}>
                <SearchBar onChange={search => setSearch(search)} showCancelOnKeyboardActive={true} />
            </View>
            <View style={styles.contentContainer}>
                <FlatList
                    keyboardShouldPersistTaps='handled'
                    data={users}
                    renderItem={renderUser}
                    onEndReached={onEndReached}
                    onEndReachedThreshold={0.1}
                    refreshing={refreshing}
                    onRefresh={onRefresh}
                    keyExtractor={item => item.id}
                    ItemSeparatorComponent={Divider}
                />
            </View>
        </SafeAreaView>
    )
}

const searchUsers = (search, offset) => {
    let params = { search, offset, limit: LIMIT }

    return Network.GET('/users/search', params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.primary
    },
    header: {
        paddingTop: 10,
        paddingBottom: 10,
        backgroundColor: colors.primary,
    },
    contentContainer: {
        flexGrow: 1,
        backgroundColor: colors.lightGray,
    }
})

export default DiscoverScreen