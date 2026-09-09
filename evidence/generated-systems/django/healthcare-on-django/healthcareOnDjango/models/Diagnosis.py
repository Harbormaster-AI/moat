from django.db import models
from healthcareOnDjango.models.DiagnosisCertainty import DiagnosisCertainty

#======================================================================
# 
# Encapsulates data for model Diagnosis
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Diagnosis Declaration
#======================================================================
class Diagnosis (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	code = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	onsetDate = models.DateField(null=True)
	encounter = models.ForeignKey('Encounter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	certainty = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DiagnosisCertainty])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.code
		str = str + self.description
		str = str + self.onsetDate
		str = str + self.certainty
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Diagnosis";
    
	def objectType(self):
		return "Diagnosis";
