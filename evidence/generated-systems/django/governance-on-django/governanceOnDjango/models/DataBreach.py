from django.db import models
from governanceOnDjango.models.BreachSeverity import BreachSeverity
from governanceOnDjango.models.IncidentStatus import IncidentStatus

#======================================================================
# 
# Encapsulates data for model DataBreach
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataBreach Declaration
#======================================================================
class DataBreach (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	incidentDate = models.DateField(null=True)
	description = models.CharField(max_length=200, null=True)
	recordsAffected = models.IntegerField(null=True)
	notificationRequired = models.BooleanField(null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	processingActivities = models.ManyToManyField('DataProcessingActivity',  blank=True, related_name='+')
	dataCategories = models.ManyToManyField('DataCategory',  blank=True, related_name='+')
	thirdParties = models.ManyToManyField('ThirdParty',  blank=True, related_name='+')
	matter = models.ForeignKey('Matter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	severity = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in BreachSeverity])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in IncidentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.incidentDate
		str = str + self.description
		str = str + self.recordsAffected
		str = str + self.notificationRequired
		str = str + self.severity
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DataBreach";
    
	def objectType(self):
		return "DataBreach";
