import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DependentService } from '../../../services/Dependent.service';
import { Dependent } from '../../../models/Dependent';
import { SubBaseComponent } from '../../Dependent/sub.base.component';

@Component({
    selector: 'app-create-dependent',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDependentComponent extends SubBaseComponent implements OnInit {

    title = 'Add Dependent';

    dependentForm: FormGroup;
    dependent: Dependent;

    constructor( http: HttpClient,
        private dependentService: DependentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.dependentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      birthDate: ['', Validators.required],
      BenefitEnrollment: ['', ],
      Employee: ['', ],
      Relationship: ['', ]
        });
    }

    
    addDependent(firstName, lastName, birthDate, BenefitEnrollment, Employee, Relationship): void {
        this.dependentService
        .addDependent(firstName, lastName, birthDate, BenefitEnrollment, Employee, Relationship)
            .subscribe(() => {
                this.router.navigate(['/indexDependent']);
            });
    }

    ngOnInit(): void {
    }
}