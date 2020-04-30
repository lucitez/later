import Constants from "expo-constants";
import { CLIENT_ID } from 'react-native-dotenv'

const ENV = {
    dev: {
        apiUrl: 'http://192.168.254.64:8080',
        clientId: CLIENT_ID,
    },
    staging: {
        apiUrl: 'https://later-api.appspot.com',
        clientId: CLIENT_ID,
    }
};

const getEnvVars = (env = Constants.manifest.releaseChannel) => {
    if (__DEV__) {
        return ENV.dev;
    } else if (env === 'staging') {
        return ENV.staging;
    }
};

export default getEnvVars();