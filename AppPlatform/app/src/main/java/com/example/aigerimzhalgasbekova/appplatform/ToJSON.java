package com.example.aigerimzhalgasbekova.appplatform;

import android.content.Context;
import android.graphics.Path;
import android.os.Environment;
import android.util.Log;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;

/**
 * Created by aigerimzhalgasbekova on 14/02/2017.
 */

public class ToJSON {

    File path = Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_MOVIES);
    File file;
    String filename = "Res.json";

    public void mCreateAndSaveFile(String mJsonResponse) {
        FileOutputStream outputStream;
        try {
            file = new File(path, "/"+filename);
            outputStream = new FileOutputStream(file);
            outputStream.write(mJsonResponse.getBytes());
            outputStream.close();
            Log.d("TAG", "save");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public String openJSFile(String name){
        //reading jsonString from storage and transform it into jsonObject
        file = new File(path, "/"+name);
        int length = (int) file.length();
        byte[] bytes = new byte[length];
        try {
            FileInputStream in = new FileInputStream(file);
            in.read(bytes);
            in.close();
        } catch (IOException e){

        }
        String contents = new String(bytes);
        return contents;
    }
}
