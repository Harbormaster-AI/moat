from django.db import models

#======================================================================
# 
# Encapsulates data for model Note
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Note Declaration
#======================================================================
class Note (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	content = models.CharField(max_length=200, null=True)
	createdAt = models.CharField(max_length=64, null=True)
	updatedAt = models.CharField(max_length=64, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owner = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	contact = models.ForeignKey('Contact', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	opportunity = models.ForeignKey('Opportunity', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	case = models.ForeignKey('Case_', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lead = models.ForeignKey('Lead', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.content
		str = str + self.createdAt
		str = str + self.updatedAt
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Note";
    
	def objectType(self):
		return "Note";
