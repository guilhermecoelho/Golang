package exemplo

import (
	"strings"
	"testing"
)

func TestImportFile(t *testing.T) {

	expectedResult := strings.Join(populateExpectedResult(), ", ")
	result := strings.Join(Readfile(), ", ")

	if strings.Compare(result, expectedResult) == -1 {
		t.Errorf("ImportFile failed, expected %v, got %v", expectedResult, result)
	} else {
		t.Logf("ImportFile success, expected %v, got %v", expectedResult, result)
	}

}

func populateExpectedResult() []string {

	expectedResult := []string{}
	expectedResult = append(expectedResult, "31-12-2020;31-12-2020;COMPRA SIRIUS FUEL LD ;30,00;;877,07;844,22;COMPRAS ;\n")
	expectedResult = append(expectedResult, "31-12-2020;31-12-2020;TRF CXDAPP ;60,00;;907,07;874,22;Diversos ;\n")
	expectedResult = append(expectedResult, "30-12-2020;30-12-2020;COMPRA CONTINENTE V A ;18,57;;967,07;874,22;COMPRAS ;\n")
	expectedResult = append(expectedResult, "30-12-2020;30-12-2020;LEVANTAMENTO Av Infan ;150,00;;985,64;892,79;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "30-12-2020;30-12-2020;TRANSFERENCIA ;200,00;;1.135,64;1.042,79;Diversos ;\n")
	expectedResult = append(expectedResult, "30-12-2020;30-12-2020;TRF EVA SOCIEDADE HOT ;;598,45;1.335,64;1.302,79;Diversos ;\n")
	expectedResult = append(expectedResult, "29-12-2020;29-12-2020;PAGAMENTO ;46,32;;737,19;737,19;AGUA ;\n")
	expectedResult = append(expectedResult, "29-12-2020;29-12-2020;COMPRA AVALANCHE CORE ;3,10;;783,51;783,51;COMPRAS ;\n")
	expectedResult = append(expectedResult, "29-12-2020;29-12-2020;TRF P2P 910xxx898 ;12,00;;786,61;786,61;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "29-12-2020;29-12-2020;COMPRA FORCA PORTUGAL ;19,99;;798,61;798,61;COMPRAS ;\n")
	expectedResult = append(expectedResult, "29-12-2020;29-12-2020;COMPRA RESTAURANTE PI ;28,50;;818,60;818,60;COMPRAS ;\n")
	expectedResult = append(expectedResult, "29-12-2020;29-12-2020;EDP COMERCIAL COMERCI ;53,64;;847,10;847,10;Diversos ;\n")
	expectedResult = append(expectedResult, "28-12-2020;28-12-2020;FAGAR FARO  Gestao de ;36,71;;900,74;900,74;AGUA ;\n")
	expectedResult = append(expectedResult, "28-12-2020;28-12-2020;FIDELIDADE   COMPANHI ;27,04;;937,45;937,45;SEGUROS ;\n")
	expectedResult = append(expectedResult, "27-12-2020;27-12-2020;COMPRA ALDI LJ014 FAR ;27,01;;964,49;964,49;COMPRAS ;\n")
	expectedResult = append(expectedResult, "27-12-2020;27-12-2020;COMPRA THE WOODS ;8,50;;991,50;991,50;COMPRAS ;\n")
	expectedResult = append(expectedResult, "24-12-2020;24-12-2020;TRF CXDAPP ;596,06;;1.000,00;1.000,00;Diversos ;\n")
	expectedResult = append(expectedResult, "24-12-2020;24-12-2020;TRF HIQ CONSULTING ;;1.316,19;1.596,06;1.596,06;Diversos ;\n")
	expectedResult = append(expectedResult, "23-12-2020;23-12-2020;COMPRA ARESTAROMATICA ;12,40;;279,87;279,87;COMPRAS ;\n")
	expectedResult = append(expectedResult, "23-12-2020;23-12-2020;ANUL COMPRA CENTROXOG ;;23,95;292,27;292,27;Diversos ;\n")
	expectedResult = append(expectedResult, "23-12-2020;23-12-2020;COMPRA CENTROXOGO ALG ;23,95;;268,32;268,32;COMPRAS ;\n")
	expectedResult = append(expectedResult, "23-12-2020;23-12-2020;ANUL COMPRA CENTROXOG ;;23,95;292,27;292,27;Diversos ;\n")
	expectedResult = append(expectedResult, "23-12-2020;23-12-2020;COMPRA CENTROXOGO ALG ;23,95;;268,32;268,32;COMPRAS ;\n")
	expectedResult = append(expectedResult, "23-12-2020;23-12-2020;TRF P2P 913xxx289 ;8,00;;292,27;292,27;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "23-12-2020;23-12-2020;COMPRA DONER KEBAB ;6,45;;300,27;300,27;COMPRAS ;\n")
	expectedResult = append(expectedResult, "23-12-2020;23-12-2020;COMPRA LIDL   CIA ;96,16;;306,72;306,72;COMPRAS ;\n")
	expectedResult = append(expectedResult, "23-12-2020;21-12-2020;COMPRA PC COMPONENTES ;227,24;;402,88;402,88;COMPRAS ;\n")
	expectedResult = append(expectedResult, "22-12-2020;22-12-2020;COBRANCA PRESTACAO ;122,38;;630,12;402,88;Diversos ;\n")
	expectedResult = append(expectedResult, "22-12-2020;22-12-2020;COMPRA REST CIDADE VE ;25,00;;752,50;525,26;COMPRAS ;\n")
	expectedResult = append(expectedResult, "22-12-2020;22-12-2020;COMPRA H3 PORTUGAL ;8,60;;777,50;550,26;COMPRAS ;\n")
	expectedResult = append(expectedResult, "22-12-2020;22-12-2020;COMPRA WOMEN SECRET F ;18,24;;786,10;558,86;COMPRAS ;\n")
	expectedResult = append(expectedResult, "22-12-2020;22-12-2020;COMPRA FORUM ALGARVE ;82,90;;804,34;577,10;COMPRAS ;\n")
	expectedResult = append(expectedResult, "22-12-2020;22-12-2020;TRF P2P 912xxx513 ;40,00;;887,24;660,00;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "22-12-2020;22-12-2020;TRANSFERENCIA ;;300,00;927,24;700,00;Diversos ;\n")
	expectedResult = append(expectedResult, "22-12-2020;22-12-2020;TRANSFERENCIA ;;400,00;627,24;400,00;Diversos ;\n")
	expectedResult = append(expectedResult, "21-12-2020;21-12-2020;COBRANCA PRESTACAO ;90,06;;227,24;0,00;Diversos ;\n")
	expectedResult = append(expectedResult, "21-12-2020;21-12-2020;COMPRA SUPERMERCADO ;3,49;;317,30;90,06;COMPRAS ;\n")
	expectedResult = append(expectedResult, "21-12-2020;21-12-2020;COMPRA DONER KEBAB ;6,45;;320,79;93,55;COMPRAS ;\n")
	expectedResult = append(expectedResult, "21-12-2020;21-12-2020;TRANSFERENCIA ;;100,00;327,24;100,00;Diversos ;\n")
	expectedResult = append(expectedResult, "21-12-2020;21-12-2020;COBRANCA PRESTACAO ;98,86;;227,24;0,00;Diversos ;\n")
	expectedResult = append(expectedResult, "20-12-2020;20-12-2020;COMPRA BCM BRICOLAGE ;683,18;;326,10;326,10;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "20-12-2020;20-12-2020;TRANSFERENCIA ;;783,18;1.009,28;1.009,28;Diversos ;\n")
	expectedResult = append(expectedResult, "20-12-2020;20-12-2020;COMPRA SIRIUS FUEL LD ;30,00;;226,10;226,10;COMPRAS ;\n")
	expectedResult = append(expectedResult, "20-12-2020;20-12-2020;COMPRA ALDI LJ014 FAR ;37,36;;256,10;256,10;COMPRAS ;\n")
	expectedResult = append(expectedResult, "20-12-2020;20-12-2020;COMPRA BAJA FRESH ;11,45;;293,46;293,46;COMPRAS ;\n")
	expectedResult = append(expectedResult, "20-12-2020;20-12-2020;COMPRA LEROY MERLIN L ;29,24;;304,91;304,91;COMPRAS ;\n")
	expectedResult = append(expectedResult, "19-12-2020;19-12-2020;COMPRA KFC FORUM ALGA ;8,00;;334,15;334,15;COMPRAS ;\n")
	expectedResult = append(expectedResult, "18-12-2020;18-12-2020;MEO  SA ;39,09;;342,15;342,15;Diversos ;\n")
	expectedResult = append(expectedResult, "18-12-2020;14-12-2020;COMPRAS C DEB IKEA LO ;7,85;;381,24;381,24;COMPRAS ;\n")
	expectedResult = append(expectedResult, "17-12-2020;17-12-2020;COMPRA ARESTAROMATICA ;12,40;;389,09;381,24;COMPRAS ;\n")
	expectedResult = append(expectedResult, "17-12-2020;13-12-2020;COMPRA WWW WORTEN PT ;205,98;;401,49;393,64;COMPRAS ;\n")
	expectedResult = append(expectedResult, "16-12-2020;13-12-2020;COMPRA IKEA PORTUGAL ;243,00;;607,47;393,64;COMPRAS ;\n")
	expectedResult = append(expectedResult, "15-12-2020;15-12-2020;TRF P2P 912xxx513 ;29,00;;850,47;393,64;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "14-12-2020;14-12-2020;LEVANTAMENTO Mar Shop ;20,00;;879,47;422,64;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "14-12-2020;14-12-2020;LEVANTAMENTO Mar Shop ;20,00;;899,47;442,64;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "14-12-2020;11-12-2020;COMPRA HBO PORTUGAL ;4,99;;919,47;470,49;COMPRAS ;\n")
	expectedResult = append(expectedResult, "10-12-2020;10-12-2020;COMPRA LEROY MERLIN ;14,91;;924,46;924,46;COMPRAS ;\n")
	expectedResult = append(expectedResult, "09-12-2020;09-12-2020;COMPRA LIDL   CIA ;31,60;;939,37;939,37;COMPRAS ;\n")
	expectedResult = append(expectedResult, "09-12-2020;09-12-2020;COMPRA FORUM ALGARVE ;102,90;;970,97;970,97;COMPRAS ;\n")
	expectedResult = append(expectedResult, "08-12-2020;08-12-2020;COMPRA AQUAPICANCO ;14,50;;1.073,87;1.073,87;COMPRAS ;\n")
	expectedResult = append(expectedResult, "08-12-2020;08-12-2020;COMPRA STARBUCKS COFF ;8,65;;1.088,37;1.088,37;COMPRAS ;\n")
	expectedResult = append(expectedResult, "06-12-2020;06-12-2020;COMPRA PINGO DOCE ;11,20;;1.097,02;1.097,02;COMPRAS ;\n")
	expectedResult = append(expectedResult, "05-12-2020;05-12-2020;COMISSAO CONTA CAIXA ;4,16;;1.108,22;1.108,22;Diversos ;\n")
	expectedResult = append(expectedResult, "04-12-2020;04-12-2020;LEVANTAMENTO ForumAlg ;20,00;;1.112,38;1.112,38;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "04-12-2020;04-12-2020;NOS Comunicacoes ;65,40;;1.132,38;1.132,38;TELE / TV / INTERNET ;\n")
	expectedResult = append(expectedResult, "02-12-2020;02-12-2020;COMPRA ALDI LJ014 FAR ;85,76;;1.197,78;1.197,78;COMPRAS ;\n")
	expectedResult = append(expectedResult, "02-12-2020;02-12-2020;TRF P2P 913xxx289 ;8,00;;1.283,54;1.283,54;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "02-12-2020;02-12-2020;TRF P2P 967xxx113 ;20,00;;1.291,54;1.291,54;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "02-12-2020;02-12-2020;DEPOSITO ;;40,00;1.311,54;1.311,54;Diversos ;\n")
	expectedResult = append(expectedResult, "02-12-2020;02-12-2020;REBELAMBITION LDA ;36,00;;1.271,54;1.271,54;Diversos ;\n")
	expectedResult = append(expectedResult, "01-12-2020;01-12-2020;COMPRA SIRIUS FUEL LD ;30,00;;1.307,54;1.307,54;COMPRAS ;\n")
	expectedResult = append(expectedResult, "01-12-2020;01-12-2020;LEVANTAMENTO Estr Sta ;40,00;;1.337,54;1.337,54;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "01-12-2020;01-12-2020;LEVANTAMENTO Estr Sta ;150,00;;1.377,54;1.377,54;LEVANTAMENTOS ;\n")
	expectedResult = append(expectedResult, "01-12-2020;01-12-2020;COMPRA PRAISUPER PROD ;21,91;;1.527,54;1.527,54;COMPRAS ;\n")
	expectedResult = append(expectedResult, "01-12-2020;01-12-2020;Medicare ;19,90;;1.549,45;1.549,45;Diversos ;\n")
	expectedResult = append(expectedResult, "01-12-2020;01-12-2020;DECO PROTESTE EDITORE ;7,30;;1.569,35;1.569,35;Diversos ;\n")
	expectedResult = append(expectedResult, "01-12-2020;01-12-2020;DECO PROTESTE EDITORE ;7,30;;1.576,65;1.576,65;Diversos ;\n")
	expectedResult = append(expectedResult, " ; ; ; ;Saldo contabilístico ;25,70 EUR ; ; ;\n")

	return expectedResult

}
