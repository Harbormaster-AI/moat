from django.db import models
from fintechOnDjango.models.CollateralType import CollateralType

#======================================================================
# 
# Encapsulates data for model Collateral
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Collateral Declaration
#======================================================================
class Collateral (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	description = models.CharField(max_length=200, null=True)
	value = Money
	loan = models.ForeignKey('Loan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	collateralType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CollateralType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.description
		str = str + self.value
		str = str + self.collateralType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Collateral";
    
	def objectType(self):
		return "Collateral";
