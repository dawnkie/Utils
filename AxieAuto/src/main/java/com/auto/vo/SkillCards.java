package com.auto.vo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.ArrayList;


@Data
@NoArgsConstructor
@AllArgsConstructor
public class SkillCards {
    private int mostCardsHero = -1;
    private ArrayList<AxieElement> coords1 = null;
    private ArrayList<AxieElement> coords2 = null;
    private ArrayList<AxieElement> coords3 = null;

    @Override
    public String toString() {
        return mostCardsHero + ": 1-" + coords1.toString() + " 2-" + coords2.toString() + " 3-" + coords3.toString();
    }
}
