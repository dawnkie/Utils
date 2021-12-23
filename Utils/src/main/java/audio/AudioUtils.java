package audio;

import org.bytedeco.javacv.FFmpegFrameGrabber;
import org.bytedeco.javacv.FFmpegFrameRecorder;
import org.bytedeco.javacv.Frame;

import java.io.File;

public class AudioUtils {

	// 音频分离
    public static void extractAudio(File file) {
        FFmpegFrameGrabber grabber = new FFmpegFrameGrabber(file.getAbsolutePath());

        try {
            grabber.start();

            String target = file.getAbsolutePath().replaceAll(".mp4", "") + ".mp3";
            System.out.print(file.getAbsolutePath() + "-->>" + target + "...");
            FFmpegFrameRecorder recorder = new FFmpegFrameRecorder(target, grabber.getAudioChannels());
            recorder.setFormat("mp3");
            recorder.setSampleRate(grabber.getSampleRate());
            recorder.setTimestamp(grabber.getTimestamp());
            recorder.setAudioQuality(0);

            recorder.start();

            while (true) {
                Frame frame = grabber.grab();
                if (frame == null) {
                    System.out.println(" Done!");
                    break;
                }
                if (frame.samples != null) {
                    recorder.recordSamples(frame.sampleRate, frame.audioChannels, frame.samples);
                }
            }

            recorder.stop();
            recorder.release();
            grabber.stop();

        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void main(String[] args) {
        extractAudio(new File("D:\\Pragrams\\EVT\\src\\main\\resources\\audio\\1. Welcome.mp4"));
    }
}