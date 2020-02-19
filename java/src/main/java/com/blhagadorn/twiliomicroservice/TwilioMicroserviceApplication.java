package com.blhagadorn.twiliomicroservice;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.*;
import org.springframework.boot.autoconfigure.*;
import org.springframework.web.bind.annotation.*;
import org.springframework.beans.factory.annotation.Value;

import java.io.*;


@RestController
@EnableAutoConfiguration
@SpringBootApplication
public class TwilioMicroserviceApplication {

  @Value("${ACCOUNT_SID}")
    private String envVar;

  @RequestMapping("/healthz")
  String home() {
      System.out.println(envVar);
      return "Hello World!";
  }

  void sendMessage() {
    System.out.println("I should send a message here");
  }

	public static void main(String[] args) {
		SpringApplication.run(TwilioMicroserviceApplication.class, args);
	}

}
