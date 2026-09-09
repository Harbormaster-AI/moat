from django.db import models
from governanceOnDjango.models.DocumentStatus import DocumentStatus

#======================================================================
# 
# Encapsulates data for model PrivacyNotice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PrivacyNotice Declaration
#======================================================================
class PrivacyNotice (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	audience = models.CharField(max_length=200, null=True)
	versionLabel = models.CharField(max_length=200, null=True)
	publicationDate = models.DateField(null=True)
	publicationUrl = URL
	processingActivities = models.ManyToManyField('DataProcessingActivity',  blank=True, related_name='+')
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	consents = models.ManyToManyField('Consent',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DocumentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.audience
		str = str + self.versionLabel
		str = str + self.publicationDate
		str = str + self.publicationUrl
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PrivacyNotice";
    
	def objectType(self):
		return "PrivacyNotice";
