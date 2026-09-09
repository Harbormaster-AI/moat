from django.db import models
from manufacturingOnDjango.models.ShiftType import ShiftType

#======================================================================
# 
# Encapsulates data for model Shift
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Shift Declaration
#======================================================================
class Shift (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	shiftName = models.CharField(max_length=200, null=True)
	startTime = models.CharField(max_length=200, null=True)
	endTime = models.CharField(max_length=200, null=True)
	plant = models.ForeignKey('Plant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	assignments = models.ManyToManyField('ShiftAssignment',  blank=True, related_name='+')
	shiftType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ShiftType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.shiftName
		str = str + self.startTime
		str = str + self.endTime
		str = str + self.shiftType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Shift";
    
	def objectType(self):
		return "Shift";
