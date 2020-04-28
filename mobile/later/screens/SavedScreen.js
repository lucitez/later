import React, { useEffect, useState } from 'react';
import { StyleSheet, View, SafeAreaView, FlatList } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { Header, SearchBar, Button, BackIcon } from '../components/common';
import { ContentFilter, ContentPreview } from '../components/content';
import { ButtonBottomSheet, EditTagBottomSheet } from '../components/modals';

const LIMIT = 20

function SavedScreen({ navigation }) {
    const [content, setContent] = useState([])
    const [search, setSearch] = useState('')
    const [filter, setFilter] = useState({})
    const [refreshing, setRefreshing] = useState(false)
    const [offset, setOffset] = useState(0)
    const [limitReached, setLimitReached] = useState(false)

    // For modal
    const [selectedContent, setSelectedContent] = useState(null)
    const [bottomSheetVisible, setBottomSheetVisible] = useState(false)
    const [editTagBottomSheetVisible, setEditTagBottomSheetVisible] = useState(false)

    const updateContent = (contentUpdateFunc) => {
        setRefreshing(true)
        getContent(search, filter, offset)
            .then(content => {
                if (content.length < LIMIT) {
                    setLimitReached(true)
                }
                setOffset(offset + content.length)
                contentUpdateFunc(content)
                setRefreshing(false)
            })
            .catch(err => {
                console.log(err)
                setRefreshing(false)
            })
    }

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

    const renderContent = ({ item }) => (
        <ContentPreview
            content={item}
            linkActive={!bottomSheetVisible && !editTagBottomSheetVisible}
            onDotPress={content => {
                setSelectedContent(content)
                setBottomSheetVisible(true)
            }}
            onTagPress={tag => navigation.navigate('Tag', { tag })}
        />
    )


    useEffect(() => {
        updateContent(content => setContent(content))
    }, [filter, search])

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
            <Header title="Saved" leftIcon={<BackIcon navigation={navigation} color={colors.white} />} />
            <SearchBar
                onChange={value => setSearch(value)}
                placeholder='Search by title or tag...'
            />
            <ContentFilter onChange={(filter) => setFilter(filter)} />
            <View style={styles.contentContainer}>
                <FlatList
                    data={content}
                    onRefresh={onRefresh}
                    refreshing={refreshing}
                    onEndReached={onEndReached}
                    onEndReachedThreshold={0.1}
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

const getContent = (search, contentFilter, offset) => {
    let params = {
        saved: true,
        search,
        limit: LIMIT,
        offset,
        ...contentFilter
    }
    return Network.GET(`/user-content/filter`, params)
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
        backgroundColor: colors.white,
    }
});

export default SavedScreen