
import React, { useState, useEffect } from 'react'
import { StyleSheet, SafeAreaView } from 'react-native'
import BottomSheet from './BottomSheet'
import { ShareFriendSelector } from '../share/index'
import { colors } from '../../assets/colors';
import Network from '../../util/Network';

export default function ForwardBottomSheet({ selectedContent, isVisible, onHide }) {
    const [visible, setVisible] = useState(isVisible)

    useEffect(() => { setVisible(isVisible) }, [isVisible])

    const onSend = (userIds, callback) => {
        forwardContent(selectedContent.contentId, userIds)
            .then(() => {
                callback()
                onHide()
            })
            .catch(err => {
                console.error(err)
                callback()
                onHide()
            })
    }

    return (
        <BottomSheet
            backdropOpacity={0.5}
            visible={visible}
            onHide={() => onHide()}
            avoidKeyboard={true}
        >
            <SafeAreaView style={styles.container}>
                <ShareFriendSelector
                    onSend={onSend}
                    contentPreview={{ ...selectedContent, createdAt: null }}
                    onSearchCancel={() => onHide()}
                />
            </SafeAreaView >

        </BottomSheet>
    )
}

const forwardContent = (contentId, userIds) => {
    let body = {
        recipientUserIds: userIds,
        contentId
    }
    return Network.POST('/shares/forward', body)
}

const styles = StyleSheet.create({
    container: {
        height: '85%',
        backgroundColor: colors.primary,
    },
})