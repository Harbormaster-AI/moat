from django.db import models
from fintechOnDjango.models.ConsentType import ConsentType
from fintechOnDjango.models.ConsentStatus import ConsentStatus

#======================================================================
# 
# Encapsulates data for model Consent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Consent Declaration
#======================================================================
class Consent (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	grantedAt = DateTime
	expiresAt = DateTime
	scope = models.CharField(max_length=200, null=True)
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	apiClient = models.ForeignKey('APIClient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	consentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ConsentType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ConsentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.grantedAt
		str = str + self.expiresAt
		str = str + self.scope
		str = str + self.consentType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Consent";
    
	def objectType(self):
		return "Consent";
