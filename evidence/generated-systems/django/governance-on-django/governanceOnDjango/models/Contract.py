from django.db import models
from governanceOnDjango.models.ContractStatus import ContractStatus

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
	title = models.CharField(max_length=200, null=True)
	effectiveDate = models.DateField(null=True)
	expiryDate = models.DateField(null=True)
	repositoryUrl = URL
	thirdParty = models.ForeignKey('ThirdParty', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	obligations = models.ManyToManyField('Obligation',  blank=True, related_name='+')
	dataProcessingActivities = models.ManyToManyField('DataProcessingActivity',  blank=True, related_name='+')
	matter = models.ForeignKey('Matter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ContractStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.effectiveDate
		str = str + self.expiryDate
		str = str + self.repositoryUrl
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Contract";
    
	def objectType(self):
		return "Contract";
