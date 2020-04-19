import React, { useState, useEffect } from 'react';
import { StyleSheet, View } from 'react-native';
import { userId } from '../util/constants';
import Network from '../util/Network';
import { colors } from '../assets/colors';
import { Button, ButtonGroup, TabBar, Icon } from '../components/common';
import { BottomSheet, BottomSheetContainer } from '../components/modals';
import { createMaterialTopTabNavigator } from '@react-navigation/material-top-tabs'
import ConversationScreen from './ConversationsScreen';
import NotificationsScreen from './NotificationsScreen';
import { UserDetails } from '../components/user';

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