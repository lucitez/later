import React, { useEffect, useState } from 'react';
import { StyleSheet, View, Text, Alert } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import Header from '../components/Header';
import { userId } from '../util/constants';
import ContentGroup from '../components/ContentGroup';
import Icon from '../components/Icon';

function ContentScreen({ navigation }) {
    const [content, setContent] = useState([])
    const [offset, setOffset] = useState(0)
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        setLoading(true)
        getContent()
            .then(content => {
                setContent(content)
                setLoading(false)
            })
            .catch(error => console.error(error))
    }, [offset])

    const onSave = (savedContent, tag) => {
        SavedContent(savedContent, tag)
            .then(setContent(content.filter(c => c.id != savedContent.id)))
            .catch(error => Alert.alert(error))
    }

    const saveIcon = (
        <Icon
            type='save'
            size={25}
            color={colors.white}
            onPress={() => navigation.navigate('Saved')}
        />
    )

    return (
        <View style={styles.container}>
            <Header name="Later" rightIcon={saveIcon} />
            <View style={styles.contentContainer}>
                {
                    content.length == 0 ?
                        <View style={{ width: '100%', alignItems: 'center', paddingTop: 15 }}>
                            <Text style={{ textAlign: 'center' }}>Check out the discover page for more content!</Text>
                        </View>
                        :
                        null
                }
                <ContentGroup
                    type='home'
                    content={content}
                    onForward={content => navigation.navigate('Forward', { contentPreview: content, previousScreen: 'Home' })}
                    onSave={onSave}
                />
                {
                    loading ?
                        <View style={{ width: '100%', alignItems: 'center', paddingTop: 10 }}>
                            <Text>Loading...</Text>
                        </View>
                        :
                        null

                }
            </View>
        </View>
    );
}

const getContent = () => {
    let params = {
        userId: userId
    }
    return Network.GET(`/user-content/filter`, params)
}

const SavedContent = (content, tag) => {
    let params = {
        id: content.id,
        tag: tag
    }

    return Network.PUT(`/user-content/save`, params)
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