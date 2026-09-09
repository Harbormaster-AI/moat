from django.db import models

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
	agency = models.ForeignKey('Agency', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	users = models.ManyToManyField('User',  blank=True, related_name='+')
	adAccounts = models.ManyToManyField('AdAccount',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Team";
    
	def objectType(self):
		return "Team";
