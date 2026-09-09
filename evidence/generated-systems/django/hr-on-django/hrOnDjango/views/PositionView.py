import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.PositionDelegate import PositionDelegate

 #======================================================================
# 
# Encapsulates data for View Position
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PositionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Position index.")

def get(request, positionId ):
	delegate = PositionDelegate()
	responseData = delegate.get( positionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	position = json.loads(request.body)
	delegate = PositionDelegate()
	responseData = delegate.createFromJson( position )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	position = json.loads(request.body)
	delegate = PositionDelegate()
	responseData = delegate.save( position )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, positionId ):
	delegate = PositionDelegate()
	responseData = delegate.delete( positionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PositionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDepartment( request, positionId, DepartmentId ):
	delegate = PositionDelegate()
	responseData = delegate.saveDepartment( positionId, DepartmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDepartment( request, positionId ):
	delegate = PositionDelegate()
	responseData = delegate.deleteDepartment( positionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignJobProfile( request, positionId, JobProfileId ):
	delegate = PositionDelegate()
	responseData = delegate.saveJobProfile( positionId, JobProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignJobProfile( request, positionId ):
	delegate = PositionDelegate()
	responseData = delegate.deleteJobProfile( positionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCostCenter( request, positionId, CostCenterId ):
	delegate = PositionDelegate()
	responseData = delegate.saveCostCenter( positionId, CostCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCostCenter( request, positionId ):
	delegate = PositionDelegate()
	responseData = delegate.deleteCostCenter( positionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLocation( request, positionId, LocationId ):
	delegate = PositionDelegate()
	responseData = delegate.saveLocation( positionId, LocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLocation( request, positionId ):
	delegate = PositionDelegate()
	responseData = delegate.deleteLocation( positionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignManagerPosition( request, positionId, ManagerPositionId ):
	delegate = PositionDelegate()
	responseData = delegate.saveManagerPosition( positionId, ManagerPositionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignManagerPosition( request, positionId ):
	delegate = PositionDelegate()
	responseData = delegate.deleteManagerPosition( positionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDirectReports( request, positionId, DirectReportsIds ):
	delegate = PositionDelegate()
	responseData = delegate.addDirectReports( positionId, DirectReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDirectReports( request, positionId, DirectReportsIds ):
	delegate = PositionDelegate()
	responseData = delegate.removeDirectReports( positionId, DirectReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAssignments( request, positionId, AssignmentsIds ):
	delegate = PositionDelegate()
	responseData = delegate.addAssignments( positionId, AssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAssignments( request, positionId, AssignmentsIds ):
	delegate = PositionDelegate()
	responseData = delegate.removeAssignments( positionId, AssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

