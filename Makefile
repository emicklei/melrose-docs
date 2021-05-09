dslmd:
	cd generators && go run *.go

wincp:
	gsutil cp gs://downloads.ernestmicklei.com/melrose/versions/windows/0.39.1/libstdc++-6.dll gs://downloads.ernestmicklei.com/melrose/versions/windows/0.40.0/
	gsutil cp gs://downloads.ernestmicklei.com/melrose/versions/windows/0.39.1/libwinpthread-1.dll gs://downloads.ernestmicklei.com/melrose/versions/windows/0.40.0/
	gsutil cp gs://downloads.ernestmicklei.com/melrose/versions/windows/0.39.1/libgcc_s_seh-1.dll gs://downloads.ernestmicklei.com/melrose/versions/windows/0.40.0/