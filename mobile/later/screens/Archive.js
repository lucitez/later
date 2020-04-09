import React, { useEffect, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import AntIcons from 'react-native-vector-icons/AntDesign';
import ContentGroup from '../components/ContentGroup';
import Header from '../components/Header';
import Colors from '../assets/colors';
import Network from '../util';

function ArchiveScreen() {
    const [content, setContent] = useState([])

    useEffect(() => {
        getContent('b6e05c09-0f62-4757-95f5-ea855adc4915')
            .then(response => response.json())
            .then(content => setContent(content))
            .catch(error => console.error(error))
    }, [])

    return (
        <View style={styles.container}>
            <View style={styles.headerContainer}>
                <View style={styles.filterContainer}>
                    <AntIcons name="filter" size={30} color={Colors.white} />
                </View>
                <View style={styles.tabContainer}>
                    <Header name="Archive" />
                </View>
                <View style={styles.searchContainer}>
                    <AntIcons name="search1" size={30} color={Colors.white} />
                </View>
            </View>
            <ContentGroup content={content} />
        </View>
    );
}

const getContent = (
    userId
) => {
    return fetch(`${Network.local}/user-content/filter?user_id=${userId}&archived=true`)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: Colors.white,
    },
    headerContainer: {
        backgroundColor: Colors.primary,
        height: '12%',
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: "flex-end",
    },
    filterContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 10
    },
    tabContainer: {
        flex: 3,
        paddingTop: '50%',
    },
    searchContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 10
    },
});

export default ArchiveScreen