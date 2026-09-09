from django.db import models

#======================================================================
# 
# Encapsulates data for model Laboratory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Laboratory Declaration
#======================================================================
class Laboratory (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	cliaNumber = models.CharField(max_length=200, null=True)
	facility = models.ForeignKey('Facility', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	laboratoryOrders = models.ManyToManyField('LaboratoryOrder',  blank=True, related_name='+')
	labResults = models.ManyToManyField('LabResult',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.cliaNumber
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Laboratory";
    
	def objectType(self):
		return "Laboratory";
