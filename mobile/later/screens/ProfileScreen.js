import React, { useState, useEffect, useContext } from 'react';
import { useDispatch } from 'react-redux';
import { StyleSheet, View, AsyncStorage } from 'react-native';
import Network from '../util/Network';
import { colors } from '../assets/colors';
import { Button, TabBar, Icon } from '../components/common';
import { ButtonBottomSheet } from '../components/modals';
import { createMaterialTopTabNavigator } from '@react-navigation/material-top-tabs'
import ConversationScreen from './ConversationsScreen';
import NotificationsScreen from './NotificationsScreen';
import { UserDetails } from '../components/user';
import { AuthContext } from '../context'
import * as actions from '../actions'

const ProfileTab = createMaterialTopTabNavigator()

const clearRefreshToken = async () => {
    try {
        await AsyncStorage.removeItem('refresh_token')
    } catch (e) {
        console.error(e)
    }
}

function ProfileScreen({ navigation, route }) {
    const dispatch = useDispatch()
    const { signOut } = useContext(AuthContext)

    const [user, setUser] = useState(null)
    const [editVisible, setEditVisible] = useState(false)

    useEffect(() => {
        getProfile()
            .then(u => setUser(u))
            .catch(err => console.error(err))
    }, [])

    useEffect(() => {
        if (route.params) {
            setUser({ ...user, ...route.params.newUserData })
        }
    }, [route.params])

    const onSignOut = () => {
        dispatch(actions.clearTokens())
        clearRefreshToken()
        signOut()
    }

    return (
        <View style={styles.container}>
            <View style={styles.topContainer}>
                <View style={styles.topContentContainer}>
                    <View style={styles.userInfoContainer}>
                        {user && <UserDetails user={user} />}
                    </View>
                    <View style={styles.gearIconContainer}>
                        <Icon type='gear' size={25} color={colors.white} onPress={() => setEditVisible(true)} />
                    </View>
                </View>

            </View>
            <View style={styles.tabContainer}>
                <ProfileTab.Navigator initialRouteName='Conversations' tabBar={props => <TabBar {...props} />}>
                    <ProfileTab.Screen name='Conversations' component={ConversationScreen} />
                    <ProfileTab.Screen name='Notifications' component={NotificationsScreen} />
                </ProfileTab.Navigator>
            </View>
            <ButtonBottomSheet
                isVisible={editVisible}
                onHide={() => setEditVisible(false)}
            >
                <Button theme='primary' name='Edit Profile' size='medium' onPress={() => {
                    setEditVisible(false)
                    navigation.navigate('Edit', { user: user })
                }} />
                <Button theme='light' name='Log Out' size='medium' onPress={() => {
                    onSignOut()
                    setEditVisible(false)
                }} />
            </ButtonBottomSheet>
        </View >
    )
}

const getProfile = () => {
    return Network.GET('/users/profile')
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
    topContainer: {
        height: '22%',
        justifyContent: 'flex-end',
        backgroundColor: colors.primary,
        paddingBottom: 5,
    },
    topContentContainer: {
        flexDirection: 'row',
    },
    tabContainer: {
        flexGrow: 1,
        flexBasis: 0,
    },
    userInfoContainer: {
        flexGrow: 1,
        flexBasis: 0
    },
    gearIconContainer: {
        height: '100%',
        paddingRight: 15,
        paddingTop: 10,
    }
})

export default ProfileScreen