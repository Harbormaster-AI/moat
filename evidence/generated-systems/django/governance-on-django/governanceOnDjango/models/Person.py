from django.db import models

#======================================================================
# 
# Encapsulates data for model Person
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Person Declaration
#======================================================================
class Person (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	email = EmailAddress
	department = models.CharField(max_length=200, null=True)
	roleAssignments = models.ManyToManyField('RoleAssignment',  blank=True, related_name='+')
	ownedPolicies = models.ManyToManyField('Policy',  blank=True, related_name='+')
	correctiveActions = models.ManyToManyField('CorrectiveAction',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.email
		str = str + self.department
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Person";
    
	def objectType(self):
		return "Person";
