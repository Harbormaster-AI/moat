from django.db import models
from insuranceOnDjango.models.RelationshipType import RelationshipType

#======================================================================
# 
# Encapsulates data for model Beneficiary
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Beneficiary Declaration
#======================================================================
class Beneficiary (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	share = Percentage
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	relationship = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RelationshipType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.share
		str = str + self.relationship
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Beneficiary";
    
	def objectType(self):
		return "Beneficiary";
