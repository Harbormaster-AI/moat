import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.TypeCertificateDelegate import TypeCertificateDelegate

 #======================================================================
# 
# Encapsulates data for View TypeCertificate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TypeCertificateView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TypeCertificate index.")

def get(request, typeCertificateId ):
	delegate = TypeCertificateDelegate()
	responseData = delegate.get( typeCertificateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	typeCertificate = json.loads(request.body)
	delegate = TypeCertificateDelegate()
	responseData = delegate.createFromJson( typeCertificate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	typeCertificate = json.loads(request.body)
	delegate = TypeCertificateDelegate()
	responseData = delegate.save( typeCertificate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, typeCertificateId ):
	delegate = TypeCertificateDelegate()
	responseData = delegate.delete( typeCertificateId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TypeCertificateDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProgram( request, typeCertificateId, ProgramId ):
	delegate = TypeCertificateDelegate()
	responseData = delegate.saveProgram( typeCertificateId, ProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProgram( request, typeCertificateId ):
	delegate = TypeCertificateDelegate()
	responseData = delegate.deleteProgram( typeCertificateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

