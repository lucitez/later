import React from 'react';
import { StyleSheet, Text, View, TouchableOpacity } from 'react-native';
import { colors } from '../../assets/colors';

function TabBar({ state, navigation }) {
    return (
        <View style={styles.tabs}>
            {state.routes.map((route, index) => {
                const label = route.name;
                let isFocused = state.index == index

                const onPress = () => {
                    const event = navigation.emit({
                        type: 'tabPress',
                        target: route.key,
                        canPreventDefault: true,
                    });

                    if (!isFocused && !event.defaultPrevented) {
                        navigation.navigate(route.name);
                    }
                };


                return (
                    <TouchableOpacity
                        key={index}
                        style={styles.tabContainer}
                        onPress={onPress}
                    >
                        <View style={[styles.tabNameContainer, isFocused ? styles.activeTab : null]}>
                            <Text style={styles.tabName}>
                                {label}
                            </Text>
                        </View>

                    </TouchableOpacity>
                );
            })}
        </View>
    );
}

const styles = StyleSheet.create({
    tabs: {
        flexDirection: 'row',
        backgroundColor: colors.primary,
    },
    tabContainer: {
        flex: 1,
        alignItems: 'center',
        marginBottom: 2,
    },
    tabNameContainer: {
        alignItems: 'center',
        width: '75%',
        paddingBottom: 10,
    },
    tabName: {
        color: colors.white,
        fontSize: 16,
    },
    activeTab: {
        borderBottomWidth: 2,
        borderColor: colors.white,
    },
});

export default TabBar