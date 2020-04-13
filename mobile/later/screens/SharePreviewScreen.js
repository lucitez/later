import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Text, Button, TouchableOpacity } from 'react-native';
import Header from '../components/Header';
import SearchBar from '../components/SearchBar';
import Network from '../util/Network';
import ContentPreview from '../components/ContentPreview';
import { colors } from '../assets/colors';

function SharePreviewScreen({ navigation }) {
    const [url, setUrl] = useState('')
    const [loading, setLoading] = useState(false)
    const [contentPreview, setContentPreview] = useState(null)

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
                    <View>
                        <View style={styles.contentPreviewContainer}>
                            <ContentPreview content={contentPreview} />
                        </View>
                        <View style={styles.footerContainer}>
                            <TouchableOpacity onPress={() => {
                                navigation.navigate('Send Share', { contentPreview: contentPreview })
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
                                <Text>{url.length == 0 ? "Paste a URL to get started!" : "We could not generate a preview of your link"}</Text>
                        }
                    </View>
            }
        </View>
    );
}

const getContentPreview = url => {
    params = {
        url: url
    }
    let queryString = `/content/preview`
    return Network.GET(queryString, params)
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