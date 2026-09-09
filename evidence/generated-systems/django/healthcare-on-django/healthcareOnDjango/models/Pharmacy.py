from django.db import models

#======================================================================
# 
# Encapsulates data for model Pharmacy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Pharmacy Declaration
#======================================================================
class Pharmacy (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	facility = models.ForeignKey('Facility', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	medicationDispenses = models.ManyToManyField('MedicationDispense',  blank=True, related_name='+')
	medicationOrders = models.ManyToManyField('MedicationOrder',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Pharmacy";
    
	def objectType(self):
		return "Pharmacy";
