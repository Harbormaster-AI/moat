from django.db import models
from hrOnDjango.models.EquityType import EquityType

#======================================================================
# 
# Encapsulates data for model EquityGrant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EquityGrant Declaration
#======================================================================
class EquityGrant (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	grantId = models.CharField(max_length=200, null=True)
	grantedUnits = models.IntegerField(null=True)
	vestingStart = models.DateField(null=True)
	compensationPackage = models.ForeignKey('CompensationPackage', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	grantType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in EquityType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.grantId
		str = str + self.grantedUnits
		str = str + self.vestingStart
		str = str + self.grantType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "EquityGrant";
    
	def objectType(self):
		return "EquityGrant";
