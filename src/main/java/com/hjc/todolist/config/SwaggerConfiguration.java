package com.hjc.todolist.config;

import io.swagger.v3.oas.models.Components;
import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.info.Contact;
import io.swagger.v3.oas.models.info.Info;
import io.swagger.v3.oas.models.security.SecurityRequirement;
import io.swagger.v3.oas.models.security.SecurityScheme;
import org.springdoc.core.models.GroupedOpenApi;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class SwaggerConfiguration {
    private static final String basePackage = "com.hjc.todolist.controller";

    private static final String headerName = "Authorization";

    @Bean
    public GroupedOpenApi groupedOpenApi() {
        return GroupedOpenApi.builder()
                .group("adminGroup")
                .addOperationCustomizer((operation, handlerMethod) -> {
                    operation.addSecurityItem(new SecurityRequirement().addList(headerName));
                    return operation;
                })
                .packagesToScan(basePackage)
                .build();
    }

    @Bean
    public OpenAPI customOpenAPI() {
        Components components = new Components();
        // 登录认证token
        components.addSecuritySchemes(headerName,
                new SecurityScheme()
                        .type(SecurityScheme.Type.APIKEY)
                        .scheme("basic")
                        .name(headerName)
                        .in(SecurityScheme.In.HEADER)
                        .description("请求头")
        );
        return new OpenAPI()
                .components(components)
                .info(apiInfo());
    }

    private Info apiInfo() {
        Contact contact = new Contact();
        contact.setEmail("hjc_0827@163.com");
        contact.setName("Huang Jiacheng");
        return new Info()
                .title("Todo List Api")
                .description("swagger Documents")
                .version("1.0")
                .contact(contact);
    }
}
