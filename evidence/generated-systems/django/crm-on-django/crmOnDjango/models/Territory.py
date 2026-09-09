from django.db import models
from crmOnDjango.models.TerritoryType import TerritoryType

#======================================================================
# 
# Encapsulates data for model Territory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Territory Declaration
#======================================================================
class Territory (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	region = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	accounts = models.ManyToManyField('Account',  blank=True, related_name='+')
	users = models.ManyToManyField('User',  blank=True, related_name='+')
	territoryType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TerritoryType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.region
		str = str + self.territoryType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Territory";
    
	def objectType(self):
		return "Territory";
