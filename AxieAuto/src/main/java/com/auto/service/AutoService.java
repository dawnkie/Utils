package com.auto.service;

import com.auto.driver.UiAutoDriver;
import com.auto.utils.CoordsDictionaries;
import com.auto.utils.ImageAnalyzeUtils;
import com.auto.vo.AxieCoords;
import com.auto.vo.AxieElement;
import com.auto.vo.SkillCards;
import io.appium.java_client.TouchAction;
import io.appium.java_client.android.AndroidDriver;
import io.appium.java_client.android.nativekey.AndroidKey;
import io.appium.java_client.android.nativekey.KeyEvent;
import io.appium.java_client.touch.WaitOptions;
import io.appium.java_client.touch.offset.PointOption;
import org.apache.commons.io.FileUtils;
import org.openqa.selenium.OutputType;
import org.openqa.selenium.WebElement;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Service;

import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;
import java.time.Duration;
import java.util.ArrayList;
import java.util.HashMap;

@Service
public class AutoService implements CommandLineRunner {
    private static final int windowX = 1280;
    private static final int windowY = 720;
    private static final int ADVENTURE = 1;
    private static final int ARENA = 2;
    // 冒险次数
    private static int adNumber = -1;
    // 匹配次数
    private static int arNumber = -1;
    // 当前关卡
    private static int level = 0;
    // 坐标集
    private HashMap<String, AxieElement> coords = null;

    @Autowired
    private AxieCoords axieCoords;
    @Autowired
    private UiAutoDriver uiAutoDriver;

    private AndroidDriver<WebElement> androidDriver;

    public void doAuto() throws InterruptedException, IOException {
        // initial
        coords = axieCoords.getCoords();
        androidDriver = uiAutoDriver.getAndroidDriver();


        System.out.println("程序开始执行...");
        Thread.sleep(8000);
        while (!ImageAnalyzeUtils.isSame100Pixels(ImageIO.read(androidDriver.getScreenshotAs(OutputType.FILE)), coords.get(CoordsDictionaries.isHOME), 0xff00b1cb)) {
            System.out.println("等待进入首页...");
            Thread.sleep(1000);
        }
        // 后续开发 容错:等待10秒,检查是否在首页 NO ? 投降, 重启 : next [androidDriver.closeApp(), androidDriver.launchApp()]

        // 1- adventure < 10 ? adventure till 10 : next
        tap(coords.get(CoordsDictionaries.QUEST), 20, 200);
        androidDriver.pressKey(new KeyEvent().withKey(AndroidKey.BACK));
        tap(coords.get(CoordsDictionaries.QUEST), 20, 2000);
        analysisAdNumber();
        tap(coords.get(CoordsDictionaries.QUEST_close), 0, 200);
        tap(coords.get(CoordsDictionaries.ADVENTURE), 0, 2000);

        // 2- adNumber == 10 ? (hasSLP && hasEnergy ? battle : next) : battle
        // 2.1 识别当前最佳关卡: 点击18关->判断是否提示->二分法得推出关卡
        for (int low = 1, high = 36, mid = 0; low < high; ) {
            mid = (low + high) / 2;
            changeMapLevel(mid);
            BufferedImage screenshoot = ImageIO.read(androidDriver.getScreenshotAs(OutputType.FILE));
            BufferedImage tips = ImageIO.read(new File("src/main/resources/screenshot/others/ADVENTURE_tips.png"));
            AxieElement axieElement = coords.get(CoordsDictionaries.ADVENTURE_tips);
            boolean isTips = ImageAnalyzeUtils.isSame100Pixels(screenshoot, tips, axieElement);
            if (isTips) {
                high = mid - 1;
                tap(coords.get(CoordsDictionaries.ADVENTURE_tips_OK), 20, 300);
            } else {
                low = mid + 1;
            }

            if (low == high) {
                level = mid;
                break;
            }
        }
        // 2.2 循环打关(输了降级,赢了继续),完成后返回首页
        while (true) {
            if (adNumber == 10) {
                BufferedImage screenshoot = ImageIO.read(androidDriver.getScreenshotAs(OutputType.FILE));
                boolean hasEnergy = ImageAnalyzeUtils.isSameImages(screenshoot, screenshoot, 187, 23, 214, 23, 15, 9);
                boolean hasSLP = ImageAnalyzeUtils.isSameImages(screenshoot, screenshoot, 79, 354, 106, 354, 15, 9);
                if (hasSLP && hasEnergy) {
                    System.out.println("[if] 开始冒险, 当前场次: " + adNumber + ", 当前关卡为: " + level);
                    tap(coords.get(CoordsDictionaries.ADVENTURE_START), 0, 200);
                    battle(AutoService.ADVENTURE);
                } else {
                    androidDriver.pressKey(new KeyEvent().withKey(AndroidKey.BACK));
                    break;
                }
            } else {
                System.out.println("[else] 开始冒险, 当前场次: " + adNumber + ", 当前关卡为: " + level);
                tap(coords.get(CoordsDictionaries.ADVENTURE_START), 0, 200);
                battle(AutoService.ADVENTURE);
            }
        }
        // 3- arNumber == 5 ? (hasEnergy ? battle : next) : battle
        // 3.1 识别匹配场次
        tap(coords.get(CoordsDictionaries.QUEST), 100, 200);
        androidDriver.pressKey(new KeyEvent().withKey(AndroidKey.BACK));
        tap(coords.get(CoordsDictionaries.QUEST), 200, 1000);
        analysisArNumber();
        // 3.2 开始匹配
        while (true) {
            if (arNumber == 5) {
                tap(coords.get(CoordsDictionaries.ARENA), 0, 2000);
                BufferedImage screenshoot = ImageIO.read(androidDriver.getScreenshotAs(OutputType.FILE));
                boolean hasEnergy = ImageAnalyzeUtils.isSameImages(screenshoot, screenshoot, 187, 23, 214, 23, 15, 9);
                androidDriver.pressKey(new KeyEvent().withKey(AndroidKey.BACK));
                if (hasEnergy) {
                    System.out.println("匹配场次已经打够,开始消耗剩余能量...");
                    tap(coords.get(CoordsDictionaries.ARENA), 20, 6000);
                    battle(AutoService.ARENA);
                } else {
                    break;
                }
            } else {
                tap(coords.get(CoordsDictionaries.ARENA), 20, 6000);
                System.out.println("开始匹配, 当前场次: " + arNumber);
                battle(AutoService.ARENA);
            }
        }
        // 4- 领取今日任务奖励
        tap(coords.get(CoordsDictionaries.QUEST), 100, 200);
        androidDriver.pressKey(new KeyEvent().withKey(AndroidKey.BACK));
        tap(coords.get(CoordsDictionaries.QUEST), 200, 1000);
        tap(coords.get(CoordsDictionaries.QUEST_CheckIn), 0, 200);
        androidDriver.pressKey(new KeyEvent().withKey(AndroidKey.BACK));
        tap(coords.get(CoordsDictionaries.QUEST), 200, 1000);
        tap(coords.get(CoordsDictionaries.QUEST_Claim), 0, 2000);
        androidDriver.pressKey(new KeyEvent().withKey(AndroidKey.BACK));
        System.out.println("任务完成,正在退出...");
        androidDriver.quit();
    }

    // type: 1 ADVENTURE | 2 ARENA
    private int battle(int type) {
        int info = 0;
        try {
            // 是否处于战场
            AxieElement axieElement = coords.get(CoordsDictionaries.isBATTLE);
            BufferedImage screenshoot = ImageIO.read(androidDriver.getScreenshotAs(OutputType.FILE));
            boolean isBattle = ImageAnalyzeUtils.isSame100Pixels(screenshoot, axieElement, 0xffc07f5b);
            while (isBattle) {
                System.out.println("类型(1 ADVENTURE | 2 ARENA)="+type + ", 是否处于战场: " + isBattle);
                // 是否为出牌时间
                axieElement = coords.get(CoordsDictionaries.isPlayingTime);
                screenshoot = ImageIO.read(androidDriver.getScreenshotAs(OutputType.FILE));
                BufferedImage battleEndTurn = ImageIO.read(new File("src/main/resources/screenshot/battle/BATTLE_END_TURN.png"));
                boolean isPlayingTime = ImageAnalyzeUtils.isSame100Pixels(screenshoot, battleEndTurn, axieElement, 3, 5);
                assert isPlayingTime : "isPlayingTime...";
                if (isPlayingTime) {
                    File screenshotAs = androidDriver.getScreenshotAs(OutputType.FILE);
                    FileUtils.copyFile(screenshotAs, new File("src/main/resources/screenshot/debug/" + adNumber + "_" + System.currentTimeMillis() + ".png"));
                    //SkillCards skillCards = ImageAnalyzeUtils.getSkillCards(screenshotAs);
                    SkillCards skillCards = null;
                    attack(skillCards);
                    axieElement = coords.get(CoordsDictionaries.BATTLE_END_TURN);
                    tap(axieElement, 100, 5500);
                } else {
                    // 是否胜利
                    axieElement = coords.get(CoordsDictionaries.VICTORY);
                    BufferedImage victory = ImageIO.read(new File("src/main/resources/screenshot/battle/VICTORY.png"));
                    boolean isVictory = ImageAnalyzeUtils.isSame100Pixels(screenshoot, victory, axieElement);
                    assert isVictory : "isVictory...";
                    if (isVictory) {
                        System.out.println("战斗胜利! 正在返回...");
                        info = 1;
                        break;
                    }
                    // 是否战败
                    axieElement = coords.get(CoordsDictionaries.DEFEATED);
                    BufferedImage defeated = ImageIO.read(new File("src/main/resources/screenshot/battle/DEFEATED.png"));
                    boolean isDefeated = ImageAnalyzeUtils.isSame100Pixels(screenshoot, defeated, axieElement);
                    assert isDefeated : "isDefeated...";
                    if (isDefeated) {
                        System.out.println("战斗失败! 正在返回...");
                        info = -1;
                        break;
                    }
                }
            }
            System.out.println("当前战斗已结束,准备中...");
            tap(new AxieElement(windowX - 15, windowY - 15), 30, 1000);
            tap(new AxieElement(windowX - 15, windowY - 15), 30, 1000);
            tap(new AxieElement(windowX - 15, windowY - 15), 30, 1000);
            tap(new AxieElement(windowX - 15, windowY - 15), 30, 5000);
            switch (info) {
                case -1:
                    if (type == AutoService.ADVENTURE) {
                        changeMapLevel(--level);
                    }
                    break;
                case 1:
                    if (type == AutoService.ADVENTURE) {
                        adNumber = Math.min(++adNumber, 10);
                    } else if (type == AutoService.ARENA) {
                        arNumber = Math.min(++arNumber, 5);
                    }
                    break;
                default:
                    System.out.println("当前页面不是战场 !");
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        return info;
    }

    private int analysisAdNumber() {
        try {
            BufferedImage screenshot = ImageIO.read(androidDriver.getScreenshotAs(OutputType.FILE));
            AxieElement axieElement = coords.get(CoordsDictionaries.isQUEST);
            if (ImageAnalyzeUtils.isSame100Pixels(screenshot, axieElement, 0xfffbb100)) {
                for (int i = 0; i < 10; i++) {
                    if (ImageAnalyzeUtils.isSameImages(screenshot, ImageIO.read(new File("src/main/resources/screenshot/quest/adventure_x991_y394_w18_h18/" + i + ".png")), 991, 394, 0, 0, 18, 18)) {
                        adNumber = i;
                    }
                }
                adNumber = adNumber == -1 ? 10 : adNumber;
            } else {
                System.out.println("analysisAdNumber(): adNumber 解析失败, 页面不对 !!!");
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        System.out.println("Adventure 当前场次: " + adNumber);
        return adNumber;
    }

    private int analysisArNumber() {
        try {
            BufferedImage screenshot = ImageIO.read(androidDriver.getScreenshotAs(OutputType.FILE));
            AxieElement axieElement = coords.get(CoordsDictionaries.isQUEST);
            if (ImageAnalyzeUtils.isSame100Pixels(screenshot, axieElement, 0xfffbb100)) {
                BufferedImage qusetScreenshot = ImageIO.read(androidDriver.getScreenshotAs(OutputType.FILE));
                for (int i = 0; i < 5; i++) {
                    if (ImageAnalyzeUtils.isSameImages(qusetScreenshot, ImageIO.read(new File("src/main/resources/screenshot/quest/arena_x490_y479_w18_h18/" + i + ".png")), 490, 479, 0, 0, 18, 18)) {
                        arNumber = i;
                    }
                }
                arNumber = arNumber == -1 ? 5 : arNumber;
            } else {
                System.out.println("analysisAdNumber(): arNumber 解析失败, 页面不对 !!!");
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        System.out.println("Arena 当前场次: " + arNumber);
        return arNumber;
    }

    private int getBattleEnergy() {
        int battleEnergy = 0;
        try {
            for (int i = 0; i <= 10; i++) {
                if (ImageAnalyzeUtils.isSameImages(ImageIO.read(androidDriver.getScreenshotAs(OutputType.FILE)), ImageIO.read(new File("src/main/resources/screenshot/battle/skill_energy/" + i + ".png")), 32, 569, 32, 569, 22, 13)) {
                    battleEnergy = i;
                    break;
                }
            }
        } catch (Exception e) {
            System.out.println("com.mtl.axieauto01.service.AutoService.getBattleEnergy()");
            e.printStackTrace();
        }
        System.out.println("com.mtl.axieauto01.service.AutoService.getBattleEnergy(): " + battleEnergy);
        return battleEnergy;

    }

    public void attack(SkillCards skillCards) {
        // 战斗能量
        int battleEnergy = getBattleEnergy();
        if (battleEnergy <= 2 || ((skillCards.getCoords1().size() + skillCards.getCoords2().size() + skillCards.getCoords3().size()) == 0)) {
            System.out.println("当前卡牌能量为: " + battleEnergy);
            return;
        }
        int mostCardsHero = skillCards.getMostCardsHero();
        // 牌最多的英雄先出
        switch (mostCardsHero) {
            case 1:
                battleEnergy = attack(skillCards.getCoords1(), battleEnergy);
                battleEnergy = attack(skillCards.getCoords2(), battleEnergy);
                battleEnergy = attack(skillCards.getCoords3(), battleEnergy);
                break;
            case 2:
                battleEnergy = attack(skillCards.getCoords2(), battleEnergy);
                battleEnergy = attack(skillCards.getCoords1(), battleEnergy);
                battleEnergy = attack(skillCards.getCoords3(), battleEnergy);
                break;
            case 3:
                battleEnergy = attack(skillCards.getCoords3(), battleEnergy);
                battleEnergy = attack(skillCards.getCoords1(), battleEnergy);
                battleEnergy = attack(skillCards.getCoords2(), battleEnergy);
                break;
            default:
                System.out.println("attack(): 未识别到英雄!");
        }
    }

    private int attack(ArrayList<AxieElement> axieElements, int battleEnergy) {
        if (axieElements == null) {
            return battleEnergy;
        }
        AxieElement axieElement1, axieElement2;
        for (int i = 0, oneTurn = 0; i < axieElements.size(); i++) {
            if (battleEnergy == 0 || oneTurn > 3) {
                return battleEnergy;
            }
            System.out.println("英雄出牌: " + (oneTurn + 1));
            axieElement1 = axieElements.get(i);
            axieElement2 = new AxieElement();
            axieElement2.setX(450 + (int) (Math.random() * 400));
            axieElement2.setY(140 + (int) (Math.random() * 180));
            swipe(axieElement1, axieElement2, 100, 100, 500);
            oneTurn++;
            battleEnergy--;
        }
        return battleEnergy;
    }

    public void swipe(AxieElement axieElement1, AxieElement axieElement2, int intervalBefore, int intervalAfter, int waitTime) {
        try {
            Thread.sleep((long) (intervalBefore + 20 * Math.random()));
            new TouchAction<>(androidDriver)
                    .press(PointOption.point(axieElement1.getRandomX(), axieElement1.getRandomY()))
                    .waitAction(WaitOptions.waitOptions(Duration.ofMillis(waitTime + (long) (20 * Math.random()))))
                    .moveTo(PointOption.point(axieElement2.getRandomX(), axieElement2.getRandomY()))
                    .release()
                    .perform();
            Thread.sleep((long) (intervalAfter + 20 * Math.random()));
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    public void tap(AxieElement axieElement, int intervalBefore, int intervalAfter) {
        System.out.println(axieElement);
        try {
            Thread.sleep((long) (intervalBefore + 20 * Math.random()));
            new TouchAction<>(androidDriver).tap(PointOption.point(axieElement.getRandomX(), axieElement.getRandomY())).perform();
            Thread.sleep((long) (intervalAfter + 20 * Math.random()));
        } catch (Exception e) {
            e.printStackTrace();
        }

    }

    public void changeMapLevel(int level) {
        try {
            AxieElement axieElement_L = coords.get(CoordsDictionaries.ADVENTURE_swipe_L);
            AxieElement axieElement_R = coords.get(CoordsDictionaries.ADVENTURE_swipe_R);
            swipe(axieElement_L, axieElement_R, 30, 30, 200);
            if (level <= 15) {
                tap(coords.get("ADVENTURE_" + level), 30, 500);
            } else if (level > 28) {
                swipe(axieElement_L, axieElement_R, 30, 30, 200);
                for (int i = 0; i < 5; i++) {
                    tap(coords.get(CoordsDictionaries.ADVENTURE_right), 30, 200);
                }
                tap(coords.get("ADVENTURE_" + level), 30, 500);
            } else {
                swipe(axieElement_L, axieElement_R, 30, 30, 200);
                for (int i = 0; i < 3; i++) {
                    tap(coords.get(CoordsDictionaries.ADVENTURE_right), 30, 200);
                }
                tap(coords.get("ADVENTURE_" + level), 30, 500);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }


    @Override
    public void run(String... args) throws Exception {
        System.out.println("程序启动: ");
        //doAuto();
    }
}
