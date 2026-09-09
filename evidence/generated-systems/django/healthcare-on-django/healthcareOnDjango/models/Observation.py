from django.db import models
from healthcareOnDjango.models.ObservationInterpretation import ObservationInterpretation

#======================================================================
# 
# Encapsulates data for model Observation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Observation Declaration
#======================================================================
class Observation (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	code = models.CharField(max_length=200, null=True)
	value = models.CharField(max_length=200, null=True)
	unit = models.CharField(max_length=200, null=True)
	effectiveDateTime = models.CharField(max_length=64, null=True)
	encounter = models.ForeignKey('Encounter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	device = models.ForeignKey('MedicalDevice', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	labResult = models.ForeignKey('LabResult', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	interpretation = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ObservationInterpretation])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.code
		str = str + self.value
		str = str + self.unit
		str = str + self.effectiveDateTime
		str = str + self.interpretation
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Observation";
    
	def objectType(self):
		return "Observation";
