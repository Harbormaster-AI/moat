from django.db import models
from healthcareOnDjango.models.RouteOfAdministration import RouteOfAdministration

#======================================================================
# 
# Encapsulates data for model MedicationOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicationOrder Declaration
#======================================================================
class MedicationOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	medicationCode = models.CharField(max_length=200, null=True)
	dose = Dose
	frequency = models.CharField(max_length=200, null=True)
	duration = models.CharField(max_length=200, null=True)
	order = models.ForeignKey('ClinicalOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	pharmacy = models.ForeignKey('Pharmacy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dispenses = models.ManyToManyField('MedicationDispense',  blank=True, related_name='+')
	route = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RouteOfAdministration])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.medicationCode
		str = str + self.dose
		str = str + self.frequency
		str = str + self.duration
		str = str + self.route
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "MedicationOrder";
    
	def objectType(self):
		return "MedicationOrder";
