import Constants from "expo-constants";
import { CLIENT_ID } from 'react-native-dotenv'

const ENV = {
    dev: {
        apiUrl: 'http://192.168.254.64:8080',
        // apiUrl: 'https://later-api.appspot.com',
        clientId: CLIENT_ID,
    },
    staging: {
        apiUrl: 'https://later-api.appspot.com',
        clientId: CLIENT_ID,
    },
    prod: {
        apiUrl: 'https://later-api.appspot.com',
        clientId: CLIENT_ID,
    }
};

const getEnvVars = (env = Constants.manifest.releaseChannel) => {
    console.log(process.env)
    if (__DEV__) {
        return ENV.dev;
    } else if (env == 'staging') {
        return ENV.staging
    } else if (env == 'prod') {
        return ENV.prod
    } else {
        return ENV.staging
    }
};

export default getEnvVars();