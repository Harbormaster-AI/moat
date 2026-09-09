from django.db import models
from crmOnDjango.models.TeamType import TeamType

#======================================================================
# 
# Encapsulates data for model Team
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Team Declaration
#======================================================================
class Team (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	users = models.ManyToManyField('User',  blank=True, related_name='+')
	accounts = models.ManyToManyField('Account',  blank=True, related_name='+')
	opportunities = models.ManyToManyField('Opportunity',  blank=True, related_name='+')
	cases = models.ManyToManyField('Case_',  blank=True, related_name='+')
	campaigns = models.ManyToManyField('Campaign',  blank=True, related_name='+')
	teamType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TeamType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.teamType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Team";
    
	def objectType(self):
		return "Team";
