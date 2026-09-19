class InitialMigration < ActiveRecord::Migration[6.1]
  def change
    create_table :organizations do |t|
      t.string :name      
      t.string :legalName      
      t.string :registrationCountry      
      t.string :website      
      t.timestamps
    end
    create_table :departments do |t|
      t.string :name      
      t.string :code      
      t.timestamps
    end
    create_table :locations do |t|
      t.string :name      
      t.string :address      
      t.string :timezone      
      t.timestamps
    end
    create_table :costCenters do |t|
      t.string :code      
      t.string :name      
      t.timestamps
    end
    create_table :jobFamilys do |t|
      t.string :name      
      t.string :description      
      t.timestamps
    end
    create_table :jobProfiles do |t|
      t.string :title      
      t.string :jobCode      
      t.integer :JobLevel      
      t.integer :ExemptStatus      
      t.timestamps
    end
    create_table :competencys do |t|
      t.string :name      
      t.string :category      
      t.timestamps
    end
    create_table :positions do |t|
      t.string :positionCode      
      t.decimal :fte      
      t.integer :Status      
      t.integer :WorkLocationType      
      t.timestamps
    end
    create_table :employees do |t|
      t.string :employeeNumber      
      t.string :name      
      t.string :workEmail      
      t.string :workPhone      
      t.date :dateOfHire      
      t.string :nationalId      
      t.integer :Status      
      t.timestamps
    end
    create_table :employmentAssignments do |t|
      t.date :startDate      
      t.date :endDate      
      t.boolean :primary      
      t.integer :AssignmentType      
      t.integer :Status      
      t.timestamps
    end
    create_table :employmentContracts do |t|
      t.string :contractNumber      
      t.date :startDate      
      t.date :endDate      
      t.decimal :workHoursPerWeek      
      t.integer :EmploymentType      
      t.integer :Status      
      t.integer :PayFrequency      
      t.timestamps
    end
    create_table :workSchedules do |t|
      t.string :name      
      t.decimal :standardHoursPerWeek      
      t.integer :ScheduleType      
      t.timestamps
    end
    create_table :workShifts do |t|
      t.datetime :startTime      
      t.datetime :endTime      
      t.integer :breakMinutes      
      t.integer :DayOfWeek      
      t.timestamps
    end
    create_table :scheduleExceptions do |t|
      t.date :date      
      t.string :reason      
      t.decimal :hours      
      t.timestamps
    end
    create_table :compensationPackages do |t|
      t.date :effectiveFrom      
      t.date :effectiveTo      
      t.string :currency      
      t.timestamps
    end
    create_table :salaryComponents do |t|
      t.string :amount      
      t.boolean :recurring      
      t.integer :ComponentType      
      t.timestamps
    end
    create_table :bonusPlans do |t|
      t.string :name      
      t.string :targetPercentage      
      t.timestamps
    end
    create_table :equityGrants do |t|
      t.string :grantId      
      t.integer :grantedUnits      
      t.date :vestingStart      
      t.integer :GrantType      
      t.timestamps
    end
    create_table :benefitPlans do |t|
      t.string :name      
      t.string :providerName      
      t.string :employeeContributionRate      
      t.string :employerContributionRate      
      t.string :eligibilityRules      
      t.integer :BenefitType      
      t.timestamps
    end
    create_table :benefitEnrollments do |t|
      t.string :enrollmentId      
      t.date :effectiveFrom      
      t.date :effectiveTo      
      t.integer :Status      
      t.integer :CoverageLevel      
      t.timestamps
    end
    create_table :dependents do |t|
      t.string :firstName      
      t.string :lastName      
      t.date :birthDate      
      t.integer :Relationship      
      t.timestamps
    end
    create_table :payrollCalendars do |t|
      t.string :name      
      t.string :country      
      t.integer :PayFrequency      
      t.timestamps
    end
    create_table :payrollRuns do |t|
      t.string :runNumber      
      t.date :periodStart      
      t.date :periodEnd      
      t.date :paymentDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :payrollItems do |t|
      t.string :amount      
      t.boolean :taxable      
      t.integer :ItemType      
      t.timestamps
    end
    create_table :taxWithholdings do |t|
      t.string :taxId      
      t.integer :allowances      
      t.string :additionalAmount      
      t.integer :FilingStatus      
      t.timestamps
    end
    create_table :paymentMethods do |t|
      t.boolean :preferred      
      t.integer :MethodType      
      t.timestamps
    end
    create_table :timesheets do |t|
      t.date :periodStart      
      t.date :periodEnd      
      t.date :submissionDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :timeEntrys do |t|
      t.date :entryDate      
      t.decimal :hoursWorked      
      t.integer :EntryType      
      t.timestamps
    end
    create_table :approvals do |t|
      t.string :approverComment      
      t.date :actionDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :leavePolicys do |t|
      t.string :name      
      t.decimal :accrualRate      
      t.boolean :carryoverAllowed      
      t.decimal :maxBalance      
      t.integer :LeaveCategory      
      t.integer :AccrualUnit      
      t.timestamps
    end
    create_table :leaveRequests do |t|
      t.string :requestNumber      
      t.date :startDate      
      t.date :endDate      
      t.string :reason      
      t.decimal :hours      
      t.integer :Status      
      t.timestamps
    end
    create_table :performanceCycles do |t|
      t.string :name      
      t.date :startDate      
      t.date :endDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :goals do |t|
      t.string :title      
      t.string :description      
      t.date :targetDate      
      t.string :weight      
      t.integer :Status      
      t.timestamps
    end
    create_table :performanceReviews do |t|
      t.string :reviewNumber      
      t.date :reviewDate      
      t.string :reviewerComments      
      t.integer :Rating      
      t.integer :Status      
      t.timestamps
    end
    create_table :competencyRatings do |t|
      t.string :comment      
      t.integer :Rating      
      t.timestamps
    end
    create_table :trainingCourses do |t|
      t.string :code      
      t.string :title      
      t.decimal :durationHours      
      t.integer :DeliveryMethod      
      t.timestamps
    end
    create_table :trainingEnrollments do |t|
      t.string :enrollmentNumber      
      t.date :completionDate      
      t.decimal :score      
      t.integer :Status      
      t.timestamps
    end
    create_table :certifications do |t|
      t.string :name      
      t.string :issuer      
      t.date :validFrom      
      t.date :validTo      
      t.string :credentialId      
      t.timestamps
    end
    create_table :jobRequisitions do |t|
      t.string :requisitionNumber      
      t.string :title      
      t.integer :openings      
      t.date :targetStartDate      
      t.integer :Status      
      t.integer :Priority      
      t.timestamps
    end
    create_table :candidates do |t|
      t.string :name      
      t.string :email      
      t.string :phone      
      t.integer :Source      
      t.timestamps
    end
    create_table :jobApplications do |t|
      t.string :applicationNumber      
      t.date :appliedDate      
      t.string :resumeUrl      
      t.integer :Status      
      t.timestamps
    end
    create_table :interviews do |t|
      t.date :interviewDate      
      t.string :feedback      
      t.integer :Stage      
      t.integer :Result      
      t.timestamps
    end
    create_table :screenings do |t|
      t.string :name      
      t.date :completedDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :offers do |t|
      t.string :offerNumber      
      t.date :proposedStartDate      
      t.string :baseSalary      
      t.string :signOnBonus      
      t.integer :Status      
      t.timestamps
    end
    create_table :onboardingTasks do |t|
      t.string :taskNumber      
      t.string :name      
      t.date :dueDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :backgroundChecks do |t|
      t.string :checkNumber      
      t.string :provider      
      t.date :completedDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :documents do |t|
      t.string :name      
      t.string :fileUrl      
      t.date :uploadedDate      
      t.integer :DocumentType      
      t.timestamps
    end
    create_table :policys do |t|
      t.string :policyNumber      
      t.string :name      
      t.date :effectiveDate      
      t.string :description      
      t.timestamps
    end
    create_table :policyAcknowledgements do |t|
      t.date :acknowledgementDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :terminations do |t|
      t.string :terminationNumber      
      t.date :terminationDate      
      t.string :notes      
      t.boolean :eligibleForRehire      
      t.integer :Reason      
      t.integer :Type      
      t.timestamps
    end
    create_table :workAuthorizations do |t|
      t.string :country      
      t.date :expirationDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :bankAccounts do |t|
      t.string :accountHolder      
      t.string :bankName      
      t.string :iban      
      t.string :bic      
      t.string :accountNumber      
      t.string :routingNumber      
      t.timestamps
    end
  end
end
