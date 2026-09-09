from django.db import models
from fintechOnDjango.models.VerificationStatus import VerificationStatus

#======================================================================
# 
# Encapsulates data for model VerifiedAddress
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class VerifiedAddress Declaration
#======================================================================
class VerifiedAddress (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	address = Address
	verifiedAt = DateTime
	kycProfile = models.ForeignKey('KYCProfile', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	verificationStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in VerificationStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.address
		str = str + self.verifiedAt
		str = str + self.verificationStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "VerifiedAddress";
    
	def objectType(self):
		return "VerifiedAddress";
