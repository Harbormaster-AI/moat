import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.JobRequisitionDelegate import JobRequisitionDelegate

 #======================================================================
# 
# Encapsulates data for View JobRequisition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobRequisitionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the JobRequisition index.")

def get(request, jobRequisitionId ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.get( jobRequisitionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	jobRequisition = json.loads(request.body)
	delegate = JobRequisitionDelegate()
	responseData = delegate.createFromJson( jobRequisition )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	jobRequisition = json.loads(request.body)
	delegate = JobRequisitionDelegate()
	responseData = delegate.save( jobRequisition )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, jobRequisitionId ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.delete( jobRequisitionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = JobRequisitionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDepartment( request, jobRequisitionId, DepartmentId ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.saveDepartment( jobRequisitionId, DepartmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDepartment( request, jobRequisitionId ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.deleteDepartment( jobRequisitionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignHiringManager( request, jobRequisitionId, HiringManagerId ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.saveHiringManager( jobRequisitionId, HiringManagerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignHiringManager( request, jobRequisitionId ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.deleteHiringManager( jobRequisitionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRecruiter( request, jobRequisitionId, RecruiterId ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.saveRecruiter( jobRequisitionId, RecruiterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRecruiter( request, jobRequisitionId ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.deleteRecruiter( jobRequisitionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignJobProfile( request, jobRequisitionId, JobProfileId ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.saveJobProfile( jobRequisitionId, JobProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignJobProfile( request, jobRequisitionId ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.deleteJobProfile( jobRequisitionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCandidates( request, jobRequisitionId, CandidatesIds ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.addCandidates( jobRequisitionId, CandidatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCandidates( request, jobRequisitionId, CandidatesIds ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.removeCandidates( jobRequisitionId, CandidatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInterviews( request, jobRequisitionId, InterviewsIds ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.addInterviews( jobRequisitionId, InterviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInterviews( request, jobRequisitionId, InterviewsIds ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.removeInterviews( jobRequisitionId, InterviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOffers( request, jobRequisitionId, OffersIds ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.addOffers( jobRequisitionId, OffersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOffers( request, jobRequisitionId, OffersIds ):
	delegate = JobRequisitionDelegate()
	responseData = delegate.removeOffers( jobRequisitionId, OffersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

