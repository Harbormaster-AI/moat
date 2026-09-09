from django.db import models

#======================================================================
# 
# Encapsulates data for model AvionicsSuite
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AvionicsSuite Declaration
#======================================================================
class AvionicsSuite (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	suiteName = models.CharField(max_length=200, null=True)
	softwareBaseline = models.CharField(max_length=200, null=True)
	supplier = models.ForeignKey('Supplier', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	variants = models.ManyToManyField('AircraftVariant',  blank=True, related_name='+')
	softwareLoads = models.ManyToManyField('SoftwareLoad',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.suiteName
		str = str + self.softwareBaseline
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AvionicsSuite";
    
	def objectType(self):
		return "AvionicsSuite";
