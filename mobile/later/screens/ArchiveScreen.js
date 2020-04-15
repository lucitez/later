import React, { useEffect, useState } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import Header from '../components/Header';
import { userId } from '../util/constants';
import ContentFilter from '../components/ContentFilter';
import ContentGroup from '../components/ContentGroup';
import SearchBar from '../components/SearchBar';
import Icon from '../components/Icon';

function ArchiveScreen({ navigation }) {
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

    const backIcon = (
        <Icon
            type='back'
            size={25}
            color={colors.white}
            onPress={() => navigation.pop()}
        />
    )

    return (
        <View style={styles.container}>
            <Header name="Later" leftIcon={backIcon} />
            <SearchBar
                onChange={value => setSearch(value)}
                placeholder='Search...'
            />
            <ContentFilter onChange={(filter) => setFilter(filter)} />
            <View style={styles.contentContainer}>
                <ContentGroup
                    type='archive'
                    noContentMessage='Your archived content shows up here'
                    content={content}
                    onForward={content => navigation.navigate('Forward', { contentPreview: content })}
                />
                {
                    loading ?
                        <View style={{ width: '100%', alignItems: 'center', paddingTop: 10 }}>
                            <Text>Loading...</Text>
                        </View>
                        :
                        content.length == 0 && Object.length == 0 ?
                            <View style={{ width: '100%', alignItems: 'center', paddingTop: 10 }}>
                                <Text>Your archived content shows up here</Text>
                            </View>
                            :
                            null
                }
            </View>
        </View>
    );
}

const getContent = (search, contentFilter) => {
    let params = {
        archived: true,
        userId: userId,
        search: search,
        ...contentFilter
    }
    return Network.GET(`/user-content/filter`, params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.white,
    },
    searchContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 10
    },
    contentContainer: {
        flexGrow: 1,
    }
});

export default ArchiveScreen