import React, { useEffect, useState } from 'react';
import { StyleSheet, View, Text, SafeAreaView } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { Header, SearchBar, Button, BackIcon } from '../components/common';
import { ContentFilter, ContentGroup2 } from '../components/content';
import { ButtonBottomSheet, EditTagBottomSheet } from '../components/modals';

function SavedScreen({ navigation }) {
    const [content, setContent] = useState([])
    const [search, setSearch] = useState('')
    const [filter, setFilter] = useState({})
    const [loading, setLoading] = useState(true)

    const [bottomSheetVisible, setBottomSheetVisible] = useState(false)
    const [editTagBottomSheetVisible, setEditTagBottomSheetVisible] = useState(false)
    const [selectedContent, setSelectedContent] = useState(null)

    useEffect(() => {
        setLoading(true)
        getContent(search, filter)
            .then(content => {
                setContent(content)
                setLoading(false)
            })
            .catch(error => console.error(error))
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
            <Header name="Saved" leftIcon={<BackIcon navigation={navigation} color={colors.white} />} />
            <SearchBar
                onChange={value => setSearch(value)}
                placeholder='Search...'
            />
            <ContentFilter onChange={(filter) => setFilter(filter)} />
            <View style={styles.contentContainer}>
                <ContentGroup2
                    type='save'
                    contents={content}
                    linkActive={!bottomSheetVisible && !editTagBottomSheetVisible}
                    onDotPress={content => {
                        setSelectedContent(content)
                        setContentBottomSheetVisible(true)
                    }}
                    onTagPress={tag => navigation.navigate('Tag', { tag })}
                />
            </View>
            <ButtonBottomSheet isVisible={bottomSheetVisible} onHide={() => setBottomSheetVisible(false)}>
                <Button
                    theme='primary'
                    name='Forward'
                    size='medium'
                    onPress={() => {
                        setBottomSheetVisible(false)
                        navigation.navigate('Forward', { contentPreview: content, previousScreen: 'Home' })
                    }}
                />
                <Button
                    theme='primary'
                    name='Edit Tag'
                    size='medium'
                    onPress={() => {
                        setBottomSheetVisible(false)
                        setTimeout(() => { setEditTagBottomSheetVisible(true) }, 400)
                    }}
                />
                <Button
                    theme='light'
                    name='Cancel'
                    size='medium'
                    onPress={() => setBottomSheetVisible(false)}
                />
            </ButtonBottomSheet>

            <EditTagBottomSheet
                isVisible={editTagBottomSheetVisible}
                onSubmit={onUpdateTag}
                onHide={() => setSaveContentBottomSheetVisible(false)}
            />

        </SafeAreaView>
    );
}

const getContent = (search, contentFilter) => {
    let params = {
        saved: true,
        search: search,
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