import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { colors } from '../../assets/colors';


function Header(props) {
    return (
        <View style={styles.container}>
            <View style={styles.leftIconContainer}>
                {props.leftIcon ? props.leftIcon : null}
            </View>
            <View style={styles.headerTitleContainer}>
                <Text style={styles.title}>{props.name.toUpperCase()}</Text>
            </View>
            <View style={styles.rightIconContainer}>
                {props.rightIcon ? props.rightIcon : null}
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        backgroundColor: colors.primary,
        height: '11%',
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: "flex-end",
        paddingTop: 40,
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