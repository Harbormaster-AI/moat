from django.db import models

#======================================================================
# 
# Encapsulates data for model Role
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Role Declaration
#======================================================================
class Role (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	responsibility = models.CharField(max_length=200, null=True)
	assignments = models.ManyToManyField('RoleAssignment',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.responsibility
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Role";
    
	def objectType(self):
		return "Role";
