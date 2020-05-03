import React, { useState, useEffect } from 'react'
import { View, Text, StyleSheet, Image } from 'react-native'
import { colors } from '../../assets/colors'
import { Link } from '../common'

export default function ContentMessage({ title, url, imageUrl, fromMe }) {

    const [imageAR, setImageAR] = useState(null)

    useEffect(() => {
        if (imageUrl) {
            Image.getSize(imageUrl, (width, height) => setImageAR(width / height), _ => null)
        }
    })

    return (
        <View style={styles.container}>
            <Link url={url} active={true}>

                <View style={[styles.detailsContainer, { backgroundColor: fromMe ? colors.blue : colors.darkGray }]}>
                    {imageAR &&
                        <Image
                            style={[styles.image, { aspectRatio: imageAR }]}
                            source={imageUrl ? { uri: imageUrl } : {}}
                        />
                    }
                    <View style={styles.titleContainer}>
                        <Text numberOfLines={2} style={styles.title}>{title}</Text>
                    </View>
                </View>
            </Link>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        width: '60%',
        margin: 5,
        marginTop: 10,
    },
    detailsContainer: {
        backgroundColor: colors.darkGray,
        borderRadius: 10,
    },
    image: {
        width: '100%',
        borderTopLeftRadius: 10,
        borderTopRightRadius: 10,
    },
    titleContainer: {
        padding: 10,
    },
    title: {
        color: colors.white,
        fontWeight: 'bold',
    }
})