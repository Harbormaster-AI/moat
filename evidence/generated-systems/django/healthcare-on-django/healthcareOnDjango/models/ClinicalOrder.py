from django.db import models
from healthcareOnDjango.models.OrderStatus import OrderStatus
from healthcareOnDjango.models.ClinicalOrderType import ClinicalOrderType
from healthcareOnDjango.models.Priority import Priority

#======================================================================
# 
# Encapsulates data for model ClinicalOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClinicalOrder Declaration
#======================================================================
class ClinicalOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	orderNumber = models.CharField(max_length=200, null=True)
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	encounter = models.ForeignKey('Encounter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	orderingClinician = models.ForeignKey('Clinician', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	medicationOrders = models.ManyToManyField('MedicationOrder',  blank=True, related_name='+')
	laboratoryOrders = models.ManyToManyField('LaboratoryOrder',  blank=True, related_name='+')
	imagingOrders = models.ManyToManyField('ImagingOrder',  blank=True, related_name='+')
	procedureOrders = models.ManyToManyField('ProcedureOrder',  blank=True, related_name='+')
	authorizations = models.ManyToManyField('Authorization',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OrderStatus])
	orderType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ClinicalOrderType])
	priority = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in Priority])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.orderNumber
		str = str + self.status
		str = str + self.orderType
		str = str + self.priority
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ClinicalOrder";
    
	def objectType(self):
		return "ClinicalOrder";
