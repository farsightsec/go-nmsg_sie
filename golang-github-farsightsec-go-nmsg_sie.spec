%global debug_package %{nil}

# https://github.com/farsightsec/go-nmsg_sie
%global goipath         github.com/farsightsec/go-nmsg_sie
Version:                0.1.1

%gometa

%global common_description %{expand:
Provides definitions for message types from sie-nmsg for use with the go-nmsg library.}

%global godocs          README.md

Name:           %{goname}
Release:        %autorelease
Summary:        SIE Message Module for go-nmsg

License:        MPLv2.0
URL:            %{gourl}
Source0:        %{gosource}

%description
%{common_description}

%gopkg

%prep
%goprep

%generate_buildrequires
%go_generate_buildrequires

%install
%gopkginstall

%if %{with check}
%check
%gocheck
%endif

%gopkgfiles

%changelog
%autochangelog
