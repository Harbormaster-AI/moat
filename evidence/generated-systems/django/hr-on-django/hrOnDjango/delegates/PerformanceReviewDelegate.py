from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.PerformanceReview import PerformanceReview
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.PerformanceCycle import PerformanceCycle
from hrOnDjango.models.CompetencyRating import CompetencyRating
from hrOnDjango.models.Goal import Goal
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PerformanceReview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceReviewDelegate Declaration
#======================================================================
class PerformanceReviewDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, performanceReviewId ):
		try:	
			performanceReview = PerformanceReview.objects.filter(id=performanceReviewId)
			return performanceReview.first();
		except PerformanceReview.DoesNotExist:
			raise ProcessingError("PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, performanceReview):
		for model in serializers.deserialize("json", performanceReview):
			model.save()
			return model;

	def create(self, performanceReview):
		performanceReview.save()
		return performanceReview;

	def saveFromJson(self, performanceReview):
		for model in serializers.deserialize("json", performanceReview):
			model.save()
			return performanceReview;
	
	def save(self, performanceReview):
		performanceReview.save()
		return performanceReview;
	
	def delete(self, performanceReviewId ):
		errMsg = "Failed to delete PerformanceReview from db using id " + str(performanceReviewId)
		
		try:
			performanceReview = PerformanceReview.objects.get(id=performanceReviewId)
			performanceReview.delete()
			return True
		except PerformanceReview.DoesNotExist:
			raise ProcessingError("PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PerformanceReview.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PerformanceReview from db")
		except Exception:
			return None;
		
	def assignEmployee( self, performanceReviewId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on PerformanceReview"

		try:
			# get the PerformanceReview from db
			performanceReview = self.get( performanceReviewId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			performanceReview.employee = employee
			
			#save it
			performanceReview.save()

			# reload and return the appropriate version					
			return self.get( performanceReviewId );
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, performanceReviewId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on PerformanceReview"

		try:
			# get the PerformanceReview from db
			performanceReview = self.get( performanceReviewId ).first()	
			
			# assign to None for unassignment
			performanceReview.employee = None			

			#save it
			performanceReview.save()

			# reload and return the appropriate version					
			return self.get( performanceReviewId );
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except Exception:
			return None;
		
	def assignReviewer( self, performanceReviewId, reviewerId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(reviewerId) + " for Reviewer on PerformanceReview"

		try:
			# get the PerformanceReview from db
			performanceReview = self.get( performanceReviewId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(reviewerId).first();
			
			# assign the Reviewer		
			performanceReview.reviewer = employee
			
			#save it
			performanceReview.save()

			# reload and return the appropriate version					
			return self.get( performanceReviewId );
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(reviewerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignReviewer( self, performanceReviewId ):
		errMsg = "Failed to unassign element " + str(reviewerId) + " for Reviewer on PerformanceReview"

		try:
			# get the PerformanceReview from db
			performanceReview = self.get( performanceReviewId ).first()	
			
			# assign to None for unassignment
			performanceReview.employee = None			

			#save it
			performanceReview.save()

			# reload and return the appropriate version					
			return self.get( performanceReviewId );
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCycle( self, performanceReviewId, cycleId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PerformanceCycleDelegate import PerformanceCycleDelegate

		errMsg = "Failed to assign element " + str(cycleId) + " for Cycle on PerformanceReview"

		try:
			# get the PerformanceReview from db
			performanceReview = self.get( performanceReviewId ).first()	
			
			# get the PerformanceCycle from db
			performanceCycle = PerformanceCycleDelegate().get(cycleId).first();
			
			# assign the Cycle		
			performanceReview.cycle = performanceCycle
			
			#save it
			performanceReview.save()

			# reload and return the appropriate version					
			return self.get( performanceReviewId );
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except PerformanceCycle.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceCycle with id " + str(cycleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCycle( self, performanceReviewId ):
		errMsg = "Failed to unassign element " + str(cycleId) + " for Cycle on PerformanceReview"

		try:
			# get the PerformanceReview from db
			performanceReview = self.get( performanceReviewId ).first()	
			
			# assign to None for unassignment
			performanceReview.performanceCycle = None			

			#save it
			performanceReview.save()

			# reload and return the appropriate version					
			return self.get( performanceReviewId );
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except Exception:
			return None;
		
	def addCompetencyRatings( self, performanceReviewId, competencyRatingsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompetencyRatingDelegate import CompetencyRatingDelegate

		errMsg = "Failed to add elements " + str(competencyRatingsIds) + " for CompetencyRatings on PerformanceReview"

		try:
			# get the PerformanceReview
			performanceReview = self.get( performanceReviewId ).first()
				
			# split on a comma with no spaces
			idList = competencyRatingsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CompetencyRating		
				competencyRating = CompetencyRatingDelegate().get(id).first();	
				# add the CompetencyRating
				performanceReview.competencyRatings.add(competencyRating)
				
			# save it		
			performanceReview.save()
			
			# reload and return the appropriate version
			return self.get( performanceReviewId );
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except CompetencyRating.DoesNotExist:
			raise ProcessingError(errMsg + " : CompetencyRating does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCompetencyRatings( self, performanceReviewId, competencyRatingsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompetencyRatingDelegate import CompetencyRatingDelegate

		errMsg = "Failed to remove elements " + str(competencyRatingsIds) + " for CompetencyRatings on PerformanceReview"

		try:
			# get the PerformanceReview
			performanceReview = self.get( performanceReviewId ).first()
				
			# split on a comma with no spaces
			idList = competencyRatingsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CompetencyRating		
				competencyRating = CompetencyRatingDelegate().get(id).first();	
				# add the CompetencyRating
				performanceReview.competencyRatings.remove(competencyRating)
				
			# save it		
			performanceReview.save()
			
			# reload and return the appropriate version
			return self.get( performanceReviewId );
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except CompetencyRating.DoesNotExist:
			raise ProcessingError(errMsg + " : CompetencyRating does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addGoals( self, performanceReviewId, goalsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.GoalDelegate import GoalDelegate

		errMsg = "Failed to add elements " + str(goalsIds) + " for Goals on PerformanceReview"

		try:
			# get the PerformanceReview
			performanceReview = self.get( performanceReviewId ).first()
				
			# split on a comma with no spaces
			idList = goalsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Goal		
				goal = GoalDelegate().get(id).first();	
				# add the Goal
				performanceReview.goals.add(goal)
				
			# save it		
			performanceReview.save()
			
			# reload and return the appropriate version
			return self.get( performanceReviewId );
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeGoals( self, performanceReviewId, goalsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.GoalDelegate import GoalDelegate

		errMsg = "Failed to remove elements " + str(goalsIds) + " for Goals on PerformanceReview"

		try:
			# get the PerformanceReview
			performanceReview = self.get( performanceReviewId ).first()
				
			# split on a comma with no spaces
			idList = goalsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Goal		
				goal = GoalDelegate().get(id).first();	
				# add the Goal
				performanceReview.goals.remove(goal)
				
			# save it		
			performanceReview.save()
			
			# reload and return the appropriate version
			return self.get( performanceReviewId );
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview with id " + str(performanceReviewId) + " does not exist.")
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
