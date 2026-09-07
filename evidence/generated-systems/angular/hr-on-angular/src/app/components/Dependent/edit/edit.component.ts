import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DependentService } from '../../../services/Dependent.service';
import { SubBaseComponent } from '../../Dependent/sub.base.component';


@Component({
    selector: 'app-edit-dependent',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDependentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Dependent';

    dependentForm: FormGroup;
    dependent: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DependentService,
        private fb: FormBuilder
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

    
    updateDependent(firstName, lastName, birthDate, BenefitEnrollment, Employee, Relationship): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDependent(firstName, lastName, birthDate, BenefitEnrollment, Employee, Relationship, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDependent']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDependent(params['id']).subscribe(res => {
                this.dependent = res;
            });
        });
    }
}