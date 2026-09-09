from django.db import models

#======================================================================
# 
# Encapsulates data for model ContentCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContentCategory Declaration
#======================================================================
class ContentCategory (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	code = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.code
		str = str + self.name
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ContentCategory";
    
	def objectType(self):
		return "ContentCategory";
