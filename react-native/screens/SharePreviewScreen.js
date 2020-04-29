import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Text, TouchableOpacity, SafeAreaView } from 'react-native';
import { Header, SearchBar } from '../components/common';
import Network from '../util/Network';
import { ContentPreview } from '../components/content';
import { colors } from '../assets/colors';

function SharePreviewScreen({ navigation, route }) {
    const [url, setUrl] = useState('')
    const [loading, setLoading] = useState(false)
    const [contentPreview, setContentPreview] = useState(null)
    const [sent, setSent] = useState(false)

    useEffect(() => {
        if (url.length > 0) {
            setLoading(true)
            getContentPreview(url)
                .then(contentPreview => {
                    setContentPreview(contentPreview)
                    setLoading(false)
                })
                .catch(() => setLoading(false))
        } else if (url.length == 0) {
            setLoading(false)
            setContentPreview(null)
        }
    }, [url])

    useEffect(() => {
        if (route.params && route.params.success) {
            setSent(true)
            setTimeout(() => { setSent(false) }, 2000)
        }
    }, [route.params])

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
                {
                    contentPreview ?
                        <View>
                            <View style={styles.contentPreviewContainer}>
                                <ContentPreview content={contentPreview} />
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
                        </View>
                        :
                        <View style={styles.noPreviewContainer}>
                            {
                                loading ?
                                    <Text>Retrieving data</Text>
                                    :
                                    <Text>
                                        {
                                            sent ? "Share successful!" :
                                                url.length == 0 ? "Paste a URL to get started" :
                                                    "We could not generate a preview of your link"
                                        }
                                    </Text>
                            }
                        </View>
                }
            </View>
        </SafeAreaView>
    );
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