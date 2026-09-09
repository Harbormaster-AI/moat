from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.PerformanceCycle import PerformanceCycle
from hrOnDjango.models.Organization import Organization
from hrOnDjango.models.PerformanceReview import PerformanceReview
from hrOnDjango.models.Goal import Goal
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PerformanceCycle
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceCycleDelegate Declaration
#======================================================================
class PerformanceCycleDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, performanceCycleId ):
		try:	
			performanceCycle = PerformanceCycle.objects.filter(id=performanceCycleId)
			return performanceCycle.first();
		except PerformanceCycle.DoesNotExist:
			raise ProcessingError("PerformanceCycle with id " + str(performanceCycleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, performanceCycle):
		for model in serializers.deserialize("json", performanceCycle):
			model.save()
			return model;

	def create(self, performanceCycle):
		performanceCycle.save()
		return performanceCycle;

	def saveFromJson(self, performanceCycle):
		for model in serializers.deserialize("json", performanceCycle):
			model.save()
			return performanceCycle;
	
	def save(self, performanceCycle):
		performanceCycle.save()
		return performanceCycle;
	
	def delete(self, performanceCycleId ):
		errMsg = "Failed to delete PerformanceCycle from db using id " + str(performanceCycleId)
		
		try:
			performanceCycle = PerformanceCycle.objects.get(id=performanceCycleId)
			performanceCycle.delete()
			return True
		except PerformanceCycle.DoesNotExist:
			raise ProcessingError("PerformanceCycle with id " + str(performanceCycleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PerformanceCycle.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PerformanceCycle from db")
		except Exception:
			return None;
		
	def assignOrganization( self, performanceCycleId, organizationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on PerformanceCycle"

		try:
			# get the PerformanceCycle from db
			performanceCycle = self.get( performanceCycleId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			performanceCycle.organization = organization
			
			#save it
			performanceCycle.save()

			# reload and return the appropriate version					
			return self.get( performanceCycleId );
		except PerformanceCycle.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceCycle with id " + str(performanceCycleId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, performanceCycleId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on PerformanceCycle"

		try:
			# get the PerformanceCycle from db
			performanceCycle = self.get( performanceCycleId ).first()	
			
			# assign to None for unassignment
			performanceCycle.organization = None			

			#save it
			performanceCycle.save()

			# reload and return the appropriate version					
			return self.get( performanceCycleId );
		except PerformanceCycle.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceCycle with id " + str(performanceCycleId) + " does not exist.")
		except Exception:
			return None;
		
	def addReviews( self, performanceCycleId, reviewsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PerformanceReviewDelegate import PerformanceReviewDelegate

		errMsg = "Failed to add elements " + str(reviewsIds) + " for Reviews on PerformanceCycle"

		try:
			# get the PerformanceCycle
			performanceCycle = self.get( performanceCycleId ).first()
				
			# split on a comma with no spaces
			idList = reviewsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PerformanceReview		
				performanceReview = PerformanceReviewDelegate().get(id).first();	
				# add the PerformanceReview
				performanceCycle.reviews.add(performanceReview)
				
			# save it		
			performanceCycle.save()
			
			# reload and return the appropriate version
			return self.get( performanceCycleId );
		except PerformanceCycle.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceCycle with id " + str(performanceCycleId) + " does not exist.")
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReviews( self, performanceCycleId, reviewsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PerformanceReviewDelegate import PerformanceReviewDelegate

		errMsg = "Failed to remove elements " + str(reviewsIds) + " for Reviews on PerformanceCycle"

		try:
			# get the PerformanceCycle
			performanceCycle = self.get( performanceCycleId ).first()
				
			# split on a comma with no spaces
			idList = reviewsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PerformanceReview		
				performanceReview = PerformanceReviewDelegate().get(id).first();	
				# add the PerformanceReview
				performanceCycle.reviews.remove(performanceReview)
				
			# save it		
			performanceCycle.save()
			
			# reload and return the appropriate version
			return self.get( performanceCycleId );
		except PerformanceCycle.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceCycle with id " + str(performanceCycleId) + " does not exist.")
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addGoals( self, performanceCycleId, goalsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.GoalDelegate import GoalDelegate

		errMsg = "Failed to add elements " + str(goalsIds) + " for Goals on PerformanceCycle"

		try:
			# get the PerformanceCycle
			performanceCycle = self.get( performanceCycleId ).first()
				
			# split on a comma with no spaces
			idList = goalsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Goal		
				goal = GoalDelegate().get(id).first();	
				# add the Goal
				performanceCycle.goals.add(goal)
				
			# save it		
			performanceCycle.save()
			
			# reload and return the appropriate version
			return self.get( performanceCycleId );
		except PerformanceCycle.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceCycle with id " + str(performanceCycleId) + " does not exist.")
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeGoals( self, performanceCycleId, goalsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.GoalDelegate import GoalDelegate

		errMsg = "Failed to remove elements " + str(goalsIds) + " for Goals on PerformanceCycle"

		try:
			# get the PerformanceCycle
			performanceCycle = self.get( performanceCycleId ).first()
				
			# split on a comma with no spaces
			idList = goalsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Goal		
				goal = GoalDelegate().get(id).first();	
				# add the Goal
				performanceCycle.goals.remove(goal)
				
			# save it		
			performanceCycle.save()
			
			# reload and return the appropriate version
			return self.get( performanceCycleId );
		except PerformanceCycle.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceCycle with id " + str(performanceCycleId) + " does not exist.")
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
