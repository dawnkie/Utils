package file;

import hash.HashUtils;

import java.io.File;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Vector;

public class FileUtils {
    private static ArrayList<File> files;

    // 获取目录内部包含的所有后缀为`ends`的文件列表
    public static ArrayList<File> getAllMp4(String dir, String ends) {
        files = new ArrayList<>();
        filterMp4(new File(dir), ends);
        return files;
    }

    // 打印目录内部包含的所有MD5哈希相同的后缀为`ends`的文件列表
    public static void showEqualFiles(String dir, String ends) {
        ArrayList<File> allMp4 = FileUtils.getAllMp4(dir, ends);
        HashMap<String, ArrayList<String>> md5Map = new HashMap<>();
        for (File file : allMp4) {
            String md5 = HashUtils.md5HashCode(file);

            ArrayList<String> paths = md5Map.get(md5);
            if (paths == null) {
                paths = new ArrayList<>();
            }
            paths.add(file.getAbsolutePath());

            md5Map.put(md5, paths);
        }
        md5Map.forEach((k, v) -> {
            if (v.size() > 1) {
                System.out.println("---------------------------------------");
                for (String s : v) {
                    System.out.println(k + " -> " + s);
                }
            }
        });
    }

    private static void filterMp4(File dir, String ends) {
        File[] paths = dir.listFiles(pathname -> pathname.isDirectory() || pathname.getPath().endsWith(ends));
        if (paths != null) {
            for (File path : paths) {
                if (path.isDirectory()) {
                    filterMp4(path, ends);
                } else {
                    files.add(path);
                }
            }
        }
    }

    public static void main(String[] args) {
        getAllMp4("D:\\Pragrams\\Utils\\src\\main\\resources\\1. Welcome.mp4", ".mp4");
    }
}