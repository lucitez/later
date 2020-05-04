import React from 'react';
import { StyleSheet, View } from 'react-native';

function Divider() {

    return (
        <View style={styles.divider}></View>
    );
}

const styles = StyleSheet.create({
    divider: {
        height: 4,
    }
});

export default Divider