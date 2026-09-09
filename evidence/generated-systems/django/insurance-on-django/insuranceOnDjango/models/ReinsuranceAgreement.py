from django.db import models
from insuranceOnDjango.models.ReinsuranceType import ReinsuranceType
from insuranceOnDjango.models.TreatyType import TreatyType

#======================================================================
# 
# Encapsulates data for model ReinsuranceAgreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReinsuranceAgreement Declaration
#======================================================================
class ReinsuranceAgreement (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	agreementNumber = models.CharField(max_length=200, null=True)
	effectivePeriod = DateRange
	retention = Money
	limit = Money
	cessionPercentage = Percentage
	insurer = models.ForeignKey('Insurer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	policies = models.ManyToManyField('Policy',  blank=True, related_name='+')
	reinsuranceType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReinsuranceType])
	treatyType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TreatyType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.agreementNumber
		str = str + self.effectivePeriod
		str = str + self.retention
		str = str + self.limit
		str = str + self.cessionPercentage
		str = str + self.reinsuranceType
		str = str + self.treatyType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ReinsuranceAgreement";
    
	def objectType(self):
		return "ReinsuranceAgreement";
