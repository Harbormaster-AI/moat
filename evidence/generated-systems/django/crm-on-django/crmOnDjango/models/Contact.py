from django.db import models
from crmOnDjango.models.ContactMethod import ContactMethod

#======================================================================
# 
# Encapsulates data for model Contact
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Contact Declaration
#======================================================================
class Contact (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	title = models.CharField(max_length=200, null=True)
	email = EmailAddress
	phone = PhoneNumber
	mobile = PhoneNumber
	mailingAddress = Address
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owner = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	activities = models.ManyToManyField('Activity',  blank=True, related_name='+')
	opportunities = models.ManyToManyField('Opportunity',  blank=True, related_name='+')
	cases = models.ManyToManyField('Case_',  blank=True, related_name='+')
	campaigns = models.ManyToManyField('Campaign',  blank=True, related_name='+')
	notes = models.ManyToManyField('Note',  blank=True, related_name='+')
	emailMessages = models.ManyToManyField('EmailMessage',  blank=True, related_name='+')
	preferredContactMethod = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ContactMethod])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.title
		str = str + self.email
		str = str + self.phone
		str = str + self.mobile
		str = str + self.mailingAddress
		str = str + self.preferredContactMethod
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Contact";
    
	def objectType(self):
		return "Contact";
