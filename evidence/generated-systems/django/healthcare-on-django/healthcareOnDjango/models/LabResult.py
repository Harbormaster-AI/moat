from django.db import models
from healthcareOnDjango.models.ResultStatus import ResultStatus

#======================================================================
# 
# Encapsulates data for model LabResult
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LabResult Declaration
#======================================================================
class LabResult (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	resultCode = models.CharField(max_length=200, null=True)
	issuedDate = models.CharField(max_length=64, null=True)
	laboratoryOrder = models.ForeignKey('LaboratoryOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	observations = models.ManyToManyField('Observation',  blank=True, related_name='+')
	laboratory = models.ForeignKey('Laboratory', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ResultStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.resultCode
		str = str + self.issuedDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "LabResult";
    
	def objectType(self):
		return "LabResult";
