import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { colors } from '../../assets/colors';


function Header({ leftIcon, titleComponent, title, rightIcon }) {
    return (
        <View style={styles.container}>
            <View style={styles.leftIconContainer}>
                {leftIcon ? leftIcon : null}
            </View>
            <View style={styles.headerTitleContainer}>
                {titleComponent ? titleComponent : <Text style={styles.title}>{title.toUpperCase()}</Text>}
            </View>
            <View style={styles.rightIconContainer}>
                {rightIcon ? rightIcon : null}
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        backgroundColor: colors.primary,
        height: '7%',
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: "center",
    },
    leftIconContainer: {
        height: '100%',
        flex: 1,
        justifyContent: 'center',
        alignItems: 'flex-start',
        paddingLeft: 10,
    },
    rightIconContainer: {
        height: '100%',
        flex: 1,
        justifyContent: 'center',
        alignItems: 'flex-end',
        paddingRight: 10,
        paddingTop: 5,
    },
    headerTitleContainer: {
        height: '100%',
        flex: 3,
        alignItems: 'center',
        justifyContent: 'center',
    },
    title: {
        color: colors.white,
        fontSize: 20,
        fontWeight: '400',
        letterSpacing: 2,
    },
});

export default Header