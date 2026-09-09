from django.db import models
from fintechOnDjango.models.DirectDebitScheme import DirectDebitScheme
from fintechOnDjango.models.MandateStatus import MandateStatus

#======================================================================
# 
# Encapsulates data for model DirectDebitMandate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DirectDebitMandate Declaration
#======================================================================
class DirectDebitMandate (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	mandateId = models.CharField(max_length=200, null=True)
	signedAt = DateTime
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	creditor = models.ForeignKey('Creditor', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	scheme = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DirectDebitScheme])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in MandateStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.mandateId
		str = str + self.signedAt
		str = str + self.scheme
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DirectDebitMandate";
    
	def objectType(self):
		return "DirectDebitMandate";
