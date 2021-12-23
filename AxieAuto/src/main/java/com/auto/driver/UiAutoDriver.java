package com.auto.driver;

import io.appium.java_client.android.AndroidDriver;
import lombok.Data;
import org.openqa.selenium.WebElement;
import org.openqa.selenium.remote.DesiredCapabilities;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

import java.net.MalformedURLException;
import java.net.URL;

@Data
@Component
@ConfigurationProperties(prefix = "appium")
public class UiAutoDriver {
    private String deviceName;
    private String platformName;
    private String appPackage;
    private String appActivity;
    private String noReset;
    private String automationName;
    private String serverIp;
    private String dontStopAppOnReset;

    public AndroidDriver<WebElement> getAndroidDriver() {
        DesiredCapabilities desiredCapabilities = new DesiredCapabilities();
        desiredCapabilities.setCapability("deviceName", deviceName);
        desiredCapabilities.setCapability("platformName", platformName);
        desiredCapabilities.setCapability("appPackage", appPackage);
        desiredCapabilities.setCapability("appActivity", appActivity);
        desiredCapabilities.setCapability("noReset", noReset);
        desiredCapabilities.setCapability("automationName", automationName);
        desiredCapabilities.setCapability("dontStopAppOnReset", dontStopAppOnReset);
        AndroidDriver<WebElement> androidDriver = null;
        try {
            androidDriver = new AndroidDriver<>(new URL(serverIp),desiredCapabilities);
        } catch (MalformedURLException e) {
            e.printStackTrace();
        }
        return androidDriver;
    }

}
