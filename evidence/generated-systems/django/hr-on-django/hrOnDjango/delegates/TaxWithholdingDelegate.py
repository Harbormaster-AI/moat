from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.TaxWithholding import TaxWithholding
from hrOnDjango.models.Employee import Employee
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TaxWithholding
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TaxWithholdingDelegate Declaration
#======================================================================
class TaxWithholdingDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, taxWithholdingId ):
		try:	
			taxWithholding = TaxWithholding.objects.filter(id=taxWithholdingId)
			return taxWithholding.first();
		except TaxWithholding.DoesNotExist:
			raise ProcessingError("TaxWithholding with id " + str(taxWithholdingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, taxWithholding):
		for model in serializers.deserialize("json", taxWithholding):
			model.save()
			return model;

	def create(self, taxWithholding):
		taxWithholding.save()
		return taxWithholding;

	def saveFromJson(self, taxWithholding):
		for model in serializers.deserialize("json", taxWithholding):
			model.save()
			return taxWithholding;
	
	def save(self, taxWithholding):
		taxWithholding.save()
		return taxWithholding;
	
	def delete(self, taxWithholdingId ):
		errMsg = "Failed to delete TaxWithholding from db using id " + str(taxWithholdingId)
		
		try:
			taxWithholding = TaxWithholding.objects.get(id=taxWithholdingId)
			taxWithholding.delete()
			return True
		except TaxWithholding.DoesNotExist:
			raise ProcessingError("TaxWithholding with id " + str(taxWithholdingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TaxWithholding.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TaxWithholding from db")
		except Exception:
			return None;
		
	def assignEmployee( self, taxWithholdingId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on TaxWithholding"

		try:
			# get the TaxWithholding from db
			taxWithholding = self.get( taxWithholdingId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			taxWithholding.employee = employee
			
			#save it
			taxWithholding.save()

			# reload and return the appropriate version					
			return self.get( taxWithholdingId );
		except TaxWithholding.DoesNotExist:
			raise ProcessingError(errMsg + " : TaxWithholding with id " + str(taxWithholdingId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, taxWithholdingId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on TaxWithholding"

		try:
			# get the TaxWithholding from db
			taxWithholding = self.get( taxWithholdingId ).first()	
			
			# assign to None for unassignment
			taxWithholding.employee = None			

			#save it
			taxWithholding.save()

			# reload and return the appropriate version					
			return self.get( taxWithholdingId );
		except TaxWithholding.DoesNotExist:
			raise ProcessingError(errMsg + " : TaxWithholding with id " + str(taxWithholdingId) + " does not exist.")
		except Exception:
			return None;
		
