import React from 'react';
import { StyleSheet, Text, View, TouchableOpacity } from 'react-native';
import { colors } from '../assets/colors';
import Icon from '../components/Icon';

function ShareWithFriendPreview(props) {
    let user = props.user
    let selected = user.selected

    return (
        <View style={styles.container}>
            <View style={styles.imageContainer}>
                <View style={styles.thumb}></View>
            </View>
            <View style={styles.detailsContainer}>
                <Text style={styles.name}>{user.firstName} {user.lastName}</Text>
                <Text style={styles.username}>@{user.username}</Text>
            </View>
            <View style={styles.chatContainer}>
                {
                    selected ?
                        <Icon type='check_filled' size={30} />
                        :
                        <Icon type='circle' size={30} />
                }
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        height: 60,
        paddingTop: 5,
        paddingBottom: 5,
        width: '100%',
        backgroundColor: colors.white
    },
    imageContainer: {
        width: 50,
        padding: 5,
        justifyContent: 'center',
        alignItems: 'center',
    },
    thumb: {
        height: '100%',
        width: '100%',
        borderRadius: 100,
        backgroundColor: 'coral',
    },
    detailsContainer: {
        height: '100%',
        flexGrow: 1,
        padding: 5,
        justifyContent: 'center',
    },
    name: {
        fontSize: 16,
        fontWeight: '500'
    },
    username: {
        fontSize: 14,
        fontWeight: '300'
    },
    chatContainer: {
        height: '100%',
        padding: 5,
        marginRight: 10,
        justifyContent: 'center',
        alignItems: 'center',
    },
});

export default ShareWithFriendPreview