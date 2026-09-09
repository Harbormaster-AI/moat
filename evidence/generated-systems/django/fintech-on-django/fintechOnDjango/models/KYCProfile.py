from django.db import models
from fintechOnDjango.models.KYCStatus import KYCStatus
from fintechOnDjango.models.VerificationLevel import VerificationLevel

#======================================================================
# 
# Encapsulates data for model KYCProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KYCProfile Declaration
#======================================================================
class KYCProfile (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	profileId = models.CharField(max_length=200, null=True)
	createdAt = DateTime
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	documents = models.ManyToManyField('KYCDocument',  blank=True, related_name='+')
	screenings = models.ManyToManyField('Screening',  blank=True, related_name='+')
	addresses = models.ManyToManyField('VerifiedAddress',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in KYCStatus])
	verificationLevel = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in VerificationLevel])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.profileId
		str = str + self.createdAt
		str = str + self.status
		str = str + self.verificationLevel
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "KYCProfile";
    
	def objectType(self):
		return "KYCProfile";
