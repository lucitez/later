import React, { useEffect, useState } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import Header from '../components/Header';
import { userId } from '../util/constants';
import ContentFilter from '../components/ContentFilter';
import ContentGroup from '../components/ContentGroup';
import SearchBar from '../components/SearchBar';

function ContentScreen({ navigation }) {
    const [content, setContent] = useState([])
    const [search, setSearch] = useState('')
    const [filter, setFilter] = useState({})
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        setLoading(true)
        getContent(search, filter)
            .then(content => {
                setContent(content)
                setLoading(false)
            })
            .catch(error => console.error(error))
    }, [filter, search])

    return (
        <View style={styles.container}>
            <Header name="Later" />
            <SearchBar
                onChange={value => setSearch(value)}
                placeholder='Search...'
            />
            <ContentFilter onChange={(filter) => setFilter(filter)} />
            <View style={styles.contentContainer}>
                {
                    loading ?
                        <View style={{ width: '100%', alignItems: 'center', paddingTop: 10 }}>
                            <Text>Loading...</Text>
                        </View>
                        :
                        <ContentGroup content={content} onForward={(content) => {
                            navigation.navigate('Share', {
                                screen: 'Send Share',
                                params: { contentPreview: content }
                            })
                        }} />
                }
            </View>
        </View>
    );
}

const getContent = (search, contentFilter) => {
    let params = {
        userId: userId,
        search: search,
        ...contentFilter
    }
    return Network.GET(`/user-content/filter`, params)
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