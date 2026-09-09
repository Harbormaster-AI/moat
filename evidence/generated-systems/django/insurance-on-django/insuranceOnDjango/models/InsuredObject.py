from django.db import models
from insuranceOnDjango.models.InsuredObjectType import InsuredObjectType

#======================================================================
# 
# Encapsulates data for model InsuredObject
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsuredObject Declaration
#======================================================================
class InsuredObject (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	description = models.CharField(max_length=200, null=True)
	serialOrId = models.CharField(max_length=200, null=True)
	primaryAddress = Address
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	coverages = models.ManyToManyField('PolicyCoverage',  blank=True, related_name='+')
	objectType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InsuredObjectType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.description
		str = str + self.serialOrId
		str = str + self.primaryAddress
		str = str + self.objectType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InsuredObject";
    
	def objectType(self):
		return "InsuredObject";
