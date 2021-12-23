package com.auto.vo;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

import java.util.HashMap;

@Data
@Component
@ConfigurationProperties("axie")
public class AxieCoords {
    private HashMap<String,AxieElement> coords;
}
