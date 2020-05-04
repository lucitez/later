import React from 'react';
import { StyleSheet, SafeAreaView } from 'react-native';
import { CommonActions } from '@react-navigation/native';
import { Header, BackIcon } from '../components/common';
import Network from '../util/Network';
import { colors } from '../assets/colors';
import { ShareFriendSelector } from '../components/share/index';

function SendShareScreen({ navigation, route }) {
    let preview = route.params.contentPreview

    const onSend = (userIds, callback) => {
        sendShares(preview.url, preview.contentType, userIds)
            .then(() => {
                callback()
                navigation.dispatch(
                    CommonActions.reset({
                        index: 3, // index of the profile tab
                        routes: [
                            { name: 'Home' },
                        ]
                    })
                )
            })
            .catch(err => {
                console.error(err)
                callback()
            })
    }

    return (
        <SafeAreaView style={styles.container}>
            <Header title='Share' leftIcon={<BackIcon navigation={navigation} />} />
            <ShareFriendSelector onSend={onSend} contentPreview={preview} />
        </SafeAreaView >
    );
}

const sendShares = (url, contentType, userIds) => {
    let body = {
        url,
        contentType,
        recipientUserIds: userIds,
    }
    return Network.POST('/shares/new', body)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.primary,
    },
});

export default SendShareScreen;