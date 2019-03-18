package noire

var colorNames = map[string]string{
	"ALICEBLUE":            "F0F8FF",
	"ANTIQUEWHITE":         "FAEBD7",
	"AQUA":                 "00FFFF",
	"AQUAMARINE":           "7FFFD4",
	"AZURE":                "F0FFFF",
	"BEIGE":                "F5F5DC",
	"BISQUE":               "FFE4C4",
	"BLACK":                "000000",
	"BLANCHEDALMOND":       "FFEBCD",
	"BLUE":                 "0000FF",
	"BLUEVIOLET":           "8A2BE2",
	"BROWN":                "A52A2A",
	"BURLYWOOD":            "DEB887",
	"CADETBLUE":            "5F9EA0",
	"CHARTREUSE":           "7FFF00",
	"CHOCOLATE":            "D2691E",
	"CORAL":                "FF7F50",
	"CORNFLOWERBLUE":       "6495ED",
	"CORNSILK":             "FFF8DC",
	"CRIMSON":              "DC143C",
	"CYAN":                 "00FFFF",
	"DARKBLUE":             "00008B",
	"DARKCYAN":             "008B8B",
	"DARKGOLDENROD":        "B8860B",
	"DARKGRAY":             "A9A9A9",
	"DARKGREY":             "A9A9A9",
	"DARKGREEN":            "006400",
	"DARKKHAKI":            "BDB76B",
	"DARKMAGENTA":          "8B008B",
	"DARKOLIVEGREEN":       "556B2F",
	"DARKORANGE":           "FF8C00",
	"DARKORCHID":           "9932CC",
	"DARKRED":              "8B0000",
	"DARKSALMON":           "E9967A",
	"DARKSEAGREEN":         "8FBC8F",
	"DARKSLATEBLUE":        "483D8B",
	"DARKSLATEGRAY":        "2F4F4F",
	"DARKSLATEGREY":        "2F4F4F",
	"DARKTURQUOISE":        "00CED1",
	"DARKVIOLET":           "9400D3",
	"DEEPPINK":             "FF1493",
	"DEEPSKYBLUE":          "00BFFF",
	"DIMGRAY":              "696969",
	"DIMGREY":              "696969",
	"DODGERBLUE":           "1E90FF",
	"FIREBRICK":            "B22222",
	"FLORALWHITE":          "FFFAF0",
	"FORESTGREEN":          "228B22",
	"FUCHSIA":              "FF00FF",
	"GAINSBORO":            "DCDCDC",
	"GHOSTWHITE":           "F8F8FF",
	"GOLD":                 "FFD700",
	"GOLDENROD":            "DAA520",
	"GRAY":                 "808080",
	"GREY":                 "808080",
	"GREEN":                "008000",
	"GREENYELLOW":          "ADFF2F",
	"HONEYDEW":             "F0FFF0",
	"HOTPINK":              "FF69B4",
	"INDIANRED ":           "CD5C5C",
	"INDIGO ":              "4B0082",
	"IVORY":                "FFFFF0",
	"KHAKI":                "F0E68C",
	"LAVENDER":             "E6E6FA",
	"LAVENDERBLUSH":        "FFF0F5",
	"LAWNGREEN":            "7CFC00",
	"LEMONCHIFFON":         "FFFACD",
	"LIGHTBLUE":            "ADD8E6",
	"LIGHTCORAL":           "F08080",
	"LIGHTCYAN":            "E0FFFF",
	"LIGHTGOLDENRODYELLOW": "FAFAD2",
	"LIGHTGRAY":            "D3D3D3",
	"LIGHTGREY":            "D3D3D3",
	"LIGHTGREEN":           "90EE90",
	"LIGHTPINK":            "FFB6C1",
	"LIGHTSALMON":          "FFA07A",
	"LIGHTSEAGREEN":        "20B2AA",
	"LIGHTSKYBLUE":         "87CEFA",
	"LIGHTSLATEGRAY":       "778899",
	"LIGHTSLATEGREY":       "778899",
	"LIGHTSTEELBLUE":       "B0C4DE",
	"LIGHTYELLOW":          "FFFFE0",
	"LIME":                 "00FF00",
	"LIMEGREEN":            "32CD32",
	"LINEN":                "FAF0E6",
	"MAGENTA":              "FF00FF",
	"MAROON":               "800000",
	"MEDIUMAQUAMARINE":     "66CDAA",
	"MEDIUMBLUE":           "0000CD",
	"MEDIUMORCHID":         "BA55D3",
	"MEDIUMPURPLE":         "9370DB",
	"MEDIUMSEAGREEN":       "3CB371",
	"MEDIUMSLATEBLUE":      "7B68EE",
	"MEDIUMSPRINGGREEN":    "00FA9A",
	"MEDIUMTURQUOISE":      "48D1CC",
	"MEDIUMVIOLETRED":      "C71585",
	"MIDNIGHTBLUE":         "191970",
	"MINTCREAM":            "F5FFFA",
	"MISTYROSE":            "FFE4E1",
	"MOCCASIN":             "FFE4B5",
	"NAVAJOWHITE":          "FFDEAD",
	"NAVY":                 "000080",
	"OLDLACE":              "FDF5E6",
	"OLIVE":                "808000",
	"OLIVEDRAB":            "6B8E23",
	"ORANGE":               "FFA500",
	"ORANGERED":            "FF4500",
	"ORCHID":               "DA70D6",
	"PALEGOLDENROD":        "EEE8AA",
	"PALEGREEN":            "98FB98",
	"PALETURQUOISE":        "AFEEEE",
	"PALEVIOLETRED":        "DB7093",
	"PAPAYAWHIP":           "FFEFD5",
	"PEACHPUFF":            "FFDAB9",
	"PERU":                 "CD853F",
	"PINK":                 "FFC0CB",
	"PLUM":                 "DDA0DD",
	"POWDERBLUE":           "B0E0E6",
	"PURPLE":               "800080",
	"REBECCAPURPLE":        "663399",
	"RED":                  "FF0000",
	"ROSYBROWN":            "BC8F8F",
	"ROYALBLUE":            "4169E1",
	"SADDLEBROWN":          "8B4513",
	"SALMON":               "FA8072",
	"SANDYBROWN":           "F4A460",
	"SEAGREEN":             "2E8B57",
	"SEASHELL":             "FFF5EE",
	"SIENNA":               "A0522D",
	"SILVER":               "C0C0C0",
	"SKYBLUE":              "87CEEB",
	"SLATEBLUE":            "6A5ACD",
	"SLATEGRAY":            "708090",
	"SLATEGREY":            "708090",
	"SNOW":                 "FFFAFA",
	"SPRINGGREEN":          "00FF7F",
	"STEELBLUE":            "4682B4",
	"TAN":                  "D2B48C",
	"TEAL":                 "008080",
	"THISTLE":              "D8BFD8",
	"TOMATO":               "FF6347",
	"TURQUOISE":            "40E0D0",
	"VIOLET":               "EE82EE",
	"WHEAT":                "F5DEB3",
	"WHITE":                "FFFFFF",
	"WHITESMOKE":           "F5F5F5",
	"YELLOW":               "FFFF00",
	"YELLOWGREEN":          "9ACD32",
}

var hexNames = map[string]string{
	"F0F8FF": "AliceBlue",
	"FAEBD7": "AntiqueWhite",
	"7FFFD4": "Aquamarine",
	"F0FFFF": "Azure",
	"F5F5DC": "Beige",
	"FFE4C4": "Bisque",
	"000000": "Black",
	"FFEBCD": "BlanchedAlmond",
	"0000FF": "Blue",
	"8A2BE2": "BlueViolet",
	"A52A2A": "Brown",
	"DEB887": "BurlyWood",
	"5F9EA0": "CadetBlue",
	"7FFF00": "Chartreuse",
	"D2691E": "Chocolate",
	"FF7F50": "Coral",
	"6495ED": "CornflowerBlue",
	"FFF8DC": "Cornsilk",
	"DC143C": "Crimson",
	"00FFFF": "Cyan",
	"00008B": "DarkBlue",
	"008B8B": "DarkCyan",
	"B8860B": "DarkGoldenRod",
	"006400": "DarkGreen",
	"BDB76B": "DarkKhaki",
	"8B008B": "DarkMagenta",
	"556B2F": "DarkOliveGreen",
	"FF8C00": "DarkOrange",
	"9932CC": "DarkOrchid",
	"8B0000": "DarkRed",
	"E9967A": "DarkSalmon",
	"8FBC8F": "DarkSeaGreen",
	"483D8B": "DarkSlateBlue",
	"2F4F4F": "DarkSlateGray",
	"00CED1": "DarkTurquoise",
	"9400D3": "DarkViolet",
	"FF1493": "DeepPink",
	"00BFFF": "DeepSkyBlue",
	"696969": "DimGray",
	"1E90FF": "DodgerBlue",
	"B22222": "FireBrick",
	"FFFAF0": "FloralWhite",
	"228B22": "ForestGreen",
	"DCDCDC": "Gainsboro",
	"F8F8FF": "GhostWhite",
	"FFD700": "Gold",
	"DAA520": "GoldenRod",
	"808080": "Gray",
	"008000": "Green",
	"ADFF2F": "GreenYellow",
	"F0FFF0": "HoneyDew",
	"FF69B4": "HotPink",
	"CD5C5C": "IndianRed ",
	"4B0082": "Indigo ",
	"FFFFF0": "Ivory",
	"F0E68C": "Khaki",
	"E6E6FA": "Lavender",
	"FFF0F5": "LavenderBlush",
	"7CFC00": "LawnGreen",
	"FFFACD": "LemonChiffon",
	"ADD8E6": "LightBlue",
	"F08080": "LightCoral",
	"E0FFFF": "LightCyan",
	"FAFAD2": "LightGoldenRodYellow",
	"D3D3D3": "LightGray",
	"90EE90": "LightGreen",
	"FFB6C1": "LightPink",
	"FFA07A": "LightSalmon",
	"20B2AA": "LightSeaGreen",
	"87CEFA": "LightSkyBlue",
	"778899": "LightSlateGray",
	"B0C4DE": "LightSteelBlue",
	"FFFFE0": "LightYellow",
	"00FF00": "Lime",
	"32CD32": "LimeGreen",
	"FAF0E6": "Linen",
	"FF00FF": "Magenta",
	"800000": "Maroon",
	"66CDAA": "MediumAquaMarine",
	"0000CD": "MediumBlue",
	"BA55D3": "MediumOrchid",
	"9370DB": "MediumPurple",
	"3CB371": "MediumSeaGreen",
	"7B68EE": "MediumSlateBlue",
	"00FA9A": "MediumSpringGreen",
	"48D1CC": "MediumTurquoise",
	"C71585": "MediumVioletRed",
	"191970": "MidnightBlue",
	"F5FFFA": "MintCream",
	"FFE4E1": "MistyRose",
	"FFE4B5": "Moccasin",
	"FFDEAD": "NavajoWhite",
	"000080": "Navy",
	"FDF5E6": "OldLace",
	"808000": "Olive",
	"6B8E23": "OliveDrab",
	"FFA500": "Orange",
	"FF4500": "OrangeRed",
	"DA70D6": "Orchid",
	"EEE8AA": "PaleGoldenRod",
	"98FB98": "PaleGreen",
	"AFEEEE": "PaleTurquoise",
	"DB7093": "PaleVioletRed",
	"FFEFD5": "PapayaWhip",
	"FFDAB9": "PeachPuff",
	"CD853F": "Peru",
	"FFC0CB": "Pink",
	"DDA0DD": "Plum",
	"B0E0E6": "PowderBlue",
	"800080": "Purple",
	"663399": "RebeccaPurple",
	"FF0000": "Red",
	"BC8F8F": "RosyBrown",
	"4169E1": "RoyalBlue",
	"8B4513": "SaddleBrown",
	"FA8072": "Salmon",
	"F4A460": "SandyBrown",
	"2E8B57": "SeaGreen",
	"FFF5EE": "SeaShell",
	"A0522D": "Sienna",
	"C0C0C0": "Silver",
	"87CEEB": "SkyBlue",
	"6A5ACD": "SlateBlue",
	"708090": "SlateGray",
	"FFFAFA": "Snow",
	"00FF7F": "SpringGreen",
	"4682B4": "SteelBlue",
	"D2B48C": "Tan",
	"008080": "Teal",
	"D8BFD8": "Thistle",
	"FF6347": "Tomato",
	"40E0D0": "Turquoise",
	"EE82EE": "Violet",
	"F5DEB3": "Wheat",
	"FFFFFF": "White",
	"F5F5F5": "WhiteSmoke",
	"FFFF00": "Yellow",
	"9ACD32": "YellowGreen",
}
