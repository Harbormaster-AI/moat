from django.db import models
from aerospaceOnDjango.models.ComponentCategory import ComponentCategory
from aerospaceOnDjango.models.SerializationMethod import SerializationMethod

#======================================================================
# 
# Encapsulates data for model Component_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Component_ Declaration
#======================================================================
class Component_ (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	partNumber = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	supplier = models.ForeignKey('Supplier', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	componentCategory = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ComponentCategory])
	serializationMethod = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SerializationMethod])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.partNumber
		str = str + self.name
		str = str + self.componentCategory
		str = str + self.serializationMethod
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Component_";
    
	def objectType(self):
		return "Component_";
