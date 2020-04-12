import React, { useEffect, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import ContentGroup from '../components/ContentGroup';
import Header from '../components/Header';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import ContentFilter from '../components/ContentFilter';

function ContentScreen(props) {
    const [allContent, setAllContent] = useState([])
    const [visibleContent, setVisibleContent] = useState([])
    const [filter, setFilter] = useState({
        'content_type': null
    })

    useEffect(() => {
        getContent('b6e05c09-0f62-4757-95f5-ea855adc4915', props.archived)
            .then(content => {
                setAllContent(content)
                setVisibleContent(content)
            })
            .catch(error => console.error(error))
    }, [])

    useEffect(() => setVisibleContent(filterContent(allContent, filter)), [filter])

    return (
        <View style={styles.container}>
            <Header name="Later" />
            <ContentFilter onChange={(filter) => setFilter(filter)} />
            <View style={styles.contentContainer}>
                <ContentGroup content={visibleContent} />
            </View>
        </View>
    );
}

const filterContent = (content, filter) => {
    if (filter['content_type'] != null) {
        return content.filter(c => c['content_type'] == filter['content_type'])
    }
    return content
}

const getContent = (userId, archived) => {
    let queryString = `/user-content/filter?user_id=${userId}`
    if (archived) {
        queryString += `&archived=true`
    }
    return Network.GET(queryString)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
    searchContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 10
    },
    contentContainer: {
        backgroundColor: colors.white,
        flexGrow: 1,
    }
});

export default ContentScreen