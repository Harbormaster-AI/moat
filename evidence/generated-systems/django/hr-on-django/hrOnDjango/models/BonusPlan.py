from django.db import models

#======================================================================
# 
# Encapsulates data for model BonusPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BonusPlan Declaration
#======================================================================
class BonusPlan (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	targetPercentage = Percentage
	compensationPackages = models.ManyToManyField('CompensationPackage',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.targetPercentage
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BonusPlan";
    
	def objectType(self):
		return "BonusPlan";
