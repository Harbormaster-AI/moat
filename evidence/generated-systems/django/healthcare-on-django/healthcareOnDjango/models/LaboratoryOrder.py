from django.db import models
from healthcareOnDjango.models.SpecimenType import SpecimenType

#======================================================================
# 
# Encapsulates data for model LaboratoryOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LaboratoryOrder Declaration
#======================================================================
class LaboratoryOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	testCode = models.CharField(max_length=200, null=True)
	fastingRequired = models.BooleanField(null=True)
	order = models.ForeignKey('ClinicalOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	laboratory = models.ForeignKey('Laboratory', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	results = models.ManyToManyField('LabResult',  blank=True, related_name='+')
	specimenType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SpecimenType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.testCode
		str = str + self.fastingRequired
		str = str + self.specimenType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "LaboratoryOrder";
    
	def objectType(self):
		return "LaboratoryOrder";
