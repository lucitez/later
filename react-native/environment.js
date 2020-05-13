import Constants from "expo-constants";
import { CLIENT_ID } from 'react-native-dotenv'

const ENV = {
    dev: {
        apiUrl: 'http://192.168.254.64:8080',
        clientId: 'bc42c734-992a-463a-9602-36218ce03152',
    },
    staging: {
        apiUrl: 'https://later-api.appspot.com',
        clientId: 'bc42c734-992a-463a-9602-36218ce03152',
    },
    prod: {
        apiUrl: 'https://later-api.appspot.com',
        clientId: 'bc42c734-992a-463a-9602-36218ce03152',
    }
};

const getEnvVars = (env = Constants.manifest.releaseChannel) => {
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