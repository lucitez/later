import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { userId } from '../util/constants';
import Network from '../util/Network';
import { colors } from '../assets/colors';
import { Icon, Button, ButtonGroup, TabBar } from '../components/common';
import { BottomSheet, BottomSheetContainer } from '../components/modals';
import { createMaterialTopTabNavigator } from '@react-navigation/material-top-tabs'
import FriendScreen from './FriendScreen';
import ConversationScreen from './ConversationsScreen';
import NotificationsScreen from './NotificationsScreen';

const ProfileTab = createMaterialTopTabNavigator()

function ProfileScreen({ navigation, route }) {

    const [user, setUser] = useState(null)
    const [editVisible, setEditVisible] = useState(false)

    useEffect(() => {
        getUser()
            .then(u => setUser(u))
            .catch(err => console.error(err))
    }, [])

    useEffect(() => {
        if (route.params) {
            setUser({ ...user, ...route.params.newUserData })
        }
    }, [route.params])

    return (
        <View style={styles.container}>
            <View style={styles.topContainer}>
                {
                    user ?
                        <View style={styles.profileContainer}>
                            <View style={styles.leftContainer}>
                                <View style={styles.imageContainer} />
                                <View style={styles.tasteContainer}>
                                    <Text style={styles.taste}>28 taste</Text>
                                </View>
                            </View>
                            <View style={styles.infoContainer}>
                                <Text style={styles.name}>{user.firstName} {user.lastName}</Text>
                                <Text style={styles.username}>@{user.username}</Text>
                            </View>
                            <View style={styles.gearIconContainer}>
                                <Icon type='gear' size={25} color={colors.white} onPress={() => setEditVisible(true)} />
                            </View>
                        </View>
                        :
                        null
                }
            </View>
            <View style={styles.tabContainer}>
                <ProfileTab.Navigator initialRouteName='Conversations' tabBar={props => <TabBar {...props} />}>
                    <ProfileTab.Screen name='Conversations' component={ConversationScreen} />
                    <ProfileTab.Screen name='Notifications' component={NotificationsScreen} />
                </ProfileTab.Navigator>
            </View>
            <BottomSheet
                visible={editVisible}
                onHide={() => setEditVisible(false)}
            >
                <BottomSheetContainer>
                    <ButtonGroup theme='primary'>
                        <Button theme='primary' name='Edit Profile' size='medium' onPress={() => {
                            setEditVisible(false)
                            navigation.navigate('Edit', { user: user })
                        }} />
                        <Button theme='light' name='Log Out' size='medium' onPress={() => setEditVisible(false)} />
                    </ButtonGroup>
                </BottomSheetContainer>
            </BottomSheet>
        </View >
    )
}

const getUser = () => {
    let params = {
        id: userId
    }
    return Network.GET('/users/by-id', params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
    topContainer: {
        height: '22%',
        backgroundColor: colors.primary,
        paddingTop: 50,
        justifyContent: 'flex-start',
    },
    profileContainer: {
        flexDirection: 'row',
    },
    tabContainer: {
        flexGrow: 1,
        flexBasis: 0,
    },
    leftContainer: {
        width: '25%',
        alignItems: 'center',
        justifyContent: 'flex-start',
    },
    imageContainer: {
        height: 75,
        width: 75,
        borderRadius: 50,
        backgroundColor: colors.lightGray,
        alignItems: 'center',
        marginBottom: 10,
    },
    infoContainer: {
        justifyContent: 'center',
        flexGrow: 1,
    },
    taste: {
        color: colors.white,
    },
    name: {
        fontSize: 18,
        fontWeight: 'bold',
        color: colors.white
    },
    username: {
        color: colors.white
    },
    gearIconContainer: {
        paddingLeft: 10,
        paddingRight: 10,
    }
})

export default ProfileScreen