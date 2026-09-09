from django.db import models
from crmOnDjango.models.ContractStatus import ContractStatus

#======================================================================
# 
# Encapsulates data for model Contract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Contract Declaration
#======================================================================
class Contract (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	contractNumber = models.CharField(max_length=200, null=True)
	startDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	renewalTermMonths = models.IntegerField(null=True)
	autoRenew = models.BooleanField(null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owner = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	orders = models.ManyToManyField('Order',  blank=True, related_name='+')
	cases = models.ManyToManyField('Case_',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ContractStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.contractNumber
		str = str + self.startDate
		str = str + self.endDate
		str = str + self.renewalTermMonths
		str = str + self.autoRenew
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Contract";
    
	def objectType(self):
		return "Contract";
