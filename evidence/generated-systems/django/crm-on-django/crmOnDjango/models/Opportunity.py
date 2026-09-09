from django.db import models
from crmOnDjango.models.OpportunityStage import OpportunityStage
from crmOnDjango.models.OpportunityType import OpportunityType
from crmOnDjango.models.ForecastCategory import ForecastCategory

#======================================================================
# 
# Encapsulates data for model Opportunity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Opportunity Declaration
#======================================================================
class Opportunity (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	amount = Money
	closeDate = models.DateField(null=True)
	probability = models.CharField(max_length=64, null=True)
	description = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owner = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	contacts = models.ManyToManyField('Contact',  blank=True, related_name='+')
	lineItems = models.ManyToManyField('OpportunityLineItem',  blank=True, related_name='+')
	stageHistory = models.ManyToManyField('OpportunityStageHistory',  blank=True, related_name='+')
	quotes = models.ManyToManyField('Quote',  blank=True, related_name='+')
	orders = models.ManyToManyField('Order',  blank=True, related_name='+')
	campaigns = models.ManyToManyField('Campaign',  blank=True, related_name='+')
	activities = models.ManyToManyField('Activity',  blank=True, related_name='+')
	teams = models.ManyToManyField('Team',  blank=True, related_name='+')
	stage = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OpportunityStage])
	type = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OpportunityType])
	forecastCategory = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ForecastCategory])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.amount
		str = str + self.closeDate
		str = str + self.probability
		str = str + self.description
		str = str + self.stage
		str = str + self.type
		str = str + self.forecastCategory
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Opportunity";
    
	def objectType(self):
		return "Opportunity";
