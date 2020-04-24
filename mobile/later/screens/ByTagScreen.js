import React, { useEffect, useState } from 'react';
import { StyleSheet, View, Text, SafeAreaView } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { Header, SearchBar, BackIcon } from '../components/common';
import { ContentGroup } from '../components/content';

function ByTagScreen({ navigation, route }) {
    let tag = route.params.tag

    const [content, setContent] = useState([])
    const [loading, setLoading] = useState(true)

    const getContent = () => {
        Network.GET('/user-content/by-tag', { tag })
            .then(c => {
                setContent(c)
                setLoading(false)
            })
            .catch(err => {
                console.error(err)
                setLoading(false)
            })
    }

    useEffect(() => {
        getContent()
    }, [])

    const onUpdateTag = (updatedContent, newTag) => {
        let prevTag = updatedContent.tag

        if (prevTag != newTag) {
            setContent(updateContentTag(content, updatedContent.id, newTag))
            updateTag(updatedContent.id, newTag)
                .then(() => getContent())
                .catch(() => setContent(updateContentTag(content, updatedContent.id, prevTag)))
        }
    }

    return (
        <SafeAreaView style={styles.container}>
            <Header name={tag} leftIcon={<BackIcon navigation={navigation} />} />
            <View style={styles.contentContainer}>
                <ContentGroup
                    type='save'
                    content={content}
                    onForward={content => navigation.navigate('Forward', { contentPreview: content, previousScreen: 'Tag' })}
                    onUpdateTag={onUpdateTag}
                />
                {
                    loading ?
                        <View style={{ width: '100%', alignItems: 'center', paddingTop: 10 }}>
                            <Text>Loading...</Text>
                        </View>
                        :
                        content.length == 0 && Object.length == 0 ?
                            <View style={{ width: '100%', alignItems: 'center', paddingTop: 10 }}>
                                <Text>Your saved content shows up here</Text>
                            </View>
                            :
                            null
                }
            </View>
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
    searchContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 10
    },
    contentContainer: {
        flexGrow: 1,
        backgroundColor: colors.white,
    }
});

export default ByTagScreen