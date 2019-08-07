package com.sbilyi.microservices.authorization.controller

import com.sbilyi.microservices.authorization.LoginController
import org.junit.jupiter.api.extension.ExtendWith
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest
import org.springframework.test.context.junit.jupiter.SpringExtension
import org.springframework.test.web.servlet.MockMvc

@ExtendWith(SpringExtension::class)
@WebMvcTest(LoginController::class)
class LoginTest {

    @Autowired
    private val mvc: MockMvc? = null


}
