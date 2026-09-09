from django.db import models
from healthcareOnDjango.models.AdmissionType import AdmissionType

#======================================================================
# 
# Encapsulates data for model Admission
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Admission Declaration
#======================================================================
class Admission (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	admitDateTime = models.CharField(max_length=64, null=True)
	bed = models.CharField(max_length=200, null=True)
	encounter = models.OneToOneField('Encounter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	facility = models.ForeignKey('Facility', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	admissionType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AdmissionType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.admitDateTime
		str = str + self.bed
		str = str + self.admissionType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Admission";
    
	def objectType(self):
		return "Admission";
