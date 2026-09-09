from django.db import models
from governanceOnDjango.models.ObligationType import ObligationType
from governanceOnDjango.models.ControlFrequency import ControlFrequency

#======================================================================
# 
# Encapsulates data for model Obligation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Obligation Declaration
#======================================================================
class Obligation (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	referenceNumber = models.CharField(max_length=200, null=True)
	descriptionText = models.CharField(max_length=200, null=True)
	regulation = models.ForeignKey('Regulation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	controls = models.ManyToManyField('Control',  blank=True, related_name='+')
	policies = models.ManyToManyField('Policy',  blank=True, related_name='+')
	contracts = models.ManyToManyField('Contract',  blank=True, related_name='+')
	obligationType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ObligationType])
	reviewFrequency = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ControlFrequency])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.referenceNumber
		str = str + self.descriptionText
		str = str + self.obligationType
		str = str + self.reviewFrequency
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Obligation";
    
	def objectType(self):
		return "Obligation";
