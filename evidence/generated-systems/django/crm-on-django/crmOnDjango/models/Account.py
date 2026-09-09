from django.db import models
from crmOnDjango.models.AccountType import AccountType
from crmOnDjango.models.AccountLifecycleStage import AccountLifecycleStage

#======================================================================
# 
# Encapsulates data for model Account
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Account Declaration
#======================================================================
class Account (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	accountNumber = models.CharField(max_length=200, null=True)
	industry = models.CharField(max_length=200, null=True)
	billingAddress = Address
	shippingAddress = Address
	website = URL
	phone = PhoneNumber
	asActive = models.BooleanField(null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	parentAccount = models.ForeignKey('self', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	childAccounts = models.ManyToManyField('Account',  blank=True, related_name='+')
	contacts = models.ManyToManyField('Contact',  blank=True, related_name='+')
	opportunities = models.ManyToManyField('Opportunity',  blank=True, related_name='+')
	cases = models.ManyToManyField('Case_',  blank=True, related_name='+')
	owner = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	territory = models.ForeignKey('Territory', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	activities = models.ManyToManyField('Activity',  blank=True, related_name='+')
	campaigns = models.ManyToManyField('Campaign',  blank=True, related_name='+')
	quotes = models.ManyToManyField('Quote',  blank=True, related_name='+')
	orders = models.ManyToManyField('Order',  blank=True, related_name='+')
	contracts = models.ManyToManyField('Contract',  blank=True, related_name='+')
	notes = models.ManyToManyField('Note',  blank=True, related_name='+')
	emailMessages = models.ManyToManyField('EmailMessage',  blank=True, related_name='+')
	accountType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AccountType])
	lifecycleStage = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AccountLifecycleStage])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.accountNumber
		str = str + self.industry
		str = str + self.billingAddress
		str = str + self.shippingAddress
		str = str + self.website
		str = str + self.phone
		str = str + self.asActive
		str = str + self.accountType
		str = str + self.lifecycleStage
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Account";
    
	def objectType(self):
		return "Account";
