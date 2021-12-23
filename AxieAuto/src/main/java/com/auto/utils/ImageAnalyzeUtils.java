package com.auto.utils;

import com.auto.vo.AxieElement;
import com.auto.vo.SkillCards;

import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;

public class ImageAnalyzeUtils {
    // RGB 单通道允许的最大偏差
    public static final int R_G_B_ALLOW = 5;
    // RBG 方差允许的最大偏差
    public static final int RGB_Deviation_ALLOW = 10;
    public static final int GRAY_AVERAGE = 0;
    public static final int GRAY_BIGGER = 1;
    public static final int GRAY_SMALLER = 2;
    public static final int GRAY_POWERFUL = 3;

    // 方差 s^2=[(x1-x)^2 +...(xn-x)^2]/n 或者s^2=[(x1-x)^2 +...(xn-x)^2]/(n-1)
    public static double variance(int[] arr) {
        long sum = 0;
        for (int j : arr) {
            sum += j;
        }
        int average = (int) (sum / arr.length);
        double dev = 0;
        for (int j : arr) {
            dev += Math.pow(j - average, 2);
        }
        return dev / arr.length;
    }

    // 标准差 σ=sqrt(s^2)
    public static double standardDeviation(int[] arr) {
        double deviation = ImageAnalyzeUtils.variance(arr);
        return Math.sqrt(deviation / arr.length);
    }

    // RGB -> R_G_B
    public static int[] getR_G_B(int rgb) {
        int[] r_g_b = new int[3];
        r_g_b[0] = (rgb & 0x00ff0000) >> 16;
        r_g_b[1] = (rgb & 0x0000ff00) >> 8;
        r_g_b[2] = (rgb & 0x000000ff);
        return r_g_b;
    }

    // R_G_B 均值
    public static int getAverageR_G_B(int rgb) {
        int[] r_g_b = getR_G_B(rgb);
        return (r_g_b[0] + r_g_b[1] + r_g_b[2]) / 3;
    }

    // 基于RGB通道的相似度分析
    public static double getRGBSimilarity(File analyzePicFile, File templatePicFile, int startX, int startY, int w, int h) {
        double similarity = 0;
        try {
            BufferedImage analyzeImg = ImageIO.read(analyzePicFile);
            BufferedImage templateImg = ImageIO.read(templatePicFile);
            // 参数：起点X，起点Y，总像素宽度，总像素高度，承接返回值的数组，从数组的什么位置开始存储，扫描分析多少宽度的像素
            int[] rgbArray1 = analyzeImg.getRGB(startX, startY, w, h, null, 0, w);
            int[] rgbArray2 = templateImg.getRGB(startX, startY, w, h, null, 0, w);
            similarity = getRGBSimilarity(rgbArray1, rgbArray2);
        } catch (IOException e) {
            e.printStackTrace();
        }
        return similarity;
    }

    // 基于RGB通道的相似度分析
    public static double getRGBSimilarity(int[] rgbArray1, int[] rgbArray2) {
        int count = 0;
        double similarity;
        for (int i = 0; i < rgbArray1.length; i++) {
            if (isRGBPixelSimilar(rgbArray1[i], rgbArray2[i])) {
                count++;
            }
        }
        similarity = (double) count / rgbArray1.length;
        return similarity;
    }

    // 基于RGB通道的 1 个像素的相似度分析
    private static boolean isRGBPixelSimilar(int RGBPixel1, int RGBPixel2) {
        int[] r_g_b1;
        int[] r_g_b2;
        int diff_R;
        int diff_G;
        int diff_B;
        boolean simFlag = true;
        r_g_b1 = getR_G_B(RGBPixel1);
        r_g_b2 = getR_G_B(RGBPixel2);
        diff_R = r_g_b1[0] - r_g_b2[0];
        diff_G = r_g_b1[1] - r_g_b2[1];
        diff_B = r_g_b1[2] - r_g_b2[2];

        final int H = 45;
        if (diff_R + diff_G + diff_B > 15) {
            simFlag = false;
        }
        return simFlag;
    }


    // 基于RGB方差的相似度分析
    public static double getDeviationSimilarity(File analyzePicFile, File templatePicFile, int startX, int startY, int w, int h) {
        double similarity = 0;
        try {
            BufferedImage analyzeImg = ImageIO.read(analyzePicFile);
            BufferedImage templateImg = ImageIO.read(templatePicFile);
            // 参数：起点X，起点Y，总像素宽度，总像素高度，承接返回值的数组，从数组的什么位置开始存储，扫描分析多少宽度的像素
            int[] rgbArray1 = analyzeImg.getRGB(startX, startY, w, h, null, 0, w);
            int[] rgbArray2 = templateImg.getRGB(startX, startY, w, h, null, 0, w);
            similarity = getDeviationSimilarity(rgbArray1, rgbArray2);
        } catch (IOException e) {
            e.printStackTrace();
        }
        return similarity;
    }

    // 基于RGB方差的相似度分析
    private static double getDeviationSimilarity(int[] rgbArray1, int[] rgbArray2) {
        double similarity;
        int count = 0;
        for (int i = 0; i < rgbArray1.length; i++) {
            if (standardDeviation(getR_G_B(rgbArray1[i])) - standardDeviation(getR_G_B(rgbArray2[i])) <= ImageAnalyzeUtils.RGB_Deviation_ALLOW) {
                count++;
            }
        }
        similarity = (double) count / rgbArray1.length;
        return similarity;
    }

    // 消除意外：将数组中“最多连续 n 个元素的值不等于corr”的现象看做 意外
    private static void excludeAccident(int[] arrayDiff, int corr, int n) {
        int count = 0;
        for (int i = 0; i < arrayDiff.length - 1; i++) {
            // 如果元素不正确，则消除意外（但如果连续不对的元素足够多，则认为其合理）
            if (arrayDiff[i] != corr) {
                count++;
                if (count <= n && arrayDiff[i + 1] == corr && (i - count < 0 || arrayDiff[i - count] == corr)) {
                    for (int j = 0; j < count; j++) {
                        arrayDiff[i - j] = corr;
                    }
                }
            } else {
                count = 0;
            }
        }
        if (arrayDiff[arrayDiff.length - 2] == corr && arrayDiff[arrayDiff.length - 1] != corr) {
            arrayDiff[arrayDiff.length - 1] = corr;
        }
    }

    public static boolean isSameImages(File analyzeImgFile, File templateImgFile, int startX1, int startY1, int startX2, int startY2, int w, int h) throws IOException {
        BufferedImage analyzeImg = ImageIO.read(analyzeImgFile);
        BufferedImage templateImg = ImageIO.read(templateImgFile);
        return isSameRGBArray(analyzeImg, templateImg, startX1, startY1, startX2, startY2, w, h);
    }

    public static boolean isSameImages(BufferedImage analyzeImg, BufferedImage templateImg, int startX1, int startY1, int startX2, int startY2, int w, int h) {
        return isSameRGBArray(analyzeImg, templateImg, startX1, startY1, startX2, startY2, w, h);
    }

    public static boolean isSame100Pixels(BufferedImage analyzeImg, BufferedImage templateImg, AxieElement axieElement) {
        return isSameRGBArray(analyzeImg, templateImg, axieElement.getX(), axieElement.getY(), axieElement.getX(), axieElement.getY(), 100, 1);
    }

    public static boolean isSame100Pixels(BufferedImage analyzeImg, BufferedImage templateImg, AxieElement axieElement, int w, int h) {
        return isSameRGBArray(analyzeImg, templateImg, axieElement.getX(), axieElement.getY(), axieElement.getX(), axieElement.getY(), w, h);
    }


    public static boolean isSameRGBArray(BufferedImage analyzeImg, BufferedImage templateImg, int startX1, int startY1, int startX2, int startY2, int w, int h) {
        int[] rgbArray1 = analyzeImg.getRGB(startX1, startY1, w, h, null, 0, w);
        int[] rgbArray2 = templateImg.getRGB(startX2, startY2, w, h, null, 0, w);
        for (int i = 0; i < rgbArray1.length; i++) {
            if (rgbArray1[i] - rgbArray2[i] != 0) {
                return false;
            }
        }
        return true;
    }


    // 图片反相
    public static int[] contraryRGB(BufferedImage analyzeImg, int x, int y, int w, int h) {
        // 比较 100个 像素
        int[] rgbArray = analyzeImg.getRGB(x, y, w, h, null, 0, w);
        int[] contraryArray = new int[rgbArray.length];
        for (int i = 0; i < rgbArray.length; i++) {
            contraryArray[i] = 0xffffffff - rgbArray[i] + 0xff000000;
        }
        return contraryArray;
    }

    public static boolean isSame100Pixels(BufferedImage analyzeImg, AxieElement axieElement, int RGB) {
        // 比较 100个 像素
        int[] rgbArray = analyzeImg.getRGB(axieElement.getX(), axieElement.getY(), 100, 1, null, 0, 100);
        for (int i : rgbArray) {
            if (i - RGB != 0) {
                return false;
            }
        }
        return true;
    }

    //// 图片像素要求：1280*720
    //public static SkillCards getSkillCards(File analyzeImgFile) {
    //    SkillCards skillCards = new SkillCards();
    //    ArrayList<Integer> cardsNumber = new ArrayList<>();
    //    ArrayList<AxieElement> coords = new ArrayList<>();
    //    try {
    //        // cards
    //        int[] rgb1 = ImageIO.read(analyzeImgFile).getRGB(60, 571, 1220, 1, null, 0, 1220);
    //        for (int i = 0, energy = 0; i < rgb1.length - 6; i++) {
    //            if (rgb1[i] == 0xffffffff) {
    //                coords.add(new AxieElement(i + 60 + 36, 630, 1));
    //                if (rgb1[i + 5] == 0xffffffff && rgb1[i + 6] == 0xffffffff) {
    //                    coords.get(coords.size() - 1).setEnergy(0);
    //                }
    //                i += 88;
    //            }
    //        }
    //        // cardsNumber
    //        for (int i = 0, count = 1; i < coords.size() - 1; i++) {
    //            if ((coords.get(i + 1).getX() - coords.get(i).getX() > 88) && (coords.get(i + 1).getX() - coords.get(i).getX() < 95)) {
    //                count++;
    //            } else {
    //                cardsNumber.add(count);
    //                count = 1;
    //            }
    //            if (i == coords.size() - 2) {
    //                cardsNumber.add(count);
    //            }
    //        }
    //        System.out.println(cardsNumber);
    //        // AxieElement.setHero()
    //        for (int i = 0, num = 0; i < cardsNumber.size(); i++) {
    //            num += i == 0 ? 0 : cardsNumber.get(i - 1);
    //            for (int j = 0; j < cardsNumber.get(i); j++) {
    //                coords.get(num + j).setHero(i + 1);
    //            }
    //        }
    //        // SkillCards.setMostCardsHero(): hero 1, hero 2, hero 3
    //        int mostCardsHero = 1;
    //        if (cardsNumber.size() != 0) {
    //            int[] temp = new int[cardsNumber.size()];
    //            for (int i = 0; i < cardsNumber.size(); i++) {
    //                temp[i] = cardsNumber.get(i);
    //            }
    //            Arrays.sort(temp);
    //            mostCardsHero = cardsNumber.indexOf(temp[cardsNumber.size() - 1]) + 1;
    //        }
    //        // SkillCards: setCoords1() setCoords2() setCoords3()
    //        ArrayList<AxieElement> subCoords1 = new ArrayList<AxieElement>();
    //        ArrayList<AxieElement> subCoords2 = new ArrayList<AxieElement>();
    //        ArrayList<AxieElement> subCoords3 = new ArrayList<AxieElement>();
    //        for (AxieElement i : coords) {
    //            switch (i.getHero()) {
    //                case 1:
    //                    subCoords1.add(i);
    //                    break;
    //                case 2:
    //                    subCoords2.add(i);
    //                    break;
    //                case 3:
    //                    subCoords3.add(i);
    //                    break;
    //                default:
    //                    System.out.println("识别出3个以上的英雄!");
    //            }
    //        }
    //        skillCards.setCoords1(subCoords1);
    //        skillCards.setCoords2(subCoords2);
    //        skillCards.setCoords3(subCoords3);
    //        skillCards.setMostCardsHero(mostCardsHero);
    //    } catch (IOException e1) {
    //        e1.printStackTrace();
    //    }
    //    System.out.println("卡牌识别结果: " + skillCards);
    //    return skillCards;
    //}

    // 二值化: 特定颜色范围二值化 算法: RGB均值大于threshold的置为黑,否则置为白
    public static void specialRGBsBinarization(BufferedImage image, int threshold) throws IOException {
        int w = image.getWidth();
        int h = image.getHeight();
        int[] imageRGB = image.getRGB(0, 0, w, h, null, 0, w);
        int[] bi = new int[w * h];
        for (int i = 0; i < imageRGB.length; i++) {
            if (getAverageR_G_B(imageRGB[i]) >= threshold) {
                bi[i] = 0xff000000;
            } else {
                bi[i] = 0xffffffff;
            }
        }
        BufferedImage biImage = new BufferedImage(w, h, BufferedImage.TYPE_BYTE_BINARY);
        biImage.setRGB(0, 0, w, h, bi, 0, w);
        ImageIO.write(biImage, "png", new File("src/main/resources/screenshot/binarization/"+System.currentTimeMillis()+".png"));
    }

    // 二值化: 像素二值化 算法: 灰度值与阈值比较, 要么置为白,要么置为黑
    public static void binarization(BufferedImage image, int grayMethod, int threshold) throws IOException {
        int h = image.getHeight();
        int w = image.getWidth();
        int[] rgb = image.getRGB(0, 0, w, h, null, 0, w);
        int[] bi = new int[w * h];
        switch (grayMethod) {
            case ImageAnalyzeUtils.GRAY_AVERAGE:
                for (int i = 0; i < rgb.length; i++) {
                    if (getGray_Bigger(rgb[i]) > threshold) {
                        bi[i] = 0xffffffff;
                    } else {
                        bi[i] = 0xff000000;
                    }
                }
                break;
            case ImageAnalyzeUtils.GRAY_BIGGER:
                for (int i = 0; i < rgb.length; i++) {
                    if (getGray_Average(rgb[i]) > threshold) {
                        bi[i] = 0xffffffff;
                    } else {
                        bi[i] = 0xff000000;
                    }
                }
                break;
            case ImageAnalyzeUtils.GRAY_SMALLER:
                for (int i = 0; i < rgb.length; i++) {
                    if (getGray_Smaller(rgb[i]) > threshold) {
                        bi[i] = 0xffffffff;
                    } else {
                        bi[i] = 0xff000000;
                    }
                }
                break;
            case ImageAnalyzeUtils.GRAY_POWERFUL:
                for (int i = 0; i < rgb.length; i++) {
                    if (getGray_Powerful(rgb[i]) > threshold) {
                        bi[i] = 0xffffffff;
                    } else {
                        bi[i] = 0xff000000;
                    }
                }
                break;
            default:
                System.out.println("grayMethod is wrong!");
                return;
        }
        for (int i = 0; i < rgb.length; i++) {
            if (getGray_Bigger(rgb[i]) >= threshold) {
                bi[i] = 0xffffffff;
            } else {
                bi[i] = 0xff000000;
            }
        }
        BufferedImage biImage = new BufferedImage(w, h, BufferedImage.TYPE_BYTE_BINARY);
        biImage.setRGB(0, 0, w, h, bi, 0, w);
        ImageIO.write(biImage, "png", new File("src/main/resources/screenshot/others/0_binarizationBig.png"));
    }

    // 二值化: 范围二值化 算法: 相对灰度值与阈值比较, 要么置为白,要么置为黑
    public static void relativeBinarization(BufferedImage image, int grayMethod, int threshold) throws IOException {
        int h = image.getHeight();
        int w = image.getWidth();
        int[][] gray = new int[w][h];
        BufferedImage biImage = new BufferedImage(w, h, BufferedImage.TYPE_BYTE_BINARY);
        switch (grayMethod) {
            case ImageAnalyzeUtils.GRAY_AVERAGE:
                for (int i = 0; i < w; i++) {
                    for (int j = 0; j < h; j++) {
                        gray[i][j] = getGray_Average(image.getRGB(i, j));
                    }
                }
                break;
            case ImageAnalyzeUtils.GRAY_BIGGER:
                for (int i = 0; i < w; i++) {
                    for (int j = 0; j < h; j++) {
                        gray[i][j] = getGray_Bigger(image.getRGB(i, j));
                    }
                }
                break;
            case ImageAnalyzeUtils.GRAY_SMALLER:
                for (int i = 0; i < w; i++) {
                    for (int j = 0; j < h; j++) {
                        gray[i][j] = getGray_Smaller(image.getRGB(i, j));
                    }
                }
                break;
            case ImageAnalyzeUtils.GRAY_POWERFUL:
                for (int i = 0; i < w; i++) {
                    for (int j = 0; j < h; j++) {
                        gray[i][j] = getGray_Powerful(image.getRGB(i, j));
                    }
                }
                break;
            default:
                System.out.println("grayMethod is wrong!");
                return;
        }
        for (int i = 0; i < w; i++) {
            for (int j = 0; j < h; j++) {
                if (getRelativeGray(gray, i, j, w, h) > threshold) {
                    biImage.setRGB(i, j, 0xffffffff);
                } else {
                    biImage.setRGB(i, j, 0xff000000);
                }
            }
        }
        ImageIO.write(biImage, "png", new File("src/main/resources/screenshot/others/0_relativeBinarizationBig.png"));
    }

    // [最大值法] Gray = R>G?:max(R,B):max(G,B)
    public static int getGray_Bigger(int rgb) {
        int[] r_g_b = getR_G_B(rgb);
        return r_g_b[0] > r_g_b[1] ? Math.max(r_g_b[0], r_g_b[2]) : Math.max(r_g_b[1], r_g_b[2]);
    }

    // [最小值法] Gray = R<G?:min(R,B):min(G,B)
    public static int getGray_Smaller(int rgb) {
        int[] r_g_b = getR_G_B(rgb);
        return r_g_b[0] < r_g_b[1] ? Math.min(r_g_b[0], r_g_b[2]) : Math.min(r_g_b[1], r_g_b[2]);
    }

    // [平均值法] Gray = (R+G+B)/3
    public static int getGray_Average(int rgb) {
        int[] r_g_b = getR_G_B(rgb);
        return (r_g_b[0] + r_g_b[1] + r_g_b[2]) / 3;
    }

    // [加权法] Gray = (int) (0.3 * R + 0.59 * G + 0.11 * B)
    public static int getGray_Powerful(int rgb) {
        int[] r_g_b = getR_G_B(rgb);
        return (int) (0.3 * r_g_b[0] + 0.59 * r_g_b[1] + 0.11 * r_g_b[2]);
    }

    // [平均值法] RelativeGray = 自己及周围8像素的灰度均值
    public static int getRelativeGray(int[][] gray, int x, int y, int w, int h) {
        int gray9 = gray[x][y]
                + (x == 0 ? 255 : gray[x - 1][y])
                + (x == 0 || y == 0 ? 255 : gray[x - 1][y - 1])
                + (x == 0 || y == h - 1 ? 255 : gray[x - 1][y + 1])
                + (y == 0 ? 255 : gray[x][y - 1])
                + (y == h - 1 ? 255 : gray[x][y + 1])
                + (x == w - 1 ? 255 : gray[x + 1][y])
                + (x == w - 1 || y == 0 ? 255 : gray[x + 1][y - 1])
                + (x == w - 1 || y == h - 1 ? 255 : gray[x + 1][y + 1]);
        return gray9 / 9;
    }


    public static void main(String[] args) throws IOException {
        // 523 531
        // 453 461
        //SkillCards skillCards = getSkillCards(new File("src/main/resources/screenshot/battle/skill_200_603_19_1.png"));
        //for (int i = 0; i < 5; i++) {
        //    String path1 = "src/main/resources/screenshot/quest/arena_x490_y479_w18_h18/quest0";
        //    String path2 = "src/main/resources/screenshot/quest/arena_x490_y479_w18_h18/";
        //    path1 = path1 + i + ".png";
        //    path2 = path2 + i + ".png";
        //    BufferedImage bufferedImage = ImageIO.read(new File(path1));
        //    BufferedImage subimage = bufferedImage.getSubimage(490, 479, 18, 18);
        //    ImageIO.write(subimage, "png", new FileOutputStream(new File(path2)));
        //}

        //relativeBinarization(ImageIO.read(new File("src/main/resources/screenshot/quest/adventure_x991_y394_w18_h18/quest00_contrary.png")), 50);
        //binarization(ImageIO.read(new File("src/main/resources/screenshot/quest/adventure_x991_y394_w18_h18/quest00_contrary.png")), 50);
        //int[] ints = contraryRGB(ImageIO.read(new File("src/main/resources/screenshot/quest/adventure_x991_y394_w18_h18/quest00.png")), 0, 0, 1280, 720);
        //BufferedImage bufferedImage = new BufferedImage(1280, 720, BufferedImage.TYPE_INT_ARGB);
        //bufferedImage.setRGB(0, 0, 1280, 720, ints, 0, 1280);
        //ImageIO.write(bufferedImage, "png", new File("src/main/resources/screenshot/quest/adventure_x991_y394_w18_h18/quest00_contrary.png"));

    }


}
