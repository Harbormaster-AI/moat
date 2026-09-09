import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.TrainingCourseDelegate import TrainingCourseDelegate

 #======================================================================
# 
# Encapsulates data for View TrainingCourse
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingCourseView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TrainingCourse index.")

def get(request, trainingCourseId ):
	delegate = TrainingCourseDelegate()
	responseData = delegate.get( trainingCourseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	trainingCourse = json.loads(request.body)
	delegate = TrainingCourseDelegate()
	responseData = delegate.createFromJson( trainingCourse )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	trainingCourse = json.loads(request.body)
	delegate = TrainingCourseDelegate()
	responseData = delegate.save( trainingCourse )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, trainingCourseId ):
	delegate = TrainingCourseDelegate()
	responseData = delegate.delete( trainingCourseId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TrainingCourseDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPrerequisites( request, trainingCourseId, PrerequisitesIds ):
	delegate = TrainingCourseDelegate()
	responseData = delegate.addPrerequisites( trainingCourseId, PrerequisitesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePrerequisites( request, trainingCourseId, PrerequisitesIds ):
	delegate = TrainingCourseDelegate()
	responseData = delegate.removePrerequisites( trainingCourseId, PrerequisitesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEnrollments( request, trainingCourseId, EnrollmentsIds ):
	delegate = TrainingCourseDelegate()
	responseData = delegate.addEnrollments( trainingCourseId, EnrollmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEnrollments( request, trainingCourseId, EnrollmentsIds ):
	delegate = TrainingCourseDelegate()
	responseData = delegate.removeEnrollments( trainingCourseId, EnrollmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addJobProfiles( request, trainingCourseId, JobProfilesIds ):
	delegate = TrainingCourseDelegate()
	responseData = delegate.addJobProfiles( trainingCourseId, JobProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeJobProfiles( request, trainingCourseId, JobProfilesIds ):
	delegate = TrainingCourseDelegate()
	responseData = delegate.removeJobProfiles( trainingCourseId, JobProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

