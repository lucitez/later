import React from 'react';
import { Linking, Alert, TouchableHighlight } from 'react-native';

function Link({ children, url, active }) {
    return (
        <TouchableHighlight activeOpacity={0.75} underlayColor="#FFFFFF" onPress={async () => {
            if (active) {
                // Checking if the link is supported for links with custom URL scheme.
                const supported = await Linking.canOpenURL(url);

                if (supported) {
                    // Opening the link with some app, if the URL scheme is "http" the web link should be opened
                    // by some browser in the mobile
                    await Linking.openURL(url);
                } else {
                    Alert.alert(`Don't know how to open this URL: ${url}`);
                }
            }
        }}>
            {children}
        </TouchableHighlight>
    )
}

export default Link