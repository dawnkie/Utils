package com.auto;

import net.sourceforge.tess4j.*;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import javax.imageio.ImageIO;
import java.awt.*;
import java.awt.image.BufferedImage;
import java.io.File;

@SpringBootTest
class AxieAutoApplicationTests {
    @Autowired
    Tesseract tesseract;

    @Test
    void contextLoads() throws Exception {
        ////ImageAnalyzeUtils.specialRGBsBinarization(ImageIO.read(new File("src/main/resources/screenshot/quest/adventure_x991_y394_w18_h18/quest00.png")), 250);
        BufferedImage bufferedImage = ImageIO.read(new File("src/main/resources/screenshot/binarization/1634389917337.png"));
        //OCRResult documentsWithResults = tesseract.createDocumentsWithResults(bufferedImage, "", "src/main/resources/screenshot/binarization/", Collections.singletonList(ITesseract.RenderedFormat.ALTO), ITessAPI.TessPageIteratorLevel.RIL_SYMBOL);
        ////List<Word> list = documentsWithResults.getWords();
        ////System.out.println(list);
        //System.out.println(tesseract.getWords(bufferedImage,2));

        //File file = new File("src/main/resources/screenshot/binarization/1634389917337.png");
        //String outputBase = "src/main/resources/screenshot/binarization/";
        //List<ITesseract.RenderedFormat> formats = new ArrayList<>(Arrays.asList(ITesseract.RenderedFormat.HOCR, ITesseract.RenderedFormat.TEXT));
        //tesseract.createDocuments(new String[]{file.getPath()}, new String[]{outputBase}, formats);
        Rectangle rect = new Rectangle(490, 479, 18, 18);
        String result = tesseract.doOCR(bufferedImage, rect);
        System.out.println(result);
    }
}
