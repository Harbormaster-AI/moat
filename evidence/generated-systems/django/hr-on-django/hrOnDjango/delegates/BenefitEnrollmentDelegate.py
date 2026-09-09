from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.BenefitEnrollment import BenefitEnrollment
from hrOnDjango.models.BenefitPlan import BenefitPlan
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.Dependent import Dependent
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BenefitEnrollment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BenefitEnrollmentDelegate Declaration
#======================================================================
class BenefitEnrollmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, benefitEnrollmentId ):
		try:	
			benefitEnrollment = BenefitEnrollment.objects.filter(id=benefitEnrollmentId)
			return benefitEnrollment.first();
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError("BenefitEnrollment with id " + str(benefitEnrollmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, benefitEnrollment):
		for model in serializers.deserialize("json", benefitEnrollment):
			model.save()
			return model;

	def create(self, benefitEnrollment):
		benefitEnrollment.save()
		return benefitEnrollment;

	def saveFromJson(self, benefitEnrollment):
		for model in serializers.deserialize("json", benefitEnrollment):
			model.save()
			return benefitEnrollment;
	
	def save(self, benefitEnrollment):
		benefitEnrollment.save()
		return benefitEnrollment;
	
	def delete(self, benefitEnrollmentId ):
		errMsg = "Failed to delete BenefitEnrollment from db using id " + str(benefitEnrollmentId)
		
		try:
			benefitEnrollment = BenefitEnrollment.objects.get(id=benefitEnrollmentId)
			benefitEnrollment.delete()
			return True
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError("BenefitEnrollment with id " + str(benefitEnrollmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BenefitEnrollment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BenefitEnrollment from db")
		except Exception:
			return None;
		
	def assignBenefitPlan( self, benefitEnrollmentId, benefitPlanId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.BenefitPlanDelegate import BenefitPlanDelegate

		errMsg = "Failed to assign element " + str(benefitPlanId) + " for BenefitPlan on BenefitEnrollment"

		try:
			# get the BenefitEnrollment from db
			benefitEnrollment = self.get( benefitEnrollmentId ).first()	
			
			# get the BenefitPlan from db
			benefitPlan = BenefitPlanDelegate().get(benefitPlanId).first();
			
			# assign the BenefitPlan		
			benefitEnrollment.benefitPlan = benefitPlan
			
			#save it
			benefitEnrollment.save()

			# reload and return the appropriate version					
			return self.get( benefitEnrollmentId );
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitEnrollment with id " + str(benefitEnrollmentId) + " does not exist.")
		except BenefitPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitPlan with id " + str(benefitPlanId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBenefitPlan( self, benefitEnrollmentId ):
		errMsg = "Failed to unassign element " + str(benefitPlanId) + " for BenefitPlan on BenefitEnrollment"

		try:
			# get the BenefitEnrollment from db
			benefitEnrollment = self.get( benefitEnrollmentId ).first()	
			
			# assign to None for unassignment
			benefitEnrollment.benefitPlan = None			

			#save it
			benefitEnrollment.save()

			# reload and return the appropriate version					
			return self.get( benefitEnrollmentId );
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitEnrollment with id " + str(benefitEnrollmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEmployee( self, benefitEnrollmentId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on BenefitEnrollment"

		try:
			# get the BenefitEnrollment from db
			benefitEnrollment = self.get( benefitEnrollmentId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			benefitEnrollment.employee = employee
			
			#save it
			benefitEnrollment.save()

			# reload and return the appropriate version					
			return self.get( benefitEnrollmentId );
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitEnrollment with id " + str(benefitEnrollmentId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, benefitEnrollmentId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on BenefitEnrollment"

		try:
			# get the BenefitEnrollment from db
			benefitEnrollment = self.get( benefitEnrollmentId ).first()	
			
			# assign to None for unassignment
			benefitEnrollment.employee = None			

			#save it
			benefitEnrollment.save()

			# reload and return the appropriate version					
			return self.get( benefitEnrollmentId );
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitEnrollment with id " + str(benefitEnrollmentId) + " does not exist.")
		except Exception:
			return None;
		
	def addDependents( self, benefitEnrollmentId, dependentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DependentDelegate import DependentDelegate

		errMsg = "Failed to add elements " + str(dependentsIds) + " for Dependents on BenefitEnrollment"

		try:
			# get the BenefitEnrollment
			benefitEnrollment = self.get( benefitEnrollmentId ).first()
				
			# split on a comma with no spaces
			idList = dependentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dependent		
				dependent = DependentDelegate().get(id).first();	
				# add the Dependent
				benefitEnrollment.dependents.add(dependent)
				
			# save it		
			benefitEnrollment.save()
			
			# reload and return the appropriate version
			return self.get( benefitEnrollmentId );
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitEnrollment with id " + str(benefitEnrollmentId) + " does not exist.")
		except Dependent.DoesNotExist:
			raise ProcessingError(errMsg + " : Dependent does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDependents( self, benefitEnrollmentId, dependentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DependentDelegate import DependentDelegate

		errMsg = "Failed to remove elements " + str(dependentsIds) + " for Dependents on BenefitEnrollment"

		try:
			# get the BenefitEnrollment
			benefitEnrollment = self.get( benefitEnrollmentId ).first()
				
			# split on a comma with no spaces
			idList = dependentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dependent		
				dependent = DependentDelegate().get(id).first();	
				# add the Dependent
				benefitEnrollment.dependents.remove(dependent)
				
			# save it		
			benefitEnrollment.save()
			
			# reload and return the appropriate version
			return self.get( benefitEnrollmentId );
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitEnrollment with id " + str(benefitEnrollmentId) + " does not exist.")
		except Dependent.DoesNotExist:
			raise ProcessingError(errMsg + " : Dependent does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
