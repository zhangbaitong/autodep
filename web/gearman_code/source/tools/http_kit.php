<?php

/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
 class http
 {
    public function get_web_result($str_url,$time_out=10)
    {
        //提交请求
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $str_url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_TIMEOUT, $time_out);        
        curl_setopt($ch, CURLOPT_HEADER, 0);
        $str_content = curl_exec($ch);
        curl_close($ch);
        return $str_content;
    }
    
    public function post_web_result($str_url,$post_data,$time_out=10)
    {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $str_url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_TIMEOUT, $time_out);        
        curl_setopt($ch, CURLOPT_POST, 1);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $post_data);
        $str_content = curl_exec($ch);
        curl_close ( $ch );        
        return $str_content;
    }
}

?>