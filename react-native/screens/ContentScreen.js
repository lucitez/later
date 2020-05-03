import React, { useEffect, useState } from 'react';
import { StyleSheet, View, Alert, SafeAreaView, FlatList } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { Header, Icon, Button, Divider } from '../components/common';
import { ContentPreview } from '../components/content';
import { ButtonBottomSheet, EditTagBottomSheet } from '../components/modals';
import ForwardBottomSheet from '../components/modals/ForwardBottomSheet';

const LIMIT = 20

function ContentScreen({ navigation }) {
    const [content, setContent] = useState([])
    const [refreshing, setRefreshing] = useState(false)
    const [offset, setOffset] = useState(0)
    const [limitReached, setLimitReached] = useState(false)

    // For modal
    const [selectedContent, setSelectedContent] = useState(null)
    const [contentBottomSheetVisible, setContentBottomSheetVisible] = useState(false)
    const [editTagBottomSheetVisible, setEditTagBottomSheetVisible] = useState(false)
    const [forwardBottomSheetVisible, setForwardBottomSheetVisible] = useState(false)

    const updateContent = (contentUpdateFunc) => {
        setRefreshing(true)
        getContent(offset)
            .then(content => {
                if (content.length < LIMIT) {
                    setLimitReached(true)
                }
                setOffset(offset + content.length)
                contentUpdateFunc(content)
                setRefreshing(false)
            })
            .catch(err => {
                console.error(err)
                setRefreshing(false)
            })
    }


    useEffect(() => {
        updateContent(content => setContent(content))
    }, [])

    const onEndReached = () => {
        if (!limitReached) {
            updateContent(nextPage => setContent(content.concat(nextPage)))
        }
    }

    const onRefresh = () => {
        setOffset(0)
        setLimitReached(false)
        updateContent(content => setContent(content))
    }

    const onSave = (tag) => {
        saveContent(selectedContent, tag)
            .then(setContent(content.filter(c => c.id != selectedContent.id)))
            .catch(error => Alert.alert(error))
    }

    const renderContent = ({ item }) => (
        <ContentPreview
            content={item}
            linkActive={!contentBottomSheetVisible && !editTagBottomSheetVisible}
            onDotPress={content => {
                setSelectedContent(content)
                setContentBottomSheetVisible(true)
            }}
        />
    )

    // TODO splash for when there is no content
    // TODO allow them to delete
    // TODO swipe to save/delete
    return (
        <SafeAreaView style={styles.container}>
            <Header
                title="Later"
                rightIcon={<Icon
                    type='save'
                    size={25}
                    color={colors.white}
                    onPress={() => navigation.navigate('Saved')}
                />}
            />
            <View style={styles.contentContainer}>
                <FlatList
                    data={content}
                    onRefresh={onRefresh}
                    refreshing={refreshing}
                    onEndReached={onEndReached}
                    onEndReachedThreshold={0.1}
                    renderItem={renderContent}
                    ItemSeparatorComponent={Divider}
                />
            </View>
            <ButtonBottomSheet isVisible={contentBottomSheetVisible} onHide={() => setContentBottomSheetVisible(false)}>
                <Button theme='primary' name='Forward' size='medium' onPress={() => {
                    setContentBottomSheetVisible(false)
                    setTimeout(() => { setForwardBottomSheetVisible(true) }, 400)
                    // navigation.navigate('Forward', { contentPreview: selectedContent, previousScreen: 'Home' })
                }} />
                <Button theme='primary' name='Save' size='medium' onPress={() => {
                    setContentBottomSheetVisible(false)
                    setTimeout(() => { setEditTagBottomSheetVisible(true) }, 400)
                }} />
                <Button theme='light' name='Cancel' size='medium' onPress={() => setContentBottomSheetVisible(false)} />
            </ButtonBottomSheet>
            <EditTagBottomSheet
                isVisible={editTagBottomSheetVisible}
                onSubmit={onSave}
                onHide={() => setEditTagBottomSheetVisible(false)}
            />
            <ForwardBottomSheet
                isVisible={forwardBottomSheetVisible}
                selectedContent={selectedContent}
                onHide={() => setForwardBottomSheetVisible(false)}
            />
        </SafeAreaView>
    );
}

const getContent = (offset) => {
    let params = { limit: LIMIT, offset }
    return Network.GET(`/user-content/filter`, params)
}

const saveContent = (content, tag) => {
    let params = {
        id: content.id,
        tag: tag
    }

    return Network.PUT(`/user-content/save`, params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.primary,
    },
    contentContainer: {
        backgroundColor: colors.lightGray,
        flexBasis: 0,
        flexGrow: 1,
    }
});

export default ContentScreen