import React from 'react';
import { StyleSheet, View } from 'react-native';
import Colors from '../assets/colors';

function Divider(props) {

    return (
        <View style={styles.divider}></View>
    );
}

const styles = StyleSheet.create({
    divider: {
        height: 1,
        width: '100%',
        backgroundColor: Colors.black
    }
});

export default Divider