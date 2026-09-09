from django.db import models
from crmOnDjango.models.CampaignStatus import CampaignStatus
from crmOnDjango.models.CampaignType import CampaignType

#======================================================================
# 
# Encapsulates data for model Campaign
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Campaign Declaration
#======================================================================
class Campaign (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	startDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	budget = Money
	actualCost = Money
	expectedRevenue = Money
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	parentCampaign = models.ForeignKey('self', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	childCampaigns = models.ManyToManyField('Campaign',  blank=True, related_name='+')
	members = models.ManyToManyField('CampaignMember',  blank=True, related_name='+')
	opportunities = models.ManyToManyField('Opportunity',  blank=True, related_name='+')
	accounts = models.ManyToManyField('Account',  blank=True, related_name='+')
	leads = models.ManyToManyField('Lead',  blank=True, related_name='+')
	contacts = models.ManyToManyField('Contact',  blank=True, related_name='+')
	teams = models.ManyToManyField('Team',  blank=True, related_name='+')
	activities = models.ManyToManyField('Activity',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CampaignStatus])
	type = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CampaignType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.startDate
		str = str + self.endDate
		str = str + self.budget
		str = str + self.actualCost
		str = str + self.expectedRevenue
		str = str + self.status
		str = str + self.type
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Campaign";
    
	def objectType(self):
		return "Campaign";
