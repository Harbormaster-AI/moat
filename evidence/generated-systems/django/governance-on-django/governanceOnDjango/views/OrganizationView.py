import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

 #======================================================================
# 
# Encapsulates data for View Organization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrganizationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Organization index.")

def get(request, organizationId ):
	delegate = OrganizationDelegate()
	responseData = delegate.get( organizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	organization = json.loads(request.body)
	delegate = OrganizationDelegate()
	responseData = delegate.createFromJson( organization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	organization = json.loads(request.body)
	delegate = OrganizationDelegate()
	responseData = delegate.save( organization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, organizationId ):
	delegate = OrganizationDelegate()
	responseData = delegate.delete( organizationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OrganizationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addGovernanceBodies( request, organizationId, GovernanceBodiesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addGovernanceBodies( organizationId, GovernanceBodiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeGovernanceBodies( request, organizationId, GovernanceBodiesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeGovernanceBodies( organizationId, GovernanceBodiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPolicies( request, organizationId, PoliciesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addPolicies( organizationId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePolicies( request, organizationId, PoliciesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removePolicies( organizationId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRisks( request, organizationId, RisksIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addRisks( organizationId, RisksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRisks( request, organizationId, RisksIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeRisks( organizationId, RisksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addThirdParties( request, organizationId, ThirdPartiesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addThirdParties( organizationId, ThirdPartiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeThirdParties( request, organizationId, ThirdPartiesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeThirdParties( organizationId, ThirdPartiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRecordsRepositories( request, organizationId, RecordsRepositoriesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addRecordsRepositories( organizationId, RecordsRepositoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRecordsRepositories( request, organizationId, RecordsRepositoriesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeRecordsRepositories( organizationId, RecordsRepositoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataProcessingActivities( request, organizationId, DataProcessingActivitiesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addDataProcessingActivities( organizationId, DataProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataProcessingActivities( request, organizationId, DataProcessingActivitiesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeDataProcessingActivities( organizationId, DataProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCompliancePrograms( request, organizationId, ComplianceProgramsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addCompliancePrograms( organizationId, ComplianceProgramsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCompliancePrograms( request, organizationId, ComplianceProgramsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeCompliancePrograms( organizationId, ComplianceProgramsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAuditPrograms( request, organizationId, AuditProgramsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addAuditPrograms( organizationId, AuditProgramsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAuditPrograms( request, organizationId, AuditProgramsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeAuditPrograms( organizationId, AuditProgramsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBusinessUnits( request, organizationId, BusinessUnitsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addBusinessUnits( organizationId, BusinessUnitsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBusinessUnits( request, organizationId, BusinessUnitsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeBusinessUnits( organizationId, BusinessUnitsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMatters( request, organizationId, MattersIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addMatters( organizationId, MattersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMatters( request, organizationId, MattersIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeMatters( organizationId, MattersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataBreaches( request, organizationId, DataBreachesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addDataBreaches( organizationId, DataBreachesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataBreaches( request, organizationId, DataBreachesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeDataBreaches( organizationId, DataBreachesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

