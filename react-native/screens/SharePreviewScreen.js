import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Text, TouchableOpacity, SafeAreaView, ActivityIndicator } from 'react-native';
import { Header, SearchBar } from '../components/common';
import Network from '../util/Network';
import { ContentPreview } from '../components/content';
import { colors } from '../assets/colors';

function SharePreviewScreen({ navigation, route }) {
    const [url, setUrl] = useState('')
    const [loading, setLoading] = useState(false)
    const [contentPreview, setContentPreview] = useState(null)
    const [sent, setSent] = useState(false)
    const [err, setErr] = useState('')

    useEffect(() => {
        setLoading(true)
        setContentPreview(null)
        setErr(null)

        if (url.length > 0 && validUrl(url)) {
            getContentPreview(url)
                .then(contentPreview => {
                    setContentPreview(contentPreview)
                    setLoading(false)
                })
                .catch(_ => {
                    setErr('Please provide a valid URL')
                    setLoading(false)
                })
        } else if (url.length > 0 && !validUrl(url)) {
            setErr('Please provide a valid URL')
            setLoading(false)
        } else {
            setLoading(false)
        }
    }, [url])

    useEffect(() => {
        if (route.params && route.params.success) {
            setSent(true)
            setTimeout(() => { setSent(false) }, 2000)
        }
    }, [route.params])

    // TODO add content type radio
    return (
        <SafeAreaView style={styles.container}>
            <Header title='Share' />
            <SearchBar
                clear={sent}
                onChange={value => setUrl(value)}
                iconName='paste'
                autoFocus={true}
                returnKeyType={contentPreview ? 'next' : 'default'}
                placeholder='Enter URL...'
            />
            <View style={styles.contentContainer}>
                {loading && <View style={styles.noPreviewContainer}><ActivityIndicator size='small' /></View>}
                {err && <View style={styles.noPreviewContainer}><Text>{err}</Text></View>}
                {contentPreview &&
                    <View>
                        <View style={styles.contentPreviewContainer}>
                            <ContentPreview content={contentPreview} onDotPress={() => null} />
                        </View>
                        <View style={styles.footerContainer}>
                            <TouchableOpacity onPress={() => {
                                navigation.navigate('Send Share', { contentPreview: contentPreview, previousScreen: 'Share' })
                            }}>
                                <View style={styles.nextButtonContainer}>
                                    <Text style={styles.nextButton}>Next</Text>
                                </View>
                            </TouchableOpacity>
                        </View>
                    </View>}
            </View>
        </SafeAreaView>
    );
}

const validUrl = url => {
    let pattern = /^(https?|chrome):\/\/[^\s$.?#].[^\s]*$/

    return url.match(pattern)
}

const getContentPreview = url => {
    let params = {
        url: url
    }
    let queryString = `/content/preview`
    return Network.GET(queryString, params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.primary,
    },
    contentContainer: {
        flexGrow: 1,
        backgroundColor: colors.lightGray,
    },
    noPreviewContainer: {
        marginTop: 10,
        width: '100%',
        alignItems: 'center',
    },
    contentPreviewContainer: {
        backgroundColor: colors.white,
        margin: 5,
        borderRadius: 5,
    },
    footerContainer: {
        width: '100%',
        flexDirection: 'row',
        justifyContent: 'flex-end',
        paddingRight: 10,
    },
    nextButtonContainer: {
        backgroundColor: colors.primary,
        padding: 7,
        paddingLeft: 10,
        paddingRight: 10,
        borderRadius: 5,
    },
    nextButton: {
        color: colors.white,
        fontSize: 16,
    }
});

export default SharePreviewScreen;