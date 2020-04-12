import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Keyboard, Text } from 'react-native';
import Header from '../components/Header';
import SearchBar from '../components/SearchBar';
import Network from '../util/Network';
import ContentPreview from '../components/ContentPreview';
import { colors } from '../assets/colors';

function ShareScreen({ navigation }) {
    const [url, setUrl] = useState('')
    const [contentPreview, setContentPreview] = useState(null)

    useEffect(() => {
        if (url.length > 0) {
            getContentPreview(url)
                .then(contentPreview => setContentPreview(contentPreview))
                .catch(error => console.error(error))
        } else if (url.length == 0) {
            setContentPreview(null)
        }
    }, [url])

    console.log(contentPreview)

    return (
        <View style={styles.container}>
            <Header name='Share' />
            <SearchBar
                onChange={value => setUrl(value)}
                iconName='paste'
                autoFocus={true}
                returnKeyType={contentPreview ? 'next' : 'default'}
                placeholder='Enter URL...'
            />
            {
                contentPreview ?
                    <View style={{ backgroundColor: colors.white, margin: 5, borderRadius: 5 }}><ContentPreview content={contentPreview} /></View> :
                    <View style={styles.noPreviewContainer}>
                        <Text>{url.length == 0 ? "Paste a URL to get started!" : "We could not generate a preview of your link"}</Text>
                    </View>
            }
        </View>
    );
}

const getContentPreview = url => {
    let queryString = `/content/preview?url=${url}`
    return Network.GET(queryString)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
    noPreviewContainer: {
        marginTop: 10,
        width: '100%',
        alignItems: 'center',
    }
});

export default ShareScreen;