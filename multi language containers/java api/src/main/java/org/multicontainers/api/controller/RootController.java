package org.multicontainers.api.controller;

import java.util.HashMap;
import java.util.Map;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@CrossOrigin(origins = "*")
@RequestMapping("/")
public class RootController {

    @GetMapping({ "", "/" })
    public ResponseEntity<Map<String, Object>> check() {
        Map<String, Object> response = new HashMap<String, Object>() {
            {
                put("message", "Welcome to the Java Spring Boot API!");
            }
        };
        return ResponseEntity.ok().body(response);
    }

}
