from django.db import models
from governanceOnDjango.models.LawfulBasis import LawfulBasis

#======================================================================
# 
# Encapsulates data for model DataProcessingActivity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataProcessingActivity Declaration
#======================================================================
class DataProcessingActivity (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	purpose = models.CharField(max_length=200, null=True)
	startDate = models.DateField(null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dataCategories = models.ManyToManyField('DataCategory',  blank=True, related_name='+')
	systems = models.ManyToManyField('System_',  blank=True, related_name='+')
	records = models.ManyToManyField('Record_',  blank=True, related_name='+')
	privacyNotices = models.ManyToManyField('PrivacyNotice',  blank=True, related_name='+')
	thirdParties = models.ManyToManyField('ThirdParty',  blank=True, related_name='+')
	consents = models.ManyToManyField('Consent',  blank=True, related_name='+')
	dataBreaches = models.ManyToManyField('DataBreach',  blank=True, related_name='+')
	dataSubjectRequests = models.ManyToManyField('DataSubjectRequest',  blank=True, related_name='+')
	lawfulBasis = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LawfulBasis])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.purpose
		str = str + self.startDate
		str = str + self.lawfulBasis
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DataProcessingActivity";
    
	def objectType(self):
		return "DataProcessingActivity";
