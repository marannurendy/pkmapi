
@REM RUNNING DI CMD

set GOOS=linux
bee pack -ba "-tags prod" -exr="^(?:images|logs|temp)$"
D:\Development\PuTTY\pscp -pw 1nd0n3s1@PNM pkmapi.tar.gz svrapp01@10.61.3.221:/var/www/html/PKM_mekaar
ssh svrapp01@10.61.3.221

@REM E:\Project\PuTTY\plink svrapp01@10.61.3.221 -pw 1nd0n3s1@PNM -m remote.txt 


@REM E:\Project\PuTTY\pscp -pw PNMdemo pkmapi.tar.gz root@103.168.134.78:/home/ubuntu/PKM_INISIASI