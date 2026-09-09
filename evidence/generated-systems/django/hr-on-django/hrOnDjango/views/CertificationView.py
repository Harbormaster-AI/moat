import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.CertificationDelegate import CertificationDelegate

 #======================================================================
# 
# Encapsulates data for View Certification
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CertificationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Certification index.")

def get(request, certificationId ):
	delegate = CertificationDelegate()
	responseData = delegate.get( certificationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	certification = json.loads(request.body)
	delegate = CertificationDelegate()
	responseData = delegate.createFromJson( certification )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	certification = json.loads(request.body)
	delegate = CertificationDelegate()
	responseData = delegate.save( certification )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, certificationId ):
	delegate = CertificationDelegate()
	responseData = delegate.delete( certificationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CertificationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, certificationId, EmployeeId ):
	delegate = CertificationDelegate()
	responseData = delegate.saveEmployee( certificationId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, certificationId ):
	delegate = CertificationDelegate()
	responseData = delegate.deleteEmployee( certificationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCourse( request, certificationId, CourseId ):
	delegate = CertificationDelegate()
	responseData = delegate.saveCourse( certificationId, CourseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCourse( request, certificationId ):
	delegate = CertificationDelegate()
	responseData = delegate.deleteCourse( certificationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

