import React, { useEffect, useState } from 'react';
import { StyleSheet, View, SafeAreaView, FlatList } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { Header, Button, BackIcon } from '../components/common';
import { ContentPreview } from '../components/content';
import { ButtonBottomSheet, EditTagBottomSheet } from '../components/modals';

function ByTagScreen({ navigation, route }) {
    let tag = route.params.tag
    const [content, setContent] = useState([])

    // For modal
    const [selectedContent, setSelectedContent] = useState(null)
    const [bottomSheetVisible, setBottomSheetVisible] = useState(false)
    const [editTagBottomSheetVisible, setEditTagBottomSheetVisible] = useState(false)

    const renderContent = ({ item }) => (
        <ContentPreview
            content={item}
            linkActive={!bottomSheetVisible && !editTagBottomSheetVisible}
            onDotPress={content => {
                setSelectedContent(content)
                setBottomSheetVisible(true)
            }}
            onTagPress={() => null}
        />
    )

    useEffect(() => {
        Network.GET(`/user-content/by-tag`, { tag })
            .then(content => setContent(content))
            .catch(err => console.log(err))
    }, [])

    const onUpdateTag = newTag => {
        let prevContent = selectedContent

        if (prevContent.tag != newTag) {
            setContent(updateContentTag(content, prevContent.id, newTag))
            updateTag(prevContent.id, newTag)
                .then()
                .catch(() => setContent(updateContentTag(content, prevContent.id, prevTag)))
        }
    }

    return (
        <SafeAreaView style={styles.container}>
            <Header title={tag} leftIcon={<BackIcon navigation={navigation} color={colors.white} />} />
            <View style={styles.contentContainer}>
                <FlatList
                    data={content}
                    renderItem={renderContent}
                />
            </View>

            <ButtonBottomSheet isVisible={bottomSheetVisible} onHide={() => setBottomSheetVisible(false)}>
                <Button theme='primary' name='Forward' size='medium' onPress={() => {
                    setBottomSheetVisible(false)
                    navigation.navigate('Forward', { contentPreview: content, previousScreen: 'Home' })
                }} />
                <Button theme='primary' name='Edit Tag' size='medium' onPress={() => {
                    setBottomSheetVisible(false)
                    setTimeout(() => { setEditTagBottomSheetVisible(true) }, 400)
                }} />
                <Button theme='light' name='Cancel' size='medium' onPress={() => setBottomSheetVisible(false)} />
            </ButtonBottomSheet>

            <EditTagBottomSheet
                isVisible={editTagBottomSheetVisible}
                onSubmit={onUpdateTag}
                onHide={() => setSaveContentBottomSheetVisible(false)}
            />

        </SafeAreaView>
    );
}

const updateTag = (contentId, tag) => {
    let params = {
        id: contentId,
        tag: tag
    }
    return Network.PUT('/user-content/update', params)
}

const updateContentTag = (content, contentId, tag) => (
    content.map(c => (
        c.id == contentId ? { ...c, ['tag']: tag } : c
    ))
)

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.primary,
    },
    contentContainer: {
        flexGrow: 1,
        backgroundColor: colors.lightGray,
    }
});

export default ByTagScreen