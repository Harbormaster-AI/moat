import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Employee} from '../models/Employee';
import {DepartmentService} from '../services/Department.service';
import {LocationService} from '../services/Location.service';
import {CostCenterService} from '../services/CostCenter.service';
import {EmploymentAssignmentService} from '../services/EmploymentAssignment.service';
import {EmploymentContractService} from '../services/EmploymentContract.service';
import {BenefitEnrollmentService} from '../services/BenefitEnrollment.service';
import {TimesheetService} from '../services/Timesheet.service';
import {LeaveRequestService} from '../services/LeaveRequest.service';
import {PerformanceReviewService} from '../services/PerformanceReview.service';
import {TrainingEnrollmentService} from '../services/TrainingEnrollment.service';
import {WorkAuthorizationService} from '../services/WorkAuthorization.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EmployeeService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	employee : Employee;

	//********************************************************************
	// Catch all for the return value of a service call
	//********************************************************************
	result: any;

	//********************************************************************
	// sole constructor, injected with the HttpClient
	//********************************************************************
	constructor(private http: HttpClient) {
		super();
	}

		//********************************************************************
	// add a Employee
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEmployee(employeeNumber, name, workEmail, workPhone, dateOfHire, nationalId, Manager, DirectReports, Department, PrimaryLocation, CostCenter, EmploymentAssignments, Contracts, BenefitEnrollments, Timesheets, LeaveRequests, PerformanceReviews, TrainingEnrollments, WorkAuthorizations, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Employee/create';
		const obj = {
			      		employeeNumber: employeeNumber,
      		name: name,
      		workEmail: workEmail,
      		workPhone: workPhone,
      		dateOfHire: dateOfHire,
      		nationalId: nationalId,
      		Manager: Manager != null && Manager.length > 0 ? Manager : null,
      		DirectReports: DirectReports != null && DirectReports.length > 0 ? DirectReports : null,
      		Department: Department != null && Department.length > 0 ? Department : null,
      		PrimaryLocation: PrimaryLocation != null && PrimaryLocation.length > 0 ? PrimaryLocation : null,
      		CostCenter: CostCenter != null && CostCenter.length > 0 ? CostCenter : null,
      		EmploymentAssignments: EmploymentAssignments != null && EmploymentAssignments.length > 0 ? EmploymentAssignments : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		BenefitEnrollments: BenefitEnrollments != null && BenefitEnrollments.length > 0 ? BenefitEnrollments : null,
      		Timesheets: Timesheets != null && Timesheets.length > 0 ? Timesheets : null,
      		LeaveRequests: LeaveRequests != null && LeaveRequests.length > 0 ? LeaveRequests : null,
      		PerformanceReviews: PerformanceReviews != null && PerformanceReviews.length > 0 ? PerformanceReviews : null,
      		TrainingEnrollments: TrainingEnrollments != null && TrainingEnrollments.length > 0 ? TrainingEnrollments : null,
      		WorkAuthorizations: WorkAuthorizations != null && WorkAuthorizations.length > 0 ? WorkAuthorizations : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEmployee(employeeNumber, name, workEmail, workPhone, dateOfHire, nationalId, Manager, DirectReports, Department, PrimaryLocation, CostCenter, EmploymentAssignments, Contracts, BenefitEnrollments, Timesheets, LeaveRequests, PerformanceReviews, TrainingEnrollments, WorkAuthorizations, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Employee/update/' + id;
		const obj = {
				      		employeeNumber: employeeNumber,
      		name: name,
      		workEmail: workEmail,
      		workPhone: workPhone,
      		dateOfHire: dateOfHire,
      		nationalId: nationalId,
      		Manager: Manager != null && Manager.length > 0 ? Manager : null,
      		DirectReports: DirectReports != null && DirectReports.length > 0 ? DirectReports : null,
      		Department: Department != null && Department.length > 0 ? Department : null,
      		PrimaryLocation: PrimaryLocation != null && PrimaryLocation.length > 0 ? PrimaryLocation : null,
      		CostCenter: CostCenter != null && CostCenter.length > 0 ? CostCenter : null,
      		EmploymentAssignments: EmploymentAssignments != null && EmploymentAssignments.length > 0 ? EmploymentAssignments : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		BenefitEnrollments: BenefitEnrollments != null && BenefitEnrollments.length > 0 ? BenefitEnrollments : null,
      		Timesheets: Timesheets != null && Timesheets.length > 0 ? Timesheets : null,
      		LeaveRequests: LeaveRequests != null && LeaveRequests.length > 0 ? LeaveRequests : null,
      		PerformanceReviews: PerformanceReviews != null && PerformanceReviews.length > 0 ? PerformanceReviews : null,
      		TrainingEnrollments: TrainingEnrollments != null && TrainingEnrollments.length > 0 ? TrainingEnrollments : null,
      		WorkAuthorizations: WorkAuthorizations != null && WorkAuthorizations.length > 0 ? WorkAuthorizations : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEmployee(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Employee/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Employee
	// returns the results untouched as an Observable Employee
	// Employee model
	// delegates via URI
	//********************************************************************
	getEmployee(id) : Observable<Employee> {
		const uri_ = this.apiUrl + '/Employee/load/' + id;

		return this.http.get<Employee>(uri_);
	}
	
	//********************************************************************
	// gets all Employee
	// returns the results untouched as JSON representation of an
	// Observable array of Employee models
	// delegates via URI
	//********************************************************************
	getEmployees() : Observable<Employee[]> {
		const uri_ = this.apiUrl + '/Employee/';

		return this
			.http.get<Employee[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Manager on a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignManager( employeeId, _managerId ): Observable<any> {

		// get the Employee from storage
		this.loadHelper( employeeId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_managerId);

	// assign the Manager
	this.employee.manager = tmp;

	// save the Employee
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Manager on a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignManager( employeeId ): Observable<any> {

		// get the Employee from storage
		this.loadHelper( employeeId );

	// assign Manager to null
	this.employee.manager = null;

	// save the Employee
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Department on a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDepartment( employeeId, _departmentId ): Observable<any> {

		// get the Employee from storage
		this.loadHelper( employeeId );

	// get the Department from storage
	var tmp 	= new DepartmentService(this.http).getDepartment(_departmentId);

	// assign the Department
	this.employee.department = tmp;

	// save the Employee
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Department on a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDepartment( employeeId ): Observable<any> {

		// get the Employee from storage
		this.loadHelper( employeeId );

	// assign Department to null
	this.employee.department = null;

	// save the Employee
	return this.saveHelper();
}

		//********************************************************************
	// assigns a PrimaryLocation on a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPrimaryLocation( employeeId, _primaryLocationId ): Observable<any> {

		// get the Employee from storage
		this.loadHelper( employeeId );

	// get the Location from storage
	var tmp 	= new LocationService(this.http).getLocation(_primaryLocationId);

	// assign the PrimaryLocation
	this.employee.primaryLocation = tmp;

	// save the Employee
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PrimaryLocation on a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPrimaryLocation( employeeId ): Observable<any> {

		// get the Employee from storage
		this.loadHelper( employeeId );

	// assign PrimaryLocation to null
	this.employee.primaryLocation = null;

	// save the Employee
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CostCenter on a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCostCenter( employeeId, _costCenterId ): Observable<any> {

		// get the Employee from storage
		this.loadHelper( employeeId );

	// get the CostCenter from storage
	var tmp 	= new CostCenterService(this.http).getCostCenter(_costCenterId);

	// assign the CostCenter
	this.employee.costCenter = tmp;

	// save the Employee
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CostCenter on a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCostCenter( employeeId ): Observable<any> {

		// get the Employee from storage
		this.loadHelper( employeeId );

	// assign CostCenter to null
	this.employee.costCenter = null;

	// save the Employee
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more directReportsIds as a DirectReports
	// to a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDirectReports( employeeId, directReportsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );

	// split on a comma with no spaces
	var idList = directReportsIds.split(',')

	// iterate over array of directReports ids
	idList.forEach(function (id) {
		// read the Employee
		var employee = new EmployeeService(this.http).getEmployee(id);
		// add the Employee if not already assigned
		if ( this.employee.directReports.indexOf(employee) == -1 )
		this.employee.directReports.push(employee);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more directReportsIds as a DirectReports
	// from a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDirectReports( employeeId, directReportsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );


	// split on a comma with no spaces
	var idList 					= directReportsIds.split(',');
	var directReports 	= this.employee.directReports;

	if ( directReports != null && directReportsIds != null ) {

		// iterate over array of directReports ids
		directReports.forEach(function (obj) {
			if ( directReportsIds.indexOf(obj._id) > -1 ) {
				// remove the Employee
				this.employee.directReports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more employmentAssignmentsIds as a EmploymentAssignments
	// to a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEmploymentAssignments( employeeId, employmentAssignmentsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );

	// split on a comma with no spaces
	var idList = employmentAssignmentsIds.split(',')

	// iterate over array of employmentAssignments ids
	idList.forEach(function (id) {
		// read the EmploymentAssignment
		var employmentAssignment = new EmploymentAssignmentService(this.http).getEmploymentAssignment(id);
		// add the EmploymentAssignment if not already assigned
		if ( this.employee.employmentAssignments.indexOf(employmentAssignment) == -1 )
		this.employee.employmentAssignments.push(employmentAssignment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more employmentAssignmentsIds as a EmploymentAssignments
	// from a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEmploymentAssignments( employeeId, employmentAssignmentsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );


	// split on a comma with no spaces
	var idList 					= employmentAssignmentsIds.split(',');
	var employmentAssignments 	= this.employee.employmentAssignments;

	if ( employmentAssignments != null && employmentAssignmentsIds != null ) {

		// iterate over array of employmentAssignments ids
		employmentAssignments.forEach(function (obj) {
			if ( employmentAssignmentsIds.indexOf(obj._id) > -1 ) {
				// remove the EmploymentAssignment
				this.employee.employmentAssignments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more contractsIds as a Contracts
	// to a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContracts( employeeId, contractsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );

	// split on a comma with no spaces
	var idList = contractsIds.split(',')

	// iterate over array of contracts ids
	idList.forEach(function (id) {
		// read the EmploymentContract
		var employmentContract = new EmploymentContractService(this.http).getEmploymentContract(id);
		// add the EmploymentContract if not already assigned
		if ( this.employee.contracts.indexOf(employmentContract) == -1 )
		this.employee.contracts.push(employmentContract);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contractsIds as a Contracts
	// from a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContracts( employeeId, contractsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );


	// split on a comma with no spaces
	var idList 					= contractsIds.split(',');
	var contracts 	= this.employee.contracts;

	if ( contracts != null && contractsIds != null ) {

		// iterate over array of contracts ids
		contracts.forEach(function (obj) {
			if ( contractsIds.indexOf(obj._id) > -1 ) {
				// remove the EmploymentContract
				this.employee.contracts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more benefitEnrollmentsIds as a BenefitEnrollments
	// to a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBenefitEnrollments( employeeId, benefitEnrollmentsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );

	// split on a comma with no spaces
	var idList = benefitEnrollmentsIds.split(',')

	// iterate over array of benefitEnrollments ids
	idList.forEach(function (id) {
		// read the BenefitEnrollment
		var benefitEnrollment = new BenefitEnrollmentService(this.http).getBenefitEnrollment(id);
		// add the BenefitEnrollment if not already assigned
		if ( this.employee.benefitEnrollments.indexOf(benefitEnrollment) == -1 )
		this.employee.benefitEnrollments.push(benefitEnrollment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more benefitEnrollmentsIds as a BenefitEnrollments
	// from a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBenefitEnrollments( employeeId, benefitEnrollmentsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );


	// split on a comma with no spaces
	var idList 					= benefitEnrollmentsIds.split(',');
	var benefitEnrollments 	= this.employee.benefitEnrollments;

	if ( benefitEnrollments != null && benefitEnrollmentsIds != null ) {

		// iterate over array of benefitEnrollments ids
		benefitEnrollments.forEach(function (obj) {
			if ( benefitEnrollmentsIds.indexOf(obj._id) > -1 ) {
				// remove the BenefitEnrollment
				this.employee.benefitEnrollments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more timesheetsIds as a Timesheets
	// to a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTimesheets( employeeId, timesheetsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );

	// split on a comma with no spaces
	var idList = timesheetsIds.split(',')

	// iterate over array of timesheets ids
	idList.forEach(function (id) {
		// read the Timesheet
		var timesheet = new TimesheetService(this.http).getTimesheet(id);
		// add the Timesheet if not already assigned
		if ( this.employee.timesheets.indexOf(timesheet) == -1 )
		this.employee.timesheets.push(timesheet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more timesheetsIds as a Timesheets
	// from a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTimesheets( employeeId, timesheetsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );


	// split on a comma with no spaces
	var idList 					= timesheetsIds.split(',');
	var timesheets 	= this.employee.timesheets;

	if ( timesheets != null && timesheetsIds != null ) {

		// iterate over array of timesheets ids
		timesheets.forEach(function (obj) {
			if ( timesheetsIds.indexOf(obj._id) > -1 ) {
				// remove the Timesheet
				this.employee.timesheets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more leaveRequestsIds as a LeaveRequests
	// to a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLeaveRequests( employeeId, leaveRequestsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );

	// split on a comma with no spaces
	var idList = leaveRequestsIds.split(',')

	// iterate over array of leaveRequests ids
	idList.forEach(function (id) {
		// read the LeaveRequest
		var leaveRequest = new LeaveRequestService(this.http).getLeaveRequest(id);
		// add the LeaveRequest if not already assigned
		if ( this.employee.leaveRequests.indexOf(leaveRequest) == -1 )
		this.employee.leaveRequests.push(leaveRequest);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more leaveRequestsIds as a LeaveRequests
	// from a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLeaveRequests( employeeId, leaveRequestsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );


	// split on a comma with no spaces
	var idList 					= leaveRequestsIds.split(',');
	var leaveRequests 	= this.employee.leaveRequests;

	if ( leaveRequests != null && leaveRequestsIds != null ) {

		// iterate over array of leaveRequests ids
		leaveRequests.forEach(function (obj) {
			if ( leaveRequestsIds.indexOf(obj._id) > -1 ) {
				// remove the LeaveRequest
				this.employee.leaveRequests.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more performanceReviewsIds as a PerformanceReviews
	// to a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPerformanceReviews( employeeId, performanceReviewsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );

	// split on a comma with no spaces
	var idList = performanceReviewsIds.split(',')

	// iterate over array of performanceReviews ids
	idList.forEach(function (id) {
		// read the PerformanceReview
		var performanceReview = new PerformanceReviewService(this.http).getPerformanceReview(id);
		// add the PerformanceReview if not already assigned
		if ( this.employee.performanceReviews.indexOf(performanceReview) == -1 )
		this.employee.performanceReviews.push(performanceReview);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more performanceReviewsIds as a PerformanceReviews
	// from a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePerformanceReviews( employeeId, performanceReviewsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );


	// split on a comma with no spaces
	var idList 					= performanceReviewsIds.split(',');
	var performanceReviews 	= this.employee.performanceReviews;

	if ( performanceReviews != null && performanceReviewsIds != null ) {

		// iterate over array of performanceReviews ids
		performanceReviews.forEach(function (obj) {
			if ( performanceReviewsIds.indexOf(obj._id) > -1 ) {
				// remove the PerformanceReview
				this.employee.performanceReviews.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more trainingEnrollmentsIds as a TrainingEnrollments
	// to a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTrainingEnrollments( employeeId, trainingEnrollmentsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );

	// split on a comma with no spaces
	var idList = trainingEnrollmentsIds.split(',')

	// iterate over array of trainingEnrollments ids
	idList.forEach(function (id) {
		// read the TrainingEnrollment
		var trainingEnrollment = new TrainingEnrollmentService(this.http).getTrainingEnrollment(id);
		// add the TrainingEnrollment if not already assigned
		if ( this.employee.trainingEnrollments.indexOf(trainingEnrollment) == -1 )
		this.employee.trainingEnrollments.push(trainingEnrollment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more trainingEnrollmentsIds as a TrainingEnrollments
	// from a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTrainingEnrollments( employeeId, trainingEnrollmentsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );


	// split on a comma with no spaces
	var idList 					= trainingEnrollmentsIds.split(',');
	var trainingEnrollments 	= this.employee.trainingEnrollments;

	if ( trainingEnrollments != null && trainingEnrollmentsIds != null ) {

		// iterate over array of trainingEnrollments ids
		trainingEnrollments.forEach(function (obj) {
			if ( trainingEnrollmentsIds.indexOf(obj._id) > -1 ) {
				// remove the TrainingEnrollment
				this.employee.trainingEnrollments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more workAuthorizationsIds as a WorkAuthorizations
	// to a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addWorkAuthorizations( employeeId, workAuthorizationsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );

	// split on a comma with no spaces
	var idList = workAuthorizationsIds.split(',')

	// iterate over array of workAuthorizations ids
	idList.forEach(function (id) {
		// read the WorkAuthorization
		var workAuthorization = new WorkAuthorizationService(this.http).getWorkAuthorization(id);
		// add the WorkAuthorization if not already assigned
		if ( this.employee.workAuthorizations.indexOf(workAuthorization) == -1 )
		this.employee.workAuthorizations.push(workAuthorization);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more workAuthorizationsIds as a WorkAuthorizations
	// from a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeWorkAuthorizations( employeeId, workAuthorizationsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );


	// split on a comma with no spaces
	var idList 					= workAuthorizationsIds.split(',');
	var workAuthorizations 	= this.employee.workAuthorizations;

	if ( workAuthorizations != null && workAuthorizationsIds != null ) {

		// iterate over array of workAuthorizations ids
		workAuthorizations.forEach(function (obj) {
			if ( workAuthorizationsIds.indexOf(obj._id) > -1 ) {
				// remove the WorkAuthorization
				this.employee.workAuthorizations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Employee
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Employee/update/' + this.employee;

	return  this.http.post(uri_, this.employee );
}

	//********************************************************************
	// loadHelper - internal helper to load a Employee
	//********************************************************************	
	loadHelper( id ) {
		this.getEmployee(id)
			.subscribe((res : Employee) => {
				this.employee = res;
			});
	}
}