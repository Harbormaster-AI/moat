from django.db import models
from governanceOnDjango.models.DataSubjectRequestType import DataSubjectRequestType
from governanceOnDjango.models.RequestStatus import RequestStatus

#======================================================================
# 
# Encapsulates data for model DataSubjectRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSubjectRequest Declaration
#======================================================================
class DataSubjectRequest (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	receivedDate = models.DateField(null=True)
	dueDate = models.DateField(null=True)
	requesterCountry = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	processingActivities = models.ManyToManyField('DataProcessingActivity',  blank=True, related_name='+')
	records = models.ManyToManyField('Record_',  blank=True, related_name='+')
	requestType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DataSubjectRequestType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RequestStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.receivedDate
		str = str + self.dueDate
		str = str + self.requesterCountry
		str = str + self.requestType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DataSubjectRequest";
    
	def objectType(self):
		return "DataSubjectRequest";
