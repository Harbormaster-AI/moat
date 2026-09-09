from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.SalaryComponent import SalaryComponent
from hrOnDjango.models.CompensationPackage import CompensationPackage
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model SalaryComponent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalaryComponentDelegate Declaration
#======================================================================
class SalaryComponentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, salaryComponentId ):
		try:	
			salaryComponent = SalaryComponent.objects.filter(id=salaryComponentId)
			return salaryComponent.first();
		except SalaryComponent.DoesNotExist:
			raise ProcessingError("SalaryComponent with id " + str(salaryComponentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, salaryComponent):
		for model in serializers.deserialize("json", salaryComponent):
			model.save()
			return model;

	def create(self, salaryComponent):
		salaryComponent.save()
		return salaryComponent;

	def saveFromJson(self, salaryComponent):
		for model in serializers.deserialize("json", salaryComponent):
			model.save()
			return salaryComponent;
	
	def save(self, salaryComponent):
		salaryComponent.save()
		return salaryComponent;
	
	def delete(self, salaryComponentId ):
		errMsg = "Failed to delete SalaryComponent from db using id " + str(salaryComponentId)
		
		try:
			salaryComponent = SalaryComponent.objects.get(id=salaryComponentId)
			salaryComponent.delete()
			return True
		except SalaryComponent.DoesNotExist:
			raise ProcessingError("SalaryComponent with id " + str(salaryComponentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = SalaryComponent.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all SalaryComponent from db")
		except Exception:
			return None;
		
	def assignCompensationPackage( self, salaryComponentId, compensationPackageId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompensationPackageDelegate import CompensationPackageDelegate

		errMsg = "Failed to assign element " + str(compensationPackageId) + " for CompensationPackage on SalaryComponent"

		try:
			# get the SalaryComponent from db
			salaryComponent = self.get( salaryComponentId ).first()	
			
			# get the CompensationPackage from db
			compensationPackage = CompensationPackageDelegate().get(compensationPackageId).first();
			
			# assign the CompensationPackage		
			salaryComponent.compensationPackage = compensationPackage
			
			#save it
			salaryComponent.save()

			# reload and return the appropriate version					
			return self.get( salaryComponentId );
		except SalaryComponent.DoesNotExist:
			raise ProcessingError(errMsg + " : SalaryComponent with id " + str(salaryComponentId) + " does not exist.")
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCompensationPackage( self, salaryComponentId ):
		errMsg = "Failed to unassign element " + str(compensationPackageId) + " for CompensationPackage on SalaryComponent"

		try:
			# get the SalaryComponent from db
			salaryComponent = self.get( salaryComponentId ).first()	
			
			# assign to None for unassignment
			salaryComponent.compensationPackage = None			

			#save it
			salaryComponent.save()

			# reload and return the appropriate version					
			return self.get( salaryComponentId );
		except SalaryComponent.DoesNotExist:
			raise ProcessingError(errMsg + " : SalaryComponent with id " + str(salaryComponentId) + " does not exist.")
		except Exception:
			return None;
		
