from django.db import models
from healthcareOnDjango.models.ProcedureStatus import ProcedureStatus

#======================================================================
# 
# Encapsulates data for model Procedure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Procedure Declaration
#======================================================================
class Procedure (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	procedureCode = models.CharField(max_length=200, null=True)
	startDateTime = models.CharField(max_length=64, null=True)
	endDateTime = models.CharField(max_length=64, null=True)
	encounter = models.ForeignKey('Encounter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	performer = models.ForeignKey('Clinician', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	procedureOrder = models.ForeignKey('ProcedureOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ProcedureStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.procedureCode
		str = str + self.startDateTime
		str = str + self.endDateTime
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Procedure";
    
	def objectType(self):
		return "Procedure";
