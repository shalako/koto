/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 * @flow strict-local
 */

import React, {useEffect, useRef} from 'react';
import {StyleSheet, SafeAreaView, Platform, Linking} from 'react-native';
import WebView from 'react-native-webview';
import messaging from '@react-native-firebase/messaging';
import DeviceInfo from 'react-native-device-info';
import CookieManager from '@react-native-community/cookies';
import axios from 'axios';
import SplashScreen from 'react-native-splash-screen';

const url = 'https://koto.at';
const inAppUrls = [
  url,
  'https://www.youtube.com',
  'https://youtube.com',
  'https://m.youtube.com'
];

const App = () => {
  let token = null;
  const webview = useRef(null);

  const requestUserPermission = async () => {
    const authStatus = await messaging().requestPermission();
    const enabled =
      authStatus === messaging.AuthorizationStatus.AUTHORIZED ||
      authStatus === messaging.AuthorizationStatus.PROVISIONAL;

    if (enabled) {
      console.log('Authorization status:', authStatus);
      token = await messaging().getToken();
      console.log(token);
    }
  };

  const onMessage = (event) => {
    const {data} = event.nativeEvent;
    console.log('[REACTNATIVE] NEW Message received from Web:', data);
  };

  useEffect(() => {
    requestUserPermission();
    SplashScreen.hide();
  });

  const onNavigationStateChange = async (navState) => {
    if (navState.url.includes('messages')) {
      const cookies =
        Platform.OS === 'ios'
          ? await CookieManager.getAll(true)
          : await CookieManager.get('https://central.koto.at');
      console.log('CookieManager.get =>', cookies);
      const cookie = Object.values(cookies)
        .map((c) => `${c.name}=${c.value};`)
        .join(' ');
      console.log(cookie);
      try {
        const data = {
          token,
          device_id: DeviceInfo.getDeviceId(),
          os: Platform.OS,
        };
        const res = await axios.request({
          url: 'https://central.koto.at/rpc.UserService/RegisterFCMToken',
          method: 'post',
          headers: {
            Cookie: cookie,
            'Content-Type': 'application/json',
          },
          data,
          withCredentials: false,
        });
        console.log(res);
      } catch (e) {
        console.log(e);
      }
    }
  };

  const urlIsInApp = (str) => {
    for (const item of inAppUrls) {
      if (str.startsWith(item)) {
        return true;
      }
    }
    return false;
  }

  const onShouldLoad = (event, t) => {
  console.log(event,t);
    if (!urlIsInApp(event.url)) {
      Linking.openURL(event.url);
      return false;
    }
    return true;
  };

  return (
    <SafeAreaView style={styles.container}>
      <WebView
        ref={webview}
        source={{uri: url}}
        onMessage={onMessage}
        onNavigationStateChange={onNavigationStateChange}
        onShouldStartLoadWithRequest={onShouldLoad}
      />
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#3f51b5',
  },
});

export default App;
