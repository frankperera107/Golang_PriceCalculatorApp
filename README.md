Simple Golang project that demonstrates the good use of interfaces, structs and modular design principles. 

Key Highlights:

1. Modular Design: Organized into multiple packages (cmdmanager, filemanager, prices, etc.) for clear separation of concerns.
2. Flexible I/O Handling: Uses an IOManager interface, allowing easy swapping between file-based (FileManager) and command-line (CMDManager) input/output.
3. Reusable Components: Easily extendable for different tax rates and input methods.
4. Error Handling: Consistent error checks ensure robust performance.

Key features of the app:  
-Input prices from a file or command line.
-Process prices with customizable tax rates.
-Output results in JSON format or display them on the console.
