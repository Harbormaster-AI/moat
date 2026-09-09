from django.db import models
from insuranceOnDjango.models.ThirdPartyType import ThirdPartyType

#======================================================================
# 
# Encapsulates data for model ThirdParty
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ThirdParty Declaration
#======================================================================
class ThirdParty (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	taxId = models.CharField(max_length=200, null=True)
	address = Address
	subrogations = models.ManyToManyField('SubrogationRecovery',  blank=True, related_name='+')
	partyType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ThirdPartyType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.taxId
		str = str + self.address
		str = str + self.partyType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ThirdParty";
    
	def objectType(self):
		return "ThirdParty";
