from django.db import models
from healthcareOnDjango.models.AnesthesiaType import AnesthesiaType

#======================================================================
# 
# Encapsulates data for model ProcedureOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProcedureOrder Declaration
#======================================================================
class ProcedureOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	procedureCode = models.CharField(max_length=200, null=True)
	consentObtained = models.BooleanField(null=True)
	order = models.ForeignKey('ClinicalOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	facility = models.ForeignKey('Facility', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	procedure = models.OneToOneField('Procedure', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	anesthesiaType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AnesthesiaType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.procedureCode
		str = str + self.consentObtained
		str = str + self.anesthesiaType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ProcedureOrder";
    
	def objectType(self):
		return "ProcedureOrder";
