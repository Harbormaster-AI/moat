from django.db import models
from healthcareOnDjango.models.DischargeDisposition import DischargeDisposition

#======================================================================
# 
# Encapsulates data for model Discharge
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Discharge Declaration
#======================================================================
class Discharge (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	dischargeDateTime = models.CharField(max_length=64, null=True)
	encounter = models.OneToOneField('Encounter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	disposition = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DischargeDisposition])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.dischargeDateTime
		str = str + self.disposition
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Discharge";
    
	def objectType(self):
		return "Discharge";
