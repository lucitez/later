import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import Colors from '../assets/colors';

function Header(props) {
    return (
        <View style={styles.headerContainer}>
            <Text style={styles.header}>{props.name}</Text>
        </View>
    );
}

const styles = StyleSheet.create({
    headerContainer: {
        height: '100%',
        width: '100%',
        justifyContent: 'center',
        alignItems: 'center'
    },
    header: {
        color: Colors.white,
        fontSize: 20,
        fontWeight: 'bold',
    },
});

export default Header