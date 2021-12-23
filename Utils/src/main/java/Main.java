import audio.AudioUtils;
import file.FileUtils;
import hash.HashUtils;

import java.io.File;
import java.io.IOException;
import java.nio.file.CopyOption;
import java.nio.file.Files;
import java.nio.file.StandardCopyOption;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Vector;

public class Main {
    public static void main(String[] args) throws IOException {
        ArrayList<File> allMp4 = FileUtils.getAllMp4("D:\\EatTheBlocks", ".mp4");
        ArrayList<File> allSrt = FileUtils.getAllMp4("D:\\EatTheBlocks", ".srt");
        ArrayList<String> equals = new ArrayList<>();
        for (int i = 0; i < allMp4.size(); i++) {
            for (int j = 0; j < allSrt.size(); j++) {
                if (allMp4.get(i).getAbsolutePath().replaceAll(".mp4", "").equals(allSrt.get(j).getAbsolutePath().replaceAll(".srt", ""))) {
                    equals.add(allMp4.get(i).getAbsolutePath());
                }
            }
        }
        for (int i = 0; i < allMp4.size(); i++) {
            if (!equals.contains(allMp4.get(i).getAbsolutePath())){
                System.out.println(allMp4.get(i));
            }
        }

    }
}
