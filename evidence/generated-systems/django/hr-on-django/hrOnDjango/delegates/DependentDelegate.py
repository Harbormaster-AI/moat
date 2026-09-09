from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Dependent import Dependent
from hrOnDjango.models.BenefitEnrollment import BenefitEnrollment
from hrOnDjango.models.Employee import Employee
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Dependent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DependentDelegate Declaration
#======================================================================
class DependentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dependentId ):
		try:	
			dependent = Dependent.objects.filter(id=dependentId)
			return dependent.first();
		except Dependent.DoesNotExist:
			raise ProcessingError("Dependent with id " + str(dependentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dependent):
		for model in serializers.deserialize("json", dependent):
			model.save()
			return model;

	def create(self, dependent):
		dependent.save()
		return dependent;

	def saveFromJson(self, dependent):
		for model in serializers.deserialize("json", dependent):
			model.save()
			return dependent;
	
	def save(self, dependent):
		dependent.save()
		return dependent;
	
	def delete(self, dependentId ):
		errMsg = "Failed to delete Dependent from db using id " + str(dependentId)
		
		try:
			dependent = Dependent.objects.get(id=dependentId)
			dependent.delete()
			return True
		except Dependent.DoesNotExist:
			raise ProcessingError("Dependent with id " + str(dependentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Dependent.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Dependent from db")
		except Exception:
			return None;
		
	def assignBenefitEnrollment( self, dependentId, benefitEnrollmentId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.BenefitEnrollmentDelegate import BenefitEnrollmentDelegate

		errMsg = "Failed to assign element " + str(benefitEnrollmentId) + " for BenefitEnrollment on Dependent"

		try:
			# get the Dependent from db
			dependent = self.get( dependentId ).first()	
			
			# get the BenefitEnrollment from db
			benefitEnrollment = BenefitEnrollmentDelegate().get(benefitEnrollmentId).first();
			
			# assign the BenefitEnrollment		
			dependent.benefitEnrollment = benefitEnrollment
			
			#save it
			dependent.save()

			# reload and return the appropriate version					
			return self.get( dependentId );
		except Dependent.DoesNotExist:
			raise ProcessingError(errMsg + " : Dependent with id " + str(dependentId) + " does not exist.")
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitEnrollment with id " + str(benefitEnrollmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBenefitEnrollment( self, dependentId ):
		errMsg = "Failed to unassign element " + str(benefitEnrollmentId) + " for BenefitEnrollment on Dependent"

		try:
			# get the Dependent from db
			dependent = self.get( dependentId ).first()	
			
			# assign to None for unassignment
			dependent.benefitEnrollment = None			

			#save it
			dependent.save()

			# reload and return the appropriate version					
			return self.get( dependentId );
		except Dependent.DoesNotExist:
			raise ProcessingError(errMsg + " : Dependent with id " + str(dependentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEmployee( self, dependentId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on Dependent"

		try:
			# get the Dependent from db
			dependent = self.get( dependentId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			dependent.employee = employee
			
			#save it
			dependent.save()

			# reload and return the appropriate version					
			return self.get( dependentId );
		except Dependent.DoesNotExist:
			raise ProcessingError(errMsg + " : Dependent with id " + str(dependentId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, dependentId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on Dependent"

		try:
			# get the Dependent from db
			dependent = self.get( dependentId ).first()	
			
			# assign to None for unassignment
			dependent.employee = None			

			#save it
			dependent.save()

			# reload and return the appropriate version					
			return self.get( dependentId );
		except Dependent.DoesNotExist:
			raise ProcessingError(errMsg + " : Dependent with id " + str(dependentId) + " does not exist.")
		except Exception:
			return None;
		
