import React from 'react';
import { StyleSheet, Text, View, Image } from 'react-native';

function ContentPreview(props) {

    return (
        <View style={styles.contentContainer}>
            <View style={styles.imageContainer}>
                <Image style={styles.thumb} source={{ uri: props.content.image_url }} />
            </View>
            <View style={styles.detailsContainer}>
                <View style={styles.titleAndDescriptionContainer}>
                    <Text style={styles.title} numberOfLines={2}>{props.content.title}</Text>
                    <Text style={styles.description} numberOfLines={1}>{props.content.description}</Text>
                </View>
                <View style={styles.tagContainer}>
                    <Text>{props.content.tag}</Text>
                </View>
                <View style={styles.sentByContainer}>
                    <Text>Recommended by {props.content.sent_by_username}</Text>
                </View>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    contentContainer: {
        display: 'flex',
        height: 120,
        width: '100%',
        flexDirection: 'row',
        justifyContent: 'space-around',
        marginTop: 5,
        marginBottom: 5
    },
    imageContainer: {
        display: 'flex',
        height: 120,
        width: 120,
        justifyContent: 'center',
        alignItems: 'center',
        padding: 5
    },
    thumb: {
        height: '100%',
        width: '100%',
        borderRadius: 5,
    },
    detailsContainer: {
        display: 'flex',
        flex: 1,
        flexDirection: 'column',
        padding: 5,
    },
    tagContainer: {
        position: 'absolute',
        right: 10,
        top: 10
    },
    titleAndDescriptionContainer: {
        flex: 2,
        marginRight: '20%',
    },
    title: {
        fontWeight: 'bold',
        overflow: 'hidden',
        fontSize: 16,
    },
    description: {
        fontSize: 12,
    },
    sentByContainer: {
        flex: 1,
        flexDirection: 'row',
        alignItems: 'flex-end'
    },
});

export default ContentPreview