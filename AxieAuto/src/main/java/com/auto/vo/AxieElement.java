package com.auto.vo;


import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class AxieElement {
    private int x;
    private int y;

    // x 坐标随机跳动
    public Integer getRandomX(int num) {
        return x + (int) (Math.random() * num);
    }

    public Integer getRandomX() {
        return x + (int) (Math.random() * 3);
    }

    // y 坐标随机跳动
    public Integer getRandomY(int num) {
        // y 坐标随机跳动
        return y + (int) (Math.random() * num);
    }

    public Integer getRandomY() {
        return y + (int) (Math.random() * 3);
    }

    @Override
    public String toString() {
        return "(" + x + ", " + y + ")";
    }

}
