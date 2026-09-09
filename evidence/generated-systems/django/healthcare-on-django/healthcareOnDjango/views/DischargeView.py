import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.DischargeDelegate import DischargeDelegate

 #======================================================================
# 
# Encapsulates data for View Discharge
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DischargeView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Discharge index.")

def get(request, dischargeId ):
	delegate = DischargeDelegate()
	responseData = delegate.get( dischargeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	discharge = json.loads(request.body)
	delegate = DischargeDelegate()
	responseData = delegate.createFromJson( discharge )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	discharge = json.loads(request.body)
	delegate = DischargeDelegate()
	responseData = delegate.save( discharge )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dischargeId ):
	delegate = DischargeDelegate()
	responseData = delegate.delete( dischargeId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DischargeDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEncounter( request, dischargeId, EncounterId ):
	delegate = DischargeDelegate()
	responseData = delegate.saveEncounter( dischargeId, EncounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEncounter( request, dischargeId ):
	delegate = DischargeDelegate()
	responseData = delegate.deleteEncounter( dischargeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

