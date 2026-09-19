Rails.application.routes.draw do
  root "application#health"
  resources :organizations do
    resources :departments
    resources :locations
    resources :jobfamilies
    resources :benefitplans
    resources :costcenters
    resources :payrollcalendars
  end
  resources :departments do
    resource :organization
    resource :manager
    resources :positions
    resources :employees
    resource :costcenter
  end
  resources :locations do
    resource :organization
    resources :departments
    resources :positions
    resources :employees
  end
  resources :costcenters do
    resource :organization
    resources :departments
    resources :positions
    resources :employees
  end
  resources :jobfamilys do
    resource :organization
    resources :jobprofiles
  end
  resources :jobprofiles do
    resource :jobfamily
    resources :competencies
    resources :trainingrecommendations
    resources :positions
  end
  resources :competencys do
    resources :jobprofiles
    resources :competencyratings
  end
  resources :positions do
    resource :department
    resource :jobprofile
    resource :costcenter
    resource :location
    resource :managerposition
    resources :directreports
    resources :assignments
  end
  resources :employees do
    resource :manager
    resources :directreports
    resource :department
    resource :primarylocation
    resource :costcenter
    resources :employmentassignments
    resources :contracts
    resources :benefitenrollments
    resources :timesheets
    resources :leaverequests
    resources :performancereviews
    resources :trainingenrollments
    resources :workauthorizations
  end
  resources :employmentassignments do
    resource :employee
    resource :position
    resource :supervisor
  end
  resources :employmentcontracts do
    resource :employee
    resource :compensationpackage
    resource :workschedule
    resource :location
    resource :payrollcalendar
  end
  resources :workschedules do
    resources :contracts
    resources :shifts
    resources :exceptions
  end
  resources :workshifts do
    resource :workschedule
  end
  resources :scheduleexceptions do
    resource :workschedule
    resource :employee
  end
  resources :compensationpackages do
    resource :contract
    resources :salarycomponents
    resources :bonusplans
    resources :equitygrants
  end
  resources :salarycomponents do
    resource :compensationpackage
  end
  resources :bonusplans do
    resources :compensationpackages
  end
  resources :equitygrants do
    resource :compensationpackage
  end
  resources :benefitplans do
    resource :organization
    resources :enrollments
  end
  resources :benefitenrollments do
    resource :benefitplan
    resource :employee
    resources :dependents
  end
  resources :dependents do
    resource :benefitenrollment
    resource :employee
  end
  resources :payrollcalendars do
    resource :organization
    resources :payrollruns
    resources :employees
  end
  resources :payrollruns do
    resource :payrollcalendar
    resources :payrollitems
  end
  resources :payrollitems do
    resource :payrollrun
    resource :employee
  end
  resources :taxwithholdings do
    resource :employee
  end
  resources :paymentmethods do
    resource :employee
    resource :bankaccount
  end
  resources :timesheets do
    resource :employee
    resources :timeentries
    resources :approvals
  end
  resources :timeentrys do
    resource :timesheet
    resource :employee
    resource :costcenter
  end
  resources :approvals do
    resource :approver
    resource :timesheet
    resource :leaverequest
  end
  resources :leavepolicys do
    resource :organization
    resources :leaverequests
  end
  resources :leaverequests do
    resource :employee
    resource :leavepolicy
    resources :approvals
  end
  resources :performancecycles do
    resource :organization
    resources :reviews
    resources :goals
  end
  resources :goals do
    resource :employee
    resource :cycle
    resource :parentgoal
    resources :childgoals
  end
  resources :performancereviews do
    resource :employee
    resource :reviewer
    resource :cycle
    resources :competencyratings
    resources :goals
  end
  resources :competencyratings do
    resource :review
    resource :competency
  end
  resources :trainingcourses do
    resources :prerequisites
    resources :enrollments
    resources :jobprofiles
  end
  resources :trainingenrollments do
    resource :course
    resource :employee
    resource :instructor
  end
  resources :certifications do
    resource :employee
    resource :course
  end
  resources :jobrequisitions do
    resource :department
    resource :hiringmanager
    resource :recruiter
    resource :jobprofile
    resources :candidates
    resources :interviews
    resources :offers
  end
  resources :candidates do
    resources :applications
    resources :interviews
    resources :offers
    resources :documents
  end
  resources :jobapplications do
    resource :candidate
    resource :requisition
    resources :screenings
  end
  resources :interviews do
    resource :requisition
    resource :candidate
    resources :interviewers
  end
  resources :screenings do
    resource :application
  end
  resources :offers do
    resource :requisition
    resource :candidate
    resource :approvedby
    resource :contract
  end
  resources :onboardingtasks do
    resource :employee
    resource :assignedto
    resources :dependencies
    resource :relatedoffer
  end
  resources :backgroundchecks do
    resource :candidate
    resource :requisition
    resource :report
  end
  resources :documents do
    resource :candidate
    resource :employee
  end
  resources :policys do
    resource :organization
    resources :acknowledgements
  end
  resources :policyacknowledgements do
    resource :policy
    resource :employee
  end
  resources :terminations do
    resource :employee
    resource :assignment
  end
  resources :workauthorizations do
    resource :employee
    resources :documents
  end
  resources :bankaccounts
end
