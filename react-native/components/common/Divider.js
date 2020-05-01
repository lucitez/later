import React from 'react';
import { StyleSheet, View } from 'react-native';
import { colors } from '../../assets/colors';

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
        backgroundColor: colors.black
    }
});

export default Divider