from django.db import models
from healthcareOnDjango.models.AllergySeverity import AllergySeverity
from healthcareOnDjango.models.AllergyStatus import AllergyStatus

#======================================================================
# 
# Encapsulates data for model Allergy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Allergy Declaration
#======================================================================
class Allergy (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	substance = models.CharField(max_length=200, null=True)
	reaction = models.CharField(max_length=200, null=True)
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	severity = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AllergySeverity])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AllergyStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.substance
		str = str + self.reaction
		str = str + self.severity
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Allergy";
    
	def objectType(self):
		return "Allergy";
