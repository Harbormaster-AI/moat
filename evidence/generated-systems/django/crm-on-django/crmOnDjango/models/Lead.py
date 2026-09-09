from django.db import models
from crmOnDjango.models.LeadStatus import LeadStatus
from crmOnDjango.models.LeadSource import LeadSource
from crmOnDjango.models.LeadRating import LeadRating

#======================================================================
# 
# Encapsulates data for model Lead
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Lead Declaration
#======================================================================
class Lead (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	company = models.CharField(max_length=200, null=True)
	email = EmailAddress
	phone = PhoneNumber
	converted = models.BooleanField(null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owner = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	activities = models.ManyToManyField('Activity',  blank=True, related_name='+')
	campaigns = models.ManyToManyField('Campaign',  blank=True, related_name='+')
	convertedAccount = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	convertedContact = models.ForeignKey('Contact', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	convertedOpportunity = models.ForeignKey('Opportunity', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	notes = models.ManyToManyField('Note',  blank=True, related_name='+')
	emailMessages = models.ManyToManyField('EmailMessage',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LeadStatus])
	source = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LeadSource])
	rating = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LeadRating])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.company
		str = str + self.email
		str = str + self.phone
		str = str + self.converted
		str = str + self.status
		str = str + self.source
		str = str + self.rating
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Lead";
    
	def objectType(self):
		return "Lead";
