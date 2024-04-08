exeargs =
args1 =

ifeq ($(OS),Windows_NT)
    EXE_EXT = .exe
else
    EXE_EXT =
endif

.PHONY: run
run: translate-api$(EXE_EXT)
	./translate-api$(EXE_EXT)

translate-api$(EXE_EXT): main.go handler.go
	$(args1) go build -o translate-api$(EXE_EXT) .
