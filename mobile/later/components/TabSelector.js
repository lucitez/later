import React, { useState } from 'react';
import { StyleSheet, Text, View, TouchableOpacity } from 'react-native';
import { colors } from '../assets/colors';

function Tab(props) {
    return (
        <View style={styles.tab}>
            <TouchableOpacity style={styles.tabNameContainer} onPress={props.onPress}>
                <Text style={styles.tabName} color='#fff'>{props.name} </Text>
            </TouchableOpacity>
            <View style={activeTabStyle(props.active)} />
        </View>
    )
}

function TabSelector(props) {

    const [activeTab, setActiveTab] = useState(props.tabs[0])

    return (
        <View style={styles.tabs}>
            {
                props.tabs.map((tabName) => (
                    <Tab
                        name={tabName}
                        key={tabName}
                        active={tabName == activeTab}
                        onPress={() => {
                            setActiveTab(tabName)
                            props.onTabChange(tabName)
                        }}
                    />
                ))
            }
        </View>
    );
}

const styles = StyleSheet.create({
    tabs: {
        height: '100%',
        width: '100%',
        flexDirection: 'row',
        justifyContent: 'space-around',
    },
    tab: {
        flex: 1,
        flexDirection: 'column',
        justifyContent: 'space-around',
        alignItems: 'center',
    },
    tabNameContainer: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
        width: '100%',
    },
    tabName: {
        color: colors.white,
        fontSize: 20,
    },
    button: {
        borderWidth: 1,
        borderColor: colors.white
    },
    activeTab: {
        borderBottomWidth: 2,
        borderColor: colors.white,
        width: '80%',
        marginBottom: 5,
    },
});

const activeTabStyle = (active) => {
    return {
        borderWidth: 2,
        borderColor: colors.white,
        width: '80%',
        marginBottom: 5,
        opacity: active ? 1 : 0
    }
}

export default TabSelector