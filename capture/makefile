# Compiler and flags
CC = gcc
# remove -Wpedantic for {} empty initializer
CFLAGS += -Wall -Wextra  -g \
          -Wformat=2 -Wno-unused-parameter -Wshadow \
          -Wwrite-strings -Wstrict-prototypes -Wold-style-definition \
          -Wredundant-decls -Wnested-externs -Wmissing-include-dirs
# Add linker flags here if needed
LDFLAGS = -lpcap
TARGET = capture

# Source files
SRC = capture.c 
# Object files (replace .c with .o)
OBJ = $(SRC:.c=.o)

# The first rule is the default target, typically the final program
all: $(TARGET)

# Rule to create the final executable
$(TARGET): $(OBJ)
	$(CC) $(OBJ) -o $(TARGET) $(LDFLAGS)

# Rule for compiling .c files into .o files (object files)
%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

# Clean up: Remove the generated files
clean:
	rm -f $(OBJ) $(TARGET)

