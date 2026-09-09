from django.db import models

#======================================================================
# 
# Encapsulates data for model APU
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class APU Declaration
#======================================================================
class APU (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	model = models.CharField(max_length=200, null=True)
	supplier = models.ForeignKey('Supplier', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	variants = models.ManyToManyField('AircraftVariant',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.model
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "APU";
    
	def objectType(self):
		return "APU";
