class InitialMigration < ActiveRecord::Migration[6.1]
  def change
    create_table :healthSystems do |t|
      t.string :name      
      t.string :legalName      
      t.string :headquartersCountry      
      t.string :website      
      t.timestamps
    end
    create_table :facilitys do |t|
      t.string :name      
      t.string :facilityCode      
      t.string :address      
      t.integer :FacilityType      
      t.timestamps
    end
    create_table :departments do |t|
      t.string :name      
      t.integer :DepartmentType      
      t.timestamps
    end
    create_table :careTeams do |t|
      t.string :name      
      t.integer :CareSetting      
      t.timestamps
    end
    create_table :clinicians do |t|
      t.string :firstName      
      t.string :lastName      
      t.string :licenseNumber      
      t.integer :ClinicianType      
      t.integer :Specialty      
      t.timestamps
    end
    create_table :patients do |t|
      t.string :firstName      
      t.string :lastName      
      t.string :mrn      
      t.date :dateOfBirth      
      t.string :address      
      t.string :primaryLanguage      
      t.integer :SexAtBirth      
      t.integer :BloodType      
      t.timestamps
    end
    create_table :appointments do |t|
      t.datetime :appointmentDate      
      t.string :reason      
      t.integer :Status      
      t.integer :Priority      
      t.timestamps
    end
    create_table :encounters do |t|
      t.string :encounterNumber      
      t.datetime :startDateTime      
      t.datetime :endDateTime      
      t.integer :Status      
      t.integer :EncounterType      
      t.timestamps
    end
    create_table :admissions do |t|
      t.datetime :admitDateTime      
      t.string :bed      
      t.integer :AdmissionType      
      t.timestamps
    end
    create_table :discharges do |t|
      t.datetime :dischargeDateTime      
      t.integer :Disposition      
      t.timestamps
    end
    create_table :clinicalOrders do |t|
      t.string :orderNumber      
      t.integer :Status      
      t.integer :OrderType      
      t.integer :Priority      
      t.timestamps
    end
    create_table :medicationOrders do |t|
      t.string :medicationCode      
      t.string :dose      
      t.string :frequency      
      t.string :duration      
      t.integer :Route      
      t.timestamps
    end
    create_table :laboratorys do |t|
      t.string :name      
      t.string :cliaNumber      
      t.timestamps
    end
    create_table :laboratoryOrders do |t|
      t.string :testCode      
      t.boolean :fastingRequired      
      t.integer :SpecimenType      
      t.timestamps
    end
    create_table :labResults do |t|
      t.string :resultCode      
      t.datetime :issuedDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :imagingCenters do |t|
      t.string :name      
      t.timestamps
    end
    create_table :imagingOrders do |t|
      t.string :bodySite      
      t.boolean :contrast      
      t.integer :Modality      
      t.timestamps
    end
    create_table :imagingReports do |t|
      t.string :reportNumber      
      t.string :impression      
      t.datetime :reportedDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :procedureOrders do |t|
      t.string :procedureCode      
      t.boolean :consentObtained      
      t.integer :AnesthesiaType      
      t.timestamps
    end
    create_table :procedures do |t|
      t.string :procedureCode      
      t.datetime :startDateTime      
      t.datetime :endDateTime      
      t.integer :Status      
      t.timestamps
    end
    create_table :pharmacys do |t|
      t.string :name      
      t.timestamps
    end
    create_table :medicationDispenses do |t|
      t.string :dispenseNumber      
      t.decimal :quantity      
      t.datetime :whenPrepared      
      t.integer :Status      
      t.timestamps
    end
    create_table :diagnosiss do |t|
      t.string :code      
      t.string :description      
      t.date :onsetDate      
      t.integer :Certainty      
      t.timestamps
    end
    create_table :observations do |t|
      t.string :code      
      t.string :value      
      t.string :unit      
      t.datetime :effectiveDateTime      
      t.integer :Interpretation      
      t.timestamps
    end
    create_table :carePlans do |t|
      t.string :planNumber      
      t.string :goalSummary      
      t.integer :Status      
      t.timestamps
    end
    create_table :careTasks do |t|
      t.string :description      
      t.date :dueDate      
      t.integer :Status      
      t.integer :Priority      
      t.timestamps
    end
    create_table :allergys do |t|
      t.string :substance      
      t.string :reaction      
      t.integer :Severity      
      t.integer :Status      
      t.timestamps
    end
    create_table :conditions do |t|
      t.string :code      
      t.date :onsetDate      
      t.date :abatementDate      
      t.integer :ClinicalStatus      
      t.integer :VerificationStatus      
      t.timestamps
    end
    create_table :insurancePayers do |t|
      t.string :name      
      t.string :website      
      t.integer :PayerType      
      t.timestamps
    end
    create_table :insurancePlans do |t|
      t.string :name      
      t.string :planCode      
      t.integer :PlanType      
      t.timestamps
    end
    create_table :coverages do |t|
      t.string :memberId      
      t.string :groupNumber      
      t.date :effectiveDate      
      t.date :endDate      
      t.integer :CoverageType      
      t.timestamps
    end
    create_table :claims do |t|
      t.string :claimNumber      
      t.string :totalAmount      
      t.integer :Status      
      t.timestamps
    end
    create_table :authorizations do |t|
      t.string :authNumber      
      t.string :requestedService      
      t.integer :Status      
      t.timestamps
    end
    create_table :invoices do |t|
      t.string :invoiceNumber      
      t.string :totalAmount      
      t.date :dueDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :payments do |t|
      t.string :paymentNumber      
      t.string :amount      
      t.date :paymentDate      
      t.integer :Method      
      t.timestamps
    end
    create_table :medicalDevices do |t|
      t.string :udi      
      t.string :manufacturer      
      t.integer :DeviceType      
      t.integer :ConnectivityStatus      
      t.timestamps
    end
    create_table :softwareUpdates do |t|
      t.string :version      
      t.datetime :appliedDate      
      t.integer :UpdateType      
      t.timestamps
    end
    create_table :medicalSuppliers do |t|
      t.string :name      
      t.string :website      
      t.integer :SupplierTier      
      t.timestamps
    end
    create_table :inventoryItems do |t|
      t.string :sku      
      t.string :name      
      t.integer :quantityOnHand      
      t.integer :quantityReserved      
      t.timestamps
    end
  end
end
