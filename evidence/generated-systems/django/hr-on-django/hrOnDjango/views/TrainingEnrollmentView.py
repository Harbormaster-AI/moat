import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.TrainingEnrollmentDelegate import TrainingEnrollmentDelegate

 #======================================================================
# 
# Encapsulates data for View TrainingEnrollment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingEnrollmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TrainingEnrollment index.")

def get(request, trainingEnrollmentId ):
	delegate = TrainingEnrollmentDelegate()
	responseData = delegate.get( trainingEnrollmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	trainingEnrollment = json.loads(request.body)
	delegate = TrainingEnrollmentDelegate()
	responseData = delegate.createFromJson( trainingEnrollment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	trainingEnrollment = json.loads(request.body)
	delegate = TrainingEnrollmentDelegate()
	responseData = delegate.save( trainingEnrollment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, trainingEnrollmentId ):
	delegate = TrainingEnrollmentDelegate()
	responseData = delegate.delete( trainingEnrollmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TrainingEnrollmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCourse( request, trainingEnrollmentId, CourseId ):
	delegate = TrainingEnrollmentDelegate()
	responseData = delegate.saveCourse( trainingEnrollmentId, CourseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCourse( request, trainingEnrollmentId ):
	delegate = TrainingEnrollmentDelegate()
	responseData = delegate.deleteCourse( trainingEnrollmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, trainingEnrollmentId, EmployeeId ):
	delegate = TrainingEnrollmentDelegate()
	responseData = delegate.saveEmployee( trainingEnrollmentId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, trainingEnrollmentId ):
	delegate = TrainingEnrollmentDelegate()
	responseData = delegate.deleteEmployee( trainingEnrollmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInstructor( request, trainingEnrollmentId, InstructorId ):
	delegate = TrainingEnrollmentDelegate()
	responseData = delegate.saveInstructor( trainingEnrollmentId, InstructorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInstructor( request, trainingEnrollmentId ):
	delegate = TrainingEnrollmentDelegate()
	responseData = delegate.deleteInstructor( trainingEnrollmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

