import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AdmissionService } from '../../../services/Admission.service';
import { Admission } from '../../../models/Admission';
import { SubBaseComponent } from '../../Admission/sub.base.component';

@Component({
    selector: 'app-create-admission',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAdmissionComponent extends SubBaseComponent implements OnInit {

    title = 'Add Admission';

    admissionForm: FormGroup;
    admission: Admission;

    constructor( http: HttpClient,
        private admissionService: AdmissionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.admissionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  admitDateTime: ['', Validators.required],
      bed: ['', Validators.required],
      Encounter: ['', ],
      Facility: ['', ],
      AdmissionType: ['', ]
        });
    }

    
    addAdmission(admitDateTime, bed, Encounter, Facility, AdmissionType): void {
        this.admissionService
        .addAdmission(admitDateTime, bed, Encounter, Facility, AdmissionType)
            .subscribe(() => {
                this.router.navigate(['/indexAdmission']);
            });
    }

    ngOnInit(): void {
    }
}