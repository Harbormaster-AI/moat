from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.BIQuery import BIQuery
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Report import Report
from analyticsOnDjango.models.Dashboard import Dashboard
from analyticsOnDjango.models.Notebook import Notebook
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BIQuery
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BIQueryDelegate Declaration
#======================================================================
class BIQueryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, bIQueryId ):
		try:	
			bIQuery = BIQuery.objects.filter(id=bIQueryId)
			return bIQuery.first();
		except BIQuery.DoesNotExist:
			raise ProcessingError("BIQuery with id " + str(bIQueryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, bIQuery):
		for model in serializers.deserialize("json", bIQuery):
			model.save()
			return model;

	def create(self, bIQuery):
		bIQuery.save()
		return bIQuery;

	def saveFromJson(self, bIQuery):
		for model in serializers.deserialize("json", bIQuery):
			model.save()
			return bIQuery;
	
	def save(self, bIQuery):
		bIQuery.save()
		return bIQuery;
	
	def delete(self, bIQueryId ):
		errMsg = "Failed to delete BIQuery from db using id " + str(bIQueryId)
		
		try:
			bIQuery = BIQuery.objects.get(id=bIQueryId)
			bIQuery.delete()
			return True
		except BIQuery.DoesNotExist:
			raise ProcessingError("BIQuery with id " + str(bIQueryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BIQuery.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BIQuery from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, bIQueryId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on BIQuery"

		try:
			# get the BIQuery from db
			bIQuery = self.get( bIQueryId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			bIQuery.workspace = analyticsWorkspace
			
			#save it
			bIQuery.save()

			# reload and return the appropriate version					
			return self.get( bIQueryId );
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery with id " + str(bIQueryId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, bIQueryId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on BIQuery"

		try:
			# get the BIQuery from db
			bIQuery = self.get( bIQueryId ).first()	
			
			# assign to None for unassignment
			bIQuery.analyticsWorkspace = None			

			#save it
			bIQuery.save()

			# reload and return the appropriate version					
			return self.get( bIQueryId );
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery with id " + str(bIQueryId) + " does not exist.")
		except Exception:
			return None;
		
	def addDatasets( self, bIQueryId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on BIQuery"

		try:
			# get the BIQuery
			bIQuery = self.get( bIQueryId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				bIQuery.datasets.add(dataSet)
				
			# save it		
			bIQuery.save()
			
			# reload and return the appropriate version
			return self.get( bIQueryId );
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery with id " + str(bIQueryId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, bIQueryId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on BIQuery"

		try:
			# get the BIQuery
			bIQuery = self.get( bIQueryId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				bIQuery.datasets.remove(dataSet)
				
			# save it		
			bIQuery.save()
			
			# reload and return the appropriate version
			return self.get( bIQueryId );
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery with id " + str(bIQueryId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReports( self, bIQueryId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to add elements " + str(reportsIds) + " for Reports on BIQuery"

		try:
			# get the BIQuery
			bIQuery = self.get( bIQueryId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				bIQuery.reports.add(report)
				
			# save it		
			bIQuery.save()
			
			# reload and return the appropriate version
			return self.get( bIQueryId );
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery with id " + str(bIQueryId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReports( self, bIQueryId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to remove elements " + str(reportsIds) + " for Reports on BIQuery"

		try:
			# get the BIQuery
			bIQuery = self.get( bIQueryId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				bIQuery.reports.remove(report)
				
			# save it		
			bIQuery.save()
			
			# reload and return the appropriate version
			return self.get( bIQueryId );
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery with id " + str(bIQueryId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDashboards( self, bIQueryId, dashboardsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to add elements " + str(dashboardsIds) + " for Dashboards on BIQuery"

		try:
			# get the BIQuery
			bIQuery = self.get( bIQueryId ).first()
				
			# split on a comma with no spaces
			idList = dashboardsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dashboard		
				dashboard = DashboardDelegate().get(id).first();	
				# add the Dashboard
				bIQuery.dashboards.add(dashboard)
				
			# save it		
			bIQuery.save()
			
			# reload and return the appropriate version
			return self.get( bIQueryId );
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery with id " + str(bIQueryId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDashboards( self, bIQueryId, dashboardsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to remove elements " + str(dashboardsIds) + " for Dashboards on BIQuery"

		try:
			# get the BIQuery
			bIQuery = self.get( bIQueryId ).first()
				
			# split on a comma with no spaces
			idList = dashboardsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dashboard		
				dashboard = DashboardDelegate().get(id).first();	
				# add the Dashboard
				bIQuery.dashboards.remove(dashboard)
				
			# save it		
			bIQuery.save()
			
			# reload and return the appropriate version
			return self.get( bIQueryId );
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery with id " + str(bIQueryId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addNotebooks( self, bIQueryId, notebooksIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.NotebookDelegate import NotebookDelegate

		errMsg = "Failed to add elements " + str(notebooksIds) + " for Notebooks on BIQuery"

		try:
			# get the BIQuery
			bIQuery = self.get( bIQueryId ).first()
				
			# split on a comma with no spaces
			idList = notebooksIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Notebook		
				notebook = NotebookDelegate().get(id).first();	
				# add the Notebook
				bIQuery.notebooks.add(notebook)
				
			# save it		
			bIQuery.save()
			
			# reload and return the appropriate version
			return self.get( bIQueryId );
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery with id " + str(bIQueryId) + " does not exist.")
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeNotebooks( self, bIQueryId, notebooksIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.NotebookDelegate import NotebookDelegate

		errMsg = "Failed to remove elements " + str(notebooksIds) + " for Notebooks on BIQuery"

		try:
			# get the BIQuery
			bIQuery = self.get( bIQueryId ).first()
				
			# split on a comma with no spaces
			idList = notebooksIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Notebook		
				notebook = NotebookDelegate().get(id).first();	
				# add the Notebook
				bIQuery.notebooks.remove(notebook)
				
			# save it		
			bIQuery.save()
			
			# reload and return the appropriate version
			return self.get( bIQueryId );
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery with id " + str(bIQueryId) + " does not exist.")
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
