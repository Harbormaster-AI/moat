from django.db import models
from advertisingOnDjango.models.AccountRole import AccountRole

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
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	email = Email
	agency = models.ForeignKey('Agency', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	teams = models.ManyToManyField('Team',  blank=True, related_name='+')
	adAccounts = models.ManyToManyField('AdAccount',  blank=True, related_name='+')
	role = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AccountRole])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.email
		str = str + self.role
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "User";
    
	def objectType(self):
		return "User";
