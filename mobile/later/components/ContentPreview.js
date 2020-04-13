import React from 'react';
import { StyleSheet, Text, View, Linking, Alert, Image, TouchableWithoutFeedback } from 'react-native';
import Icon from '../components/Icon';
import Tag from '../components/Tag';
import { contentTypes } from '../assets/colors';

function ContentPreview(props) {

    let content = props.content

    return (
        <View style={styles.contentContainer}>
            <View style={styles.imageContainer}>
                <Image style={styles.thumb} source={content.imageUrl ? { uri: content.imageUrl } : {}} />
            </View>
            <View style={styles.detailsContainer}>
                <View style={styles.topDetailsContainer}>
                    <View style={styles.titleAndDescriptionContainer}>
                        <TouchableWithoutFeedback onPress={async () => {
                            // Checking if the link is supported for links with custom URL scheme.
                            const supported = await Linking.canOpenURL(content.url);

                            if (supported) {
                                // Opening the link with some app, if the URL scheme is "http" the web link should be opened
                                // by some browser in the mobile
                                await Linking.openURL(content.url);
                            } else {
                                Alert.alert(`Don't know how to open this URL: ${content.url}`);
                            }

                        }}>
                            <View>
                                <Text style={styles.title} numberOfLines={2}>{content.title}</Text>
                                <Text style={styles.description} numberOfLines={1}>{content.description}</Text>
                            </View>
                        </TouchableWithoutFeedback>
                    </View>
                    {
                        content.tag ?
                            <View style={styles.tagContainer}>
                                <Tag name={content.tag} />
                            </View>
                            : null
                    }
                </View>
                <View style={styles.bottomDetailsContainer}>
                    {
                        content.sentByUsername ?
                            <Text>Recommended by {content.sentByUsername}</Text> :
                            null
                    }
                    {
                        content.contentType ?
                            <View style={styles.iconContainer}>
                                <Icon type={content.contentType} size={25} color={contentTypes[content.contentType].color} />
                            </View> :
                            null
                    }

                </View>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    contentContainer: {
        flexDirection: 'row',
        height: 120,
        width: '100%',
        marginTop: 5,
        marginBottom: 5,
    },
    imageContainer: {
        width: 120,
        justifyContent: 'center',
        alignItems: 'center',
        padding: 5,
    },
    detailsContainer: {
        flexDirection: 'column',
        flexGrow: 1,
        padding: 5,
    },
    topDetailsContainer: {
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    tagContainer: {
        flex: 1,
    },
    titleAndDescriptionContainer: {
        flex: 4,
    },
    thumb: {
        height: '100%',
        width: '100%',
        borderRadius: 5,
    },
    title: {
        fontWeight: 'bold',
        fontSize: 16,
    },
    description: {
        fontSize: 12,
    },
    bottomDetailsContainer: {
        flexGrow: 1,
        flexDirection: 'row',
        alignItems: 'flex-end',
        justifyContent: 'space-between',
    },
    iconContainer: {
        paddingRight: 10,
        marginBottom: -5
    }
});

export default ContentPreview