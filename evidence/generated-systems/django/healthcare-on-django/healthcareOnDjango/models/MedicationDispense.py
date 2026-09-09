from django.db import models
from healthcareOnDjango.models.DispenseStatus import DispenseStatus

#======================================================================
# 
# Encapsulates data for model MedicationDispense
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicationDispense Declaration
#======================================================================
class MedicationDispense (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	dispenseNumber = models.CharField(max_length=200, null=True)
	quantity = models.CharField(max_length=64, null=True)
	whenPrepared = models.CharField(max_length=64, null=True)
	medicationOrder = models.ForeignKey('MedicationOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	pharmacy = models.ForeignKey('Pharmacy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DispenseStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.dispenseNumber
		str = str + self.quantity
		str = str + self.whenPrepared
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "MedicationDispense";
    
	def objectType(self):
		return "MedicationDispense";
