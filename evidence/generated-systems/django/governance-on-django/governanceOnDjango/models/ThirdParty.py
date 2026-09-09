from django.db import models
from governanceOnDjango.models.ThirdPartyType import ThirdPartyType
from governanceOnDjango.models.VendorCriticality import VendorCriticality

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
	country = models.CharField(max_length=200, null=True)
	contactEmail = EmailAddress
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	processingActivities = models.ManyToManyField('DataProcessingActivity',  blank=True, related_name='+')
	assessments = models.ManyToManyField('ThirdPartyAssessment',  blank=True, related_name='+')
	contracts = models.ManyToManyField('Contract',  blank=True, related_name='+')
	obligations = models.ManyToManyField('Obligation',  blank=True, related_name='+')
	dataBreaches = models.ManyToManyField('DataBreach',  blank=True, related_name='+')
	thirdPartyType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ThirdPartyType])
	criticality = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in VendorCriticality])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.country
		str = str + self.contactEmail
		str = str + self.thirdPartyType
		str = str + self.criticality
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ThirdParty";
    
	def objectType(self):
		return "ThirdParty";
