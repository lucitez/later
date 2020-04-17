import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { userId } from '../util/constants';
import Network from '../util/Network';
import { colors } from '../assets/colors';
import Icon from '../components/Icon';
import BottomSheet from '../components/BottomSheet';
import BottomSheetContainer from '../components/BottomSheetContainer';
import ButtonGroup from '../components/ButtonGroup';
import Button from '../components/Button';

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
        height: '25%',
        backgroundColor: colors.primary,
        paddingTop: 50,
        justifyContent: 'flex-start',
    },
    profileContainer: {
        flexDirection: 'row',
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