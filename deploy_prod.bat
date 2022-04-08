set GOOS=linux
bee pack -ba "-tags prod" -exr="^(?:images|logs|temp|swagger)$"
D:\Development\PuTTY\pscp -pw 1nd0n3s1@PNM pkmapi.tar.gz svrapp01@10.61.3.233:/var/www/html/PKM_INISIASI
D:\Development\PuTTY\pscp -pw 1nd0n3s1@PNM pkmapi.tar.gz svrapp01@10.61.3.234:/var/www/html/PKM_INISIASI
D:\Development\PuTTY\pscp -pw 1nd0n3s1@PNM pkmapi.tar.gz svrapp01@10.61.3.235:/var/www/html/PKM_INISIASI
D:\Development\PuTTY\pscp -pw 1nd0n3s1@PNM pkmapi.tar.gz svrapp01@10.61.3.236:/var/www/html/PKM_INISIASI
ssh svrapp01@10.61.3.233
ssh svrapp01@10.61.3.234
ssh svrapp01@10.61.3.235
ssh svrapp01@10.61.3.236