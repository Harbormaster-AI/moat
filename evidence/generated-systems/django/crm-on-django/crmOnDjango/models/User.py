from django.db import models
from crmOnDjango.models.UserRole import UserRole
from crmOnDjango.models.UserStatus import UserStatus

#======================================================================
# 
# Encapsulates data for model User
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class User Declaration
#======================================================================
class User (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	username = models.CharField(max_length=200, null=True)
	fullName = models.CharField(max_length=200, null=True)
	email = EmailAddress
	locale = _Locale
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	teams = models.ManyToManyField('Team',  blank=True, related_name='+')
	activities = models.ManyToManyField('Activity',  blank=True, related_name='+')
	ownedAccounts = models.ManyToManyField('Account',  blank=True, related_name='+')
	ownedLeads = models.ManyToManyField('Lead',  blank=True, related_name='+')
	ownedOpportunities = models.ManyToManyField('Opportunity',  blank=True, related_name='+')
	ownedCases = models.ManyToManyField('Case_',  blank=True, related_name='+')
	quotes = models.ManyToManyField('Quote',  blank=True, related_name='+')
	orders = models.ManyToManyField('Order',  blank=True, related_name='+')
	contracts = models.ManyToManyField('Contract',  blank=True, related_name='+')
	emailMessages = models.ManyToManyField('EmailMessage',  blank=True, related_name='+')
	role = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in UserRole])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in UserStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.username
		str = str + self.fullName
		str = str + self.email
		str = str + self.locale
		str = str + self.role
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "User";
    
	def objectType(self):
		return "User";
