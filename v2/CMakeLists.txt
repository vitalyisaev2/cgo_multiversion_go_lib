cmake_minimum_required(VERSION 3.10)

file(GLOB src "*.c")
file(GLOB headers "*.h")
add_library(cgomultiversion SHARED "")
set_target_properties(cgomultiversion PROPERTIES PUBLIC_HEADER "${headers}")
set_target_properties(cgomultiversion PROPERTIES SOVERSION "2")
set_target_properties(cgomultiversion PROPERTIES VERSION "2.0.0")
target_sources(cgomultiversion PRIVATE "${src}")

install(
    TARGETS cgomultiversion
    LIBRARY DESTINATION lib
    PUBLIC_HEADER DESTINATION include/cgomultiversion_v2
)