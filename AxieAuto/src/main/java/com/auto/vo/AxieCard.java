package com.auto.vo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class AxieCard {
    private int mana;
    private int hero;
    private AxieElement axieElement;
}
