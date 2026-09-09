from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Goal import Goal
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.PerformanceCycle import PerformanceCycle
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Goal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoalDelegate Declaration
#======================================================================
class GoalDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, goalId ):
		try:	
			goal = Goal.objects.filter(id=goalId)
			return goal.first();
		except Goal.DoesNotExist:
			raise ProcessingError("Goal with id " + str(goalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, goal):
		for model in serializers.deserialize("json", goal):
			model.save()
			return model;

	def create(self, goal):
		goal.save()
		return goal;

	def saveFromJson(self, goal):
		for model in serializers.deserialize("json", goal):
			model.save()
			return goal;
	
	def save(self, goal):
		goal.save()
		return goal;
	
	def delete(self, goalId ):
		errMsg = "Failed to delete Goal from db using id " + str(goalId)
		
		try:
			goal = Goal.objects.get(id=goalId)
			goal.delete()
			return True
		except Goal.DoesNotExist:
			raise ProcessingError("Goal with id " + str(goalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Goal.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Goal from db")
		except Exception:
			return None;
		
	def assignEmployee( self, goalId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on Goal"

		try:
			# get the Goal from db
			goal = self.get( goalId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			goal.employee = employee
			
			#save it
			goal.save()

			# reload and return the appropriate version					
			return self.get( goalId );
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal with id " + str(goalId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, goalId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on Goal"

		try:
			# get the Goal from db
			goal = self.get( goalId ).first()	
			
			# assign to None for unassignment
			goal.employee = None			

			#save it
			goal.save()

			# reload and return the appropriate version					
			return self.get( goalId );
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal with id " + str(goalId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCycle( self, goalId, cycleId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PerformanceCycleDelegate import PerformanceCycleDelegate

		errMsg = "Failed to assign element " + str(cycleId) + " for Cycle on Goal"

		try:
			# get the Goal from db
			goal = self.get( goalId ).first()	
			
			# get the PerformanceCycle from db
			performanceCycle = PerformanceCycleDelegate().get(cycleId).first();
			
			# assign the Cycle		
			goal.cycle = performanceCycle
			
			#save it
			goal.save()

			# reload and return the appropriate version					
			return self.get( goalId );
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal with id " + str(goalId) + " does not exist.")
		except PerformanceCycle.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceCycle with id " + str(cycleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCycle( self, goalId ):
		errMsg = "Failed to unassign element " + str(cycleId) + " for Cycle on Goal"

		try:
			# get the Goal from db
			goal = self.get( goalId ).first()	
			
			# assign to None for unassignment
			goal.performanceCycle = None			

			#save it
			goal.save()

			# reload and return the appropriate version					
			return self.get( goalId );
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal with id " + str(goalId) + " does not exist.")
		except Exception:
			return None;
		
	def assignParentGoal( self, goalId, parentGoalId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.GoalDelegate import GoalDelegate

		errMsg = "Failed to assign element " + str(parentGoalId) + " for ParentGoal on Goal"

		try:
			# get the Goal from db
			goal = self.get( goalId ).first()	
			
			# get the Goal from db
			goal = GoalDelegate().get(parentGoalId).first();
			
			# assign the ParentGoal		
			goal.parentGoal = goal
			
			#save it
			goal.save()

			# reload and return the appropriate version					
			return self.get( goalId );
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal with id " + str(goalId) + " does not exist.")
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal with id " + str(parentGoalId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignParentGoal( self, goalId ):
		errMsg = "Failed to unassign element " + str(parentGoalId) + " for ParentGoal on Goal"

		try:
			# get the Goal from db
			goal = self.get( goalId ).first()	
			
			# assign to None for unassignment
			goal.goal = None			

			#save it
			goal.save()

			# reload and return the appropriate version					
			return self.get( goalId );
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal with id " + str(goalId) + " does not exist.")
		except Exception:
			return None;
		
	def addChildGoals( self, goalId, childGoalsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.GoalDelegate import GoalDelegate

		errMsg = "Failed to add elements " + str(childGoalsIds) + " for ChildGoals on Goal"

		try:
			# get the Goal
			goal = self.get( goalId ).first()
				
			# split on a comma with no spaces
			idList = childGoalsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Goal		
				goal = GoalDelegate().get(id).first();	
				# add the Goal
				goal.childGoals.add(goal)
				
			# save it		
			goal.save()
			
			# reload and return the appropriate version
			return self.get( goalId );
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal with id " + str(goalId) + " does not exist.")
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChildGoals( self, goalId, childGoalsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.GoalDelegate import GoalDelegate

		errMsg = "Failed to remove elements " + str(childGoalsIds) + " for ChildGoals on Goal"

		try:
			# get the Goal
			goal = self.get( goalId ).first()
				
			# split on a comma with no spaces
			idList = childGoalsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Goal		
				goal = GoalDelegate().get(id).first();	
				# add the Goal
				goal.childGoals.remove(goal)
				
			# save it		
			goal.save()
			
			# reload and return the appropriate version
			return self.get( goalId );
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal with id " + str(goalId) + " does not exist.")
		except Goal.DoesNotExist:
			raise ProcessingError(errMsg + " : Goal does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
