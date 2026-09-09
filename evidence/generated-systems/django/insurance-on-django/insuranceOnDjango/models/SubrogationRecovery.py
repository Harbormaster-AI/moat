from django.db import models
from insuranceOnDjango.models.SubrogationStatus import SubrogationStatus

#======================================================================
# 
# Encapsulates data for model SubrogationRecovery
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubrogationRecovery Declaration
#======================================================================
class SubrogationRecovery (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	recoveryReference = models.CharField(max_length=200, null=True)
	amount = Money
	recoveryDate = models.DateField(null=True)
	claim = models.ForeignKey('Claim', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	exposure = models.ForeignKey('Exposure', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	counterparty = models.ForeignKey('ThirdParty', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SubrogationStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.recoveryReference
		str = str + self.amount
		str = str + self.recoveryDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "SubrogationRecovery";
    
	def objectType(self):
		return "SubrogationRecovery";
