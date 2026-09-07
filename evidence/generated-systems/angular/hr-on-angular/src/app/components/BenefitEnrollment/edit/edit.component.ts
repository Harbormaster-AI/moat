import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BenefitEnrollmentService } from '../../../services/BenefitEnrollment.service';
import { SubBaseComponent } from '../../BenefitEnrollment/sub.base.component';


@Component({
    selector: 'app-edit-benefitEnrollment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBenefitEnrollmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BenefitEnrollment';

    benefitEnrollmentForm: FormGroup;
    benefitEnrollment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BenefitEnrollmentService,
        private fb: FormBuilder
) {
        super(http);
        this.benefitEnrollmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  enrollmentId: ['', Validators.required],
      effectiveFrom: ['', Validators.required],
      effectiveTo: ['', Validators.required],
      BenefitPlan: ['', ],
      Employee: ['', ],
      Dependents: ['', ],
      Status: ['', ],
      CoverageLevel: ['', ]
        });
    }

    
    updateBenefitEnrollment(enrollmentId, effectiveFrom, effectiveTo, BenefitPlan, Employee, Dependents, Status, CoverageLevel): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBenefitEnrollment(enrollmentId, effectiveFrom, effectiveTo, BenefitPlan, Employee, Dependents, Status, CoverageLevel, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBenefitEnrollment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBenefitEnrollment(params['id']).subscribe(res => {
                this.benefitEnrollment = res;
            });
        });
    }
}