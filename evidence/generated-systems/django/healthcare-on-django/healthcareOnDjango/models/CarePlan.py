from django.db import models
from healthcareOnDjango.models.CarePlanStatus import CarePlanStatus

#======================================================================
# 
# Encapsulates data for model CarePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CarePlan Declaration
#======================================================================
class CarePlan (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	planNumber = models.CharField(max_length=200, null=True)
	goalSummary = models.CharField(max_length=200, null=True)
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	encounters = models.ManyToManyField('Encounter',  blank=True, related_name='+')
	tasks = models.ManyToManyField('CareTask',  blank=True, related_name='+')
	careTeam = models.ForeignKey('CareTeam', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CarePlanStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.planNumber
		str = str + self.goalSummary
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CarePlan";
    
	def objectType(self):
		return "CarePlan";
