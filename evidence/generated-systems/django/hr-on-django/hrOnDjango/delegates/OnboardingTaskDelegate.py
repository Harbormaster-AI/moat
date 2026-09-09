from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.OnboardingTask import OnboardingTask
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.Offer import Offer
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model OnboardingTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OnboardingTaskDelegate Declaration
#======================================================================
class OnboardingTaskDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, onboardingTaskId ):
		try:	
			onboardingTask = OnboardingTask.objects.filter(id=onboardingTaskId)
			return onboardingTask.first();
		except OnboardingTask.DoesNotExist:
			raise ProcessingError("OnboardingTask with id " + str(onboardingTaskId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, onboardingTask):
		for model in serializers.deserialize("json", onboardingTask):
			model.save()
			return model;

	def create(self, onboardingTask):
		onboardingTask.save()
		return onboardingTask;

	def saveFromJson(self, onboardingTask):
		for model in serializers.deserialize("json", onboardingTask):
			model.save()
			return onboardingTask;
	
	def save(self, onboardingTask):
		onboardingTask.save()
		return onboardingTask;
	
	def delete(self, onboardingTaskId ):
		errMsg = "Failed to delete OnboardingTask from db using id " + str(onboardingTaskId)
		
		try:
			onboardingTask = OnboardingTask.objects.get(id=onboardingTaskId)
			onboardingTask.delete()
			return True
		except OnboardingTask.DoesNotExist:
			raise ProcessingError("OnboardingTask with id " + str(onboardingTaskId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = OnboardingTask.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all OnboardingTask from db")
		except Exception:
			return None;
		
	def assignEmployee( self, onboardingTaskId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on OnboardingTask"

		try:
			# get the OnboardingTask from db
			onboardingTask = self.get( onboardingTaskId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			onboardingTask.employee = employee
			
			#save it
			onboardingTask.save()

			# reload and return the appropriate version					
			return self.get( onboardingTaskId );
		except OnboardingTask.DoesNotExist:
			raise ProcessingError(errMsg + " : OnboardingTask with id " + str(onboardingTaskId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, onboardingTaskId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on OnboardingTask"

		try:
			# get the OnboardingTask from db
			onboardingTask = self.get( onboardingTaskId ).first()	
			
			# assign to None for unassignment
			onboardingTask.employee = None			

			#save it
			onboardingTask.save()

			# reload and return the appropriate version					
			return self.get( onboardingTaskId );
		except OnboardingTask.DoesNotExist:
			raise ProcessingError(errMsg + " : OnboardingTask with id " + str(onboardingTaskId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAssignedTo( self, onboardingTaskId, assignedToId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(assignedToId) + " for AssignedTo on OnboardingTask"

		try:
			# get the OnboardingTask from db
			onboardingTask = self.get( onboardingTaskId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(assignedToId).first();
			
			# assign the AssignedTo		
			onboardingTask.assignedTo = employee
			
			#save it
			onboardingTask.save()

			# reload and return the appropriate version					
			return self.get( onboardingTaskId );
		except OnboardingTask.DoesNotExist:
			raise ProcessingError(errMsg + " : OnboardingTask with id " + str(onboardingTaskId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(assignedToId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAssignedTo( self, onboardingTaskId ):
		errMsg = "Failed to unassign element " + str(assignedToId) + " for AssignedTo on OnboardingTask"

		try:
			# get the OnboardingTask from db
			onboardingTask = self.get( onboardingTaskId ).first()	
			
			# assign to None for unassignment
			onboardingTask.employee = None			

			#save it
			onboardingTask.save()

			# reload and return the appropriate version					
			return self.get( onboardingTaskId );
		except OnboardingTask.DoesNotExist:
			raise ProcessingError(errMsg + " : OnboardingTask with id " + str(onboardingTaskId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRelatedOffer( self, onboardingTaskId, relatedOfferId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OfferDelegate import OfferDelegate

		errMsg = "Failed to assign element " + str(relatedOfferId) + " for RelatedOffer on OnboardingTask"

		try:
			# get the OnboardingTask from db
			onboardingTask = self.get( onboardingTaskId ).first()	
			
			# get the Offer from db
			offer = OfferDelegate().get(relatedOfferId).first();
			
			# assign the RelatedOffer		
			onboardingTask.relatedOffer = offer
			
			#save it
			onboardingTask.save()

			# reload and return the appropriate version					
			return self.get( onboardingTaskId );
		except OnboardingTask.DoesNotExist:
			raise ProcessingError(errMsg + " : OnboardingTask with id " + str(onboardingTaskId) + " does not exist.")
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer with id " + str(relatedOfferId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRelatedOffer( self, onboardingTaskId ):
		errMsg = "Failed to unassign element " + str(relatedOfferId) + " for RelatedOffer on OnboardingTask"

		try:
			# get the OnboardingTask from db
			onboardingTask = self.get( onboardingTaskId ).first()	
			
			# assign to None for unassignment
			onboardingTask.offer = None			

			#save it
			onboardingTask.save()

			# reload and return the appropriate version					
			return self.get( onboardingTaskId );
		except OnboardingTask.DoesNotExist:
			raise ProcessingError(errMsg + " : OnboardingTask with id " + str(onboardingTaskId) + " does not exist.")
		except Exception:
			return None;
		
	def addDependencies( self, onboardingTaskId, dependenciesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OnboardingTaskDelegate import OnboardingTaskDelegate

		errMsg = "Failed to add elements " + str(dependenciesIds) + " for Dependencies on OnboardingTask"

		try:
			# get the OnboardingTask
			onboardingTask = self.get( onboardingTaskId ).first()
				
			# split on a comma with no spaces
			idList = dependenciesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the OnboardingTask		
				onboardingTask = OnboardingTaskDelegate().get(id).first();	
				# add the OnboardingTask
				onboardingTask.dependencies.add(onboardingTask)
				
			# save it		
			onboardingTask.save()
			
			# reload and return the appropriate version
			return self.get( onboardingTaskId );
		except OnboardingTask.DoesNotExist:
			raise ProcessingError(errMsg + " : OnboardingTask with id " + str(onboardingTaskId) + " does not exist.")
		except OnboardingTask.DoesNotExist:
			raise ProcessingError(errMsg + " : OnboardingTask does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDependencies( self, onboardingTaskId, dependenciesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OnboardingTaskDelegate import OnboardingTaskDelegate

		errMsg = "Failed to remove elements " + str(dependenciesIds) + " for Dependencies on OnboardingTask"

		try:
			# get the OnboardingTask
			onboardingTask = self.get( onboardingTaskId ).first()
				
			# split on a comma with no spaces
			idList = dependenciesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the OnboardingTask		
				onboardingTask = OnboardingTaskDelegate().get(id).first();	
				# add the OnboardingTask
				onboardingTask.dependencies.remove(onboardingTask)
				
			# save it		
			onboardingTask.save()
			
			# reload and return the appropriate version
			return self.get( onboardingTaskId );
		except OnboardingTask.DoesNotExist:
			raise ProcessingError(errMsg + " : OnboardingTask with id " + str(onboardingTaskId) + " does not exist.")
		except OnboardingTask.DoesNotExist:
			raise ProcessingError(errMsg + " : OnboardingTask does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
