from django.db import models
from aerospaceOnDjango.models.OperatorType import OperatorType

#======================================================================
# 
# Encapsulates data for model Operator
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Operator Declaration
#======================================================================
class Operator (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	icaoDesignator = models.CharField(max_length=200, null=True)
	aircraftOrders = models.ManyToManyField('AircraftOrder',  blank=True, related_name='+')
	operatedAircraft = models.ManyToManyField('Aircraft',  blank=True, related_name='+')
	salesRegion = models.ForeignKey('SalesRegion', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	operatorType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OperatorType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.icaoDesignator
		str = str + self.operatorType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Operator";
    
	def objectType(self):
		return "Operator";
