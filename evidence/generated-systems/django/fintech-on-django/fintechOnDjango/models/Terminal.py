from django.db import models
from fintechOnDjango.models.TerminalType import TerminalType
from fintechOnDjango.models.TerminalStatus import TerminalStatus

#======================================================================
# 
# Encapsulates data for model Terminal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Terminal Declaration
#======================================================================
class Terminal (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	location = Address
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	type = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TerminalType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TerminalStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.location
		str = str + self.type
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Terminal";
    
	def objectType(self):
		return "Terminal";
