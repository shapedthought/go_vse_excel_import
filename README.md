# GO VSE Excel Import

This program allows you to fill in an Excel template with all your workloads and then converts it into a VSE compatible workload txt file.

Simply fill in the fields using lower case on any yes/no fields, then run the exe with:

    .\vse_excel_import.exe VSE_Workloads.xlsx

The result will be a vse_input_file.txt which can be loaded via Workloads > Import workloads in the VSE.

You can compile the program yourself via:

    go run build .

Or you can download the latest pre-compiled release.

This program uses the following library for Excel import:

https://github.com/360EntSecGroup-Skylar/excelize