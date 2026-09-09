from django.db import models
from healthcareOnDjango.models.ConditionStatus import ConditionStatus
from healthcareOnDjango.models.DiagnosisCertainty import DiagnosisCertainty

#======================================================================
# 
# Encapsulates data for model Condition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Condition Declaration
#======================================================================
class Condition (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	code = models.CharField(max_length=200, null=True)
	onsetDate = models.DateField(null=True)
	abatementDate = models.DateField(null=True)
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	clinicalStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ConditionStatus])
	verificationStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DiagnosisCertainty])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.code
		str = str + self.onsetDate
		str = str + self.abatementDate
		str = str + self.clinicalStatus
		str = str + self.verificationStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Condition";
    
	def objectType(self):
		return "Condition";
