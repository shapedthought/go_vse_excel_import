# GO VSE Excel Import

This program allows you to fill in the Excel template (included) with all your workloads and then converts it into a VSE compatible workload txt file.

Simply fill in the fields using lower case on any yes/no fields, then run the exe with:

    .\vse_excel_import.exe VSE_Workloads.xlsx

The result will be a vse_input_file.txt which can be loaded via Workloads > Import workloads in the VSE.

The format of the import is the same as the Excel Export 'Workload' page from the VSE so can be used as a method of updating multiple Workloads in existing estimations. For example:

    .\vse_excel_import.exe edited_output_report_excel.xlsx

Note that the settings used are the base ones in the VSE, if you wish to change anything I recommend opening the file in your favorite text editor or IDE.

## Installation

You can compile the program yourself via:

    go build .

You will need GO installed of course!

Or you can download the latest pre-compiled release >>

This program uses the following library for Excel import:

https://github.com/360EntSecGroup-Skylar/excelize