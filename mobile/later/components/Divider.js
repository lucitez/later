import React from 'react';
import { StyleSheet, View } from 'react-native';
import Colors from '../assets/colors';

function Divider() {

    return (
        <View style={styles.divider}></View>
    );
}

const styles = StyleSheet.create({
    divider: {
        height: 0.5,
        alignSelf: 'center',
        width: '98%',
        backgroundColor: Colors.black
    }
});

export default Divider