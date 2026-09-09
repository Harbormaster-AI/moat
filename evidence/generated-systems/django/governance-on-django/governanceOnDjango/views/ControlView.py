import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

 #======================================================================
# 
# Encapsulates data for View Control
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ControlView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Control index.")

def get(request, controlId ):
	delegate = ControlDelegate()
	responseData = delegate.get( controlId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	control = json.loads(request.body)
	delegate = ControlDelegate()
	responseData = delegate.createFromJson( control )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	control = json.loads(request.body)
	delegate = ControlDelegate()
	responseData = delegate.save( control )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, controlId ):
	delegate = ControlDelegate()
	responseData = delegate.delete( controlId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ControlDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, controlId, PolicyId ):
	delegate = ControlDelegate()
	responseData = delegate.savePolicy( controlId, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, controlId ):
	delegate = ControlDelegate()
	responseData = delegate.deletePolicy( controlId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addControlTests( request, controlId, ControlTestsIds ):
	delegate = ControlDelegate()
	responseData = delegate.addControlTests( controlId, ControlTestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeControlTests( request, controlId, ControlTestsIds ):
	delegate = ControlDelegate()
	responseData = delegate.removeControlTests( controlId, ControlTestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEvidence( request, controlId, EvidenceIds ):
	delegate = ControlDelegate()
	responseData = delegate.addEvidence( controlId, EvidenceIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEvidence( request, controlId, EvidenceIds ):
	delegate = ControlDelegate()
	responseData = delegate.removeEvidence( controlId, EvidenceIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRisks( request, controlId, RisksIds ):
	delegate = ControlDelegate()
	responseData = delegate.addRisks( controlId, RisksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRisks( request, controlId, RisksIds ):
	delegate = ControlDelegate()
	responseData = delegate.removeRisks( controlId, RisksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addObligations( request, controlId, ObligationsIds ):
	delegate = ControlDelegate()
	responseData = delegate.addObligations( controlId, ObligationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeObligations( request, controlId, ObligationsIds ):
	delegate = ControlDelegate()
	responseData = delegate.removeObligations( controlId, ObligationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcedures( request, controlId, ProceduresIds ):
	delegate = ControlDelegate()
	responseData = delegate.addProcedures( controlId, ProceduresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcedures( request, controlId, ProceduresIds ):
	delegate = ControlDelegate()
	responseData = delegate.removeProcedures( controlId, ProceduresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addIssues( request, controlId, IssuesIds ):
	delegate = ControlDelegate()
	responseData = delegate.addIssues( controlId, IssuesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeIssues( request, controlId, IssuesIds ):
	delegate = ControlDelegate()
	responseData = delegate.removeIssues( controlId, IssuesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

