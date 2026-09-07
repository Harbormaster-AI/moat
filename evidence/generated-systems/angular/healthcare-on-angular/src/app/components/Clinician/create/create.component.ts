import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ClinicianService } from '../../../services/Clinician.service';
import { Clinician } from '../../../models/Clinician';
import { SubBaseComponent } from '../../Clinician/sub.base.component';

@Component({
    selector: 'app-create-clinician',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateClinicianComponent extends SubBaseComponent implements OnInit {

    title = 'Add Clinician';

    clinicianForm: FormGroup;
    clinician: Clinician;

    constructor( http: HttpClient,
        private clinicianService: ClinicianService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.clinicianForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      licenseNumber: ['', Validators.required],
      CareTeams: ['', ],
      Appointments: ['', ],
      Encounters: ['', ],
      Procedures: ['', ],
      ImagingReports: ['', ],
      ClinicianType: ['', ],
      Specialty: ['', ]
        });
    }

    
    addClinician(firstName, lastName, licenseNumber, CareTeams, Appointments, Encounters, Procedures, ImagingReports, ClinicianType, Specialty): void {
        this.clinicianService
        .addClinician(firstName, lastName, licenseNumber, CareTeams, Appointments, Encounters, Procedures, ImagingReports, ClinicianType, Specialty)
            .subscribe(() => {
                this.router.navigate(['/indexClinician']);
            });
    }

    ngOnInit(): void {
    }
}