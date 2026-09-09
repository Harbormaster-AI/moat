from django.db import models
from governanceOnDjango.models.ConsentType import ConsentType
from governanceOnDjango.models.ConsentStatus import ConsentStatus

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
	subjectIdentifier = models.CharField(max_length=200, null=True)
	captureDate = models.DateField(null=True)
	expiryDate = models.DateField(null=True)
	processingActivities = models.ManyToManyField('DataProcessingActivity',  blank=True, related_name='+')
	privacyNotice = models.ForeignKey('PrivacyNotice', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	consentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ConsentType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ConsentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.subjectIdentifier
		str = str + self.captureDate
		str = str + self.expiryDate
		str = str + self.consentType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Consent";
    
	def objectType(self):
		return "Consent";
