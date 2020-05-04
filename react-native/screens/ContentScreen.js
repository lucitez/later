import React, { useEffect, useState } from 'react';
import { StyleSheet, View, Alert, SafeAreaView, FlatList } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { Header, SearchBar, Icon, Divider, BackIcon } from '../components/common';
import { ContentFilter, ContentPreview } from '../components/content';
import { EditTagBottomSheet, ForwardBottomSheet } from '../components/modals';
import { ContentScreenPlaceholder } from '../components/placeholders';

const LIMIT = 20

function ContentScreen({ navigation, route, kind }) {
    let tag = route && route.params.tag

    const [content, setContent] = useState([])
    const [refreshing, setRefreshing] = useState(false)
    const [search, setSearch] = useState('')
    const [filter, setFilter] = useState({})
    const [offset, setOffset] = useState(0)
    const [limitReached, setLimitReached] = useState(false)

    // For modal
    const [selectedContent, setSelectedContent] = useState(null)
    const [saveBottomSheetVisible, setSaveBottomSheetVisible] = useState(false)
    const [editTagBottomSheetVisible, setEditTagBottomSheetVisible] = useState(false)
    const [forwardBottomSheetVisible, setForwardBottomSheetVisible] = useState(false)

    const updateContent = (contentUpdateFunc) => {
        setRefreshing(true)
        getContent(kind, search, offset, filter, tag)
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
        setContent([])
        updateContent(content => setContent(content))
    }, [filter, search])

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

    const onSave = tag => {
        saveContent(selectedContent, tag)
            .then(setContent(content.filter(c => c.id != selectedContent.id)))
            .catch(error => Alert.alert(error))
    }

    const onUpdateTag = newTag => {
        let prevContent = selectedContent

        if (prevContent.tag != newTag) {
            setContent(transformContentTag(content, prevContent.id, newTag))
            updateTag(prevContent.id, newTag)
                .then()
                .catch(() => setContent(transformContentTag(content, prevContent.id, prevTag)))
        }
    }

    const onDeletePress = toDelete => {
        Alert.alert('Confirm delete', null, [
            {
                text: 'Cancel',
                onPress: () => null,
                style: 'cancel'
            },
            {
                text: 'Confirm',
                onPress: () => {
                    deleteContent(toDelete.id)
                        .then(() => setContent(content.filter(c => c.id != toDelete.id)))
                        .catch(err => console.log(err))
                }
            }
        ])
    }

    const renderContent = ({ item }) => (
        <ContentPreview
            content={item}
            includeFooter={true}
            kind={kind}
            linkActive={!editTagBottomSheetVisible && !forwardBottomSheetVisible && !saveBottomSheetVisible}
            onTagPress={tag => navigation.navigate('Tag', { tag })}
            onDeletePress={(content) => onDeletePress(content)}
            onForwardPress={content => {
                setSelectedContent(content)
                setForwardBottomSheetVisible(true)
            }}
            onSavePress={content => {
                setSelectedContent(content)
                setSaveBottomSheetVisible(true)
            }}
            onEditTagPress={content => {
                setSelectedContent(content)
                setEditTagBottomSheetVisible(true)
            }}
        />
    )

    // TODO splash for when there is no content
    // TODO swipe to save/delete
    return (
        <SafeAreaView style={styles.container}>
            <Header
                title="Later"
                leftIcon={kind != 'home' && <BackIcon navigation={navigation} color={colors.white} />}
                rightIcon={kind == 'home' && <Icon
                    type='save'
                    size={25}
                    color={colors.white}
                    onPress={() => navigation.navigate('Saved')}
                />}
            />
            {kind == 'saved' && <>
                <SearchBar onChange={value => setSearch(value)} placeholder='Search by title or tag...' />
                <ContentFilter onChange={(filter) => setFilter(filter)} />
            </>}
            <View style={styles.contentContainer}>
                {!refreshing && content.length == 0 && <ContentScreenPlaceholder navigation={navigation} kind={kind} />}
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
            <EditTagBottomSheet
                isVisible={saveBottomSheetVisible}
                onSubmit={onSave}
                onHide={() => setSaveBottomSheetVisible(false)}
            />
            <EditTagBottomSheet
                isVisible={editTagBottomSheetVisible}
                onSubmit={onUpdateTag}
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

const getContent = (kind, search, offset, contentFilter, tag) => {
    let homeParams = { limit: LIMIT, offset }

    let savedParams = {
        saved: true,
        search,
        limit: LIMIT,
        offset,
        ...contentFilter
    }

    switch (kind) {
        case 'home':
            return Network.GET('/user-content/filter', homeParams)
        case 'saved':
            return Network.GET('/user-content/filter', savedParams)
        case 'byTag':
            return Network.GET('/user-content/by-tag', { tag })
    }
}

const updateTag = (contentId, tag) => {
    let params = {
        id: contentId,
        tag: tag
    }
    return Network.PUT('/user-content/update', params)
}

const saveContent = (content, tag) => {
    let params = {
        id: content.id,
        tag: tag
    }

    return Network.PUT('/user-content/save', params)
}

const deleteContent = id => {
    return Network.PUT('/user-content/delete', { id })
}

const transformContentTag = (content, contentId, tag) => (
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
        backgroundColor: colors.lightGray,
        flexBasis: 0,
        flexGrow: 1,
    }
});

export default ContentScreen